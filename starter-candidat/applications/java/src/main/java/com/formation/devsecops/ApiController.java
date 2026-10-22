package com.formation.devsecops;

import java.net.InetAddress;
import java.net.UnknownHostException;
import java.time.Instant;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * Endpoints de l'API :
 *   GET /             message de bienvenue
 *   GET /health       etat de sante
 *   GET /api/info     informations d'execution
 *   GET /api/version  version applicative
 */
@RestController
public class ApiController {

    private final String appName;
    private final String appVersion;
    private final String appEnvironment;

    public ApiController(@Value("${spring.application.name}") String appName,
                         @Value("${app.version}") String appVersion,
                         @Value("${app.environment}") String appEnvironment) {
        this.appName = appName;
        this.appVersion = appVersion;
        this.appEnvironment = appEnvironment;
    }

    @GetMapping("/")
    public Map<String, Object> home() {
        Map<String, Object> body = new LinkedHashMap<>();
        body.put("message", "Bienvenue sur l'API Java Spring Boot du TP Docker");
        body.put("application", appName);
        body.put("endpoints", List.of("/", "/health", "/api/info", "/api/version"));
        return body;
    }

    @GetMapping("/health")
    public Map<String, Object> health() {
        Map<String, Object> body = new LinkedHashMap<>();
        body.put("status", "UP");
        return body;
    }

    @GetMapping("/api/info")
    public Map<String, Object> info() {
        Map<String, Object> body = new LinkedHashMap<>();
        body.put("application", appName);
        body.put("version", appVersion);
        body.put("environment", appEnvironment);
        body.put("language", "Java");
        body.put("runtime", System.getProperty("java.vm.name") + " " + System.getProperty("java.version"));
        body.put("hostname", hostname());
        body.put("platform", System.getProperty("os.name") + "/" + System.getProperty("os.arch"));
        body.put("user", System.getProperty("user.name"));
        body.put("maxHeapMb", Runtime.getRuntime().maxMemory() / (1024 * 1024));
        body.put("timestamp", Instant.now().toString());
        return body;
    }

    @GetMapping("/api/version")
    public Map<String, Object> version() {
        Map<String, Object> body = new LinkedHashMap<>();
        body.put("application", appName);
        body.put("version", appVersion);
        return body;
    }

    private static String hostname() {
        try {
            return InetAddress.getLocalHost().getHostName();
        } catch (UnknownHostException e) {
            return "unknown";
        }
    }
}
