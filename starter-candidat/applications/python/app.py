"""
TP Docker DevSecOps - Application 2 : API REST Python / Flask
Auteur : Khlifi Houcem - Formateur DevSecOps et Cloud

Endpoints :
    GET /            message de bienvenue
    GET /health      etat de sante
    GET /api/info    informations d'execution
    GET /api/time    heure du serveur
"""
import os
import platform
import socket
import time
from datetime import datetime, timezone

from flask import Flask, jsonify

APP_NAME = os.getenv("APP_NAME", "tp-docker-python")
APP_VERSION = os.getenv("APP_VERSION", "1.0.0")
APP_ENV = os.getenv("APP_ENV", "dev")

app = Flask(__name__)
app.json.sort_keys = False  # conserve l'ordre des cles dans les reponses JSON


def _current_user() -> str:
    """Retourne 'nom (uid=...)' de l'utilisateur qui execute le processus."""
    uid = os.getuid()
    try:
        import pwd

        name = pwd.getpwuid(uid).pw_name
    except (ImportError, KeyError):
        name = "unknown"
    return f"{name} (uid={uid})"


@app.get("/")
def index():
    return jsonify(
        message="Bienvenue sur l'API Python Flask du TP Docker",
        application=APP_NAME,
        endpoints=["/", "/health", "/api/info", "/api/time"],
    )


@app.get("/health")
def health():
    return jsonify(status="UP")


@app.get("/api/info")
def info():
    return jsonify(
        application=APP_NAME,
        version=APP_VERSION,
        environment=APP_ENV,
        language="Python",
        runtime=f"{platform.python_implementation()} {platform.python_version()}",
        hostname=socket.gethostname(),
        platform=f"{platform.system()}/{platform.machine()}",
        user=_current_user(),
        timestamp=datetime.now(timezone.utc).isoformat(),
    )


@app.get("/api/time")
def server_time():
    now_utc = datetime.now(timezone.utc)
    local = now_utc.astimezone()
    return jsonify(
        utc=now_utc.isoformat(),
        local=local.isoformat(),
        timezone=str(local.tzinfo),
        epoch=int(time.time()),
    )


if __name__ == "__main__":
    # Serveur de developpement Flask (methode A uniquement).
    # En production, l'application est servie par Gunicorn (methodes B et C).
    app.run(host="0.0.0.0", port=int(os.getenv("PORT", "5000")))
