"""
Challenge 9.6 - Service "tp-crash" : petit service HTTP (bibliotheque standard).
Ce code est CORRECT : ne le modifiez pas. Le probleme se situe dans la facon
dont l'image est construite et dont le conteneur est lance.
"""
import json
import os
import sys
from datetime import datetime, timezone
from http.server import BaseHTTPRequestHandler, HTTPServer

SECRET = os.environ.get("APP_SECRET_KEY")
if not SECRET:
    print("ERREUR FATALE : la variable d'environnement APP_SECRET_KEY est absente", file=sys.stderr)
    sys.exit(3)

PORT = int(os.environ.get("PORT", "8000"))
STATE_DIR = os.environ.get("STATE_DIR", "/srv/state")

# Ecriture d'un fichier d'etat au demarrage
os.makedirs(STATE_DIR, exist_ok=True)
with open(os.path.join(STATE_DIR, "started.json"), "w", encoding="utf-8") as f:
    json.dump({"started_at": datetime.now(timezone.utc).isoformat(), "pid": os.getpid()}, f)


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/health":
            body, code = {"status": "UP"}, 200
        else:
            body, code = {"service": "tp-crash", "message": "Diagnostic reussi !"}, 200
        payload = json.dumps(body).encode()
        self.send_response(code)
        self.send_header("Content-Type", "application/json")
        self.send_header("Content-Length", str(len(payload)))
        self.end_headers()
        self.wfile.write(payload)


print(f"tp-crash ecoute sur 0.0.0.0:{PORT}", flush=True)
HTTPServer(("0.0.0.0", PORT), Handler).serve_forever()
