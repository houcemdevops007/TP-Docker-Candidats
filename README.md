# TP Docker — Création et optimisation d'images

**DOCUMENT CANDIDAT — TRAVAUX PRATIQUES**
Khlifi Houcem — Formateur DevSecOps et Cloud · Version 1.0 · 22/10/2026

Ce dépôt contient tout ce dont vous avez besoin pour réaliser le TP :

| Élément | Rôle |
|---|---|
| `TP_DOCKER_ENONCE.html` | Énoncé interactif : à ouvrir dans un navigateur (fonctionne hors ligne). |
| `starter-candidat/` | Kit de démarrage : code des 4 applications (Java, Python, Node.js, Go), Dockerfiles et `.dockerignore` à compléter (`TODO`), challenges, squelette de script. |

## Démarrage

```bash
git clone https://github.com/houcemdevops007/TP-Docker-Candidats.git ~/tp-docker
cd ~/tp-docker
open TP_DOCKER_ENONCE.html        # macOS  (Linux : xdg-open, Windows : start)
cd starter-candidat
```

Suivez ensuite l'énoncé à partir du chapitre 1. Le détail du kit et des conventions
(ports, noms d'images) se trouve dans [`starter-candidat/README.md`](starter-candidat/README.md).

## Prérequis

- Docker Engine 24+ ou Docker Desktop récent, avec BuildKit (`docker buildx version`).
- Accès à Docker Hub et à `gcr.io` (images de base), environ 10 Go d'espace disque libre.
- `curl`, `git` et un navigateur récent.

## Règles

- Aucun tag `latest`, aucun secret dans les Dockerfiles.
- Aucune commande de suppression globale (`docker system prune -a`…) : nettoyez
  uniquement vos ressources (préfixe `tp-`, images `tp-docker/*`).

## Rendu

En fin de séance, exportez vos réponses depuis l'énoncé (bouton **Exporter mes réponses (JSON)**)
et remettez ce fichier au formateur avec votre dossier `starter-candidat/` complété
(archive ou dépôt Git personnel, selon la consigne donnée en séance).
