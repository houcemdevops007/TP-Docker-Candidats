// TP Docker DevSecOps - Application 4 : API HTTP Go (bibliotheque standard)
// Auteur : Khlifi Houcem - Formateur DevSecOps et Cloud
//
// Endpoints :
//
//	GET /             message de bienvenue
//	GET /health       etat de sante
//	GET /api/info     informations d'execution
//	GET /api/version  version applicative
//
// Le binaire integre son propre healthcheck (option -healthcheck) afin de
// pouvoir etre surveille dans des images sans shell ni curl (scratch, distroless).
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// Valeurs surchargeables a la compilation :
//
//	go build -ldflags "-X main.version=1.2.3 -X main.commit=abc123"
var (
	version   = "1.0.0"
	commit    = "dev"
	buildDate = "unknown"
)

const appName = "tp-docker-golang"

func main() {
	healthcheck := flag.Bool("healthcheck", false, "interroge /health sur le port local puis quitte (0 = sain, 1 = en echec)")
	flag.Parse()

	port := getenv("PORT", "8080")

	if *healthcheck {
		os.Exit(runHealthcheck(port))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", handleRoot)
	mux.HandleFunc("GET /health", handleHealth)
	mux.HandleFunc("GET /api/info", handleInfo)
	mux.HandleFunc("GET /api/version", handleVersion)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           logRequests(mux),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		log.Printf("[%s] version %s ecoute sur :%s", appName, version, port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("erreur serveur : %v", err)
		}
	}()

	// Arret propre sur SIGTERM (docker stop) ou SIGINT (Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stop
	log.Printf("[%s] signal %s recu, arret en cours...", appName, sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("arret force : %v", err)
	}
}

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"message":     "Bienvenue sur l'API Go du TP Docker",
		"application": appName,
		"endpoints":   []string{"/", "/health", "/api/info", "/api/version"},
	})
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "UP"})
}

func handleInfo(w http.ResponseWriter, _ *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"application": appName,
		"version":     version,
		"commit":      commit,
		"buildDate":   buildDate,
		"environment": getenv("APP_ENV", "dev"),
		"language":    "Go",
		"runtime":     runtime.Version(),
		"hostname":    hostname,
		"platform":    runtime.GOOS + "/" + runtime.GOARCH,
		"user":        fmt.Sprintf("uid=%d gid=%d", os.Getuid(), os.Getgid()),
		"timestamp":   time.Now().UTC().Format(time.RFC3339),
	})
}

func handleVersion(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"application": appName,
		"version":     version,
		"commit":      commit,
	})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("erreur d'encodage JSON : %v", err)
	}
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func runHealthcheck(port string) int {
	client := http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get("http://127.0.0.1:" + port + "/health")
	if err != nil {
		fmt.Fprintln(os.Stderr, "healthcheck en echec :", err)
		return 1
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, "healthcheck en echec : HTTP", resp.StatusCode)
		return 1
	}
	return 0
}

func getenv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
