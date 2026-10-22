# TP Docker — Création et optimisation d'images · Kit de démarrage candidat

**DOCUMENT CANDIDAT — TRAVAUX PRATIQUES**
Khlifi Houcem — Formateur DevSecOps et Cloud · Version 1.0 · 22/10/2026

Ce dossier contient **uniquement** les fichiers dont vous avez besoin pour démarrer.
Les consignes détaillées se trouvent dans `TP_DOCKER_ENONCE.html` (à ouvrir dans un navigateur).

## Contenu

```text
starter-candidat/
├── applications/
│   ├── java/      pom.xml, src/ (code fourni)      + Dockerfile.* et .dockerignore à compléter
│   ├── python/    app.py, requirements.txt         + Dockerfile.* et .dockerignore à compléter
│   ├── nodejs/    server.js, package*.json, test/  + Dockerfile.* et .dockerignore à compléter
│   └── golang/    main.go, go.mod                  + Dockerfile.* et .dockerignore à compléter
├── challenges/
│   ├── 9.2-dockerfile-a-corriger/   Dockerfile.bad à analyser et corriger
│   └── 9.6-conteneur-qui-s-arrete/  service qui refuse de démarrer
└── scripts/
    └── scan-images.sh               squelette du script de scan Trivy (challenge 9.8)
```

## Règles

- Le **code applicatif est fourni et fonctionnel** : votre travail porte sur la conteneurisation.
- Chaque fichier `Dockerfile.*` contient des blocs `TODO` : remplacez-les par les instructions adaptées.
- N'utilisez **jamais** le tag `latest` ; utilisez les tags imposés par l'énoncé.
- N'écrivez **aucun secret** (mot de passe, token) dans un Dockerfile.
- Ne lancez **pas** de commandes de suppression globale (`docker system prune -a`, `docker rm -f $(docker ps -aq)`…).
  Supprimez uniquement les ressources que vous avez créées (préfixe `tp-`, images `tp-docker/*`).

## Conventions

| Application | Port conteneur | Port hôte (exercices) | Images à produire |
|---|---|---|---|
| Java    | 8080 | 8181 → 8184 | `tp-docker/java:simple`, `:jar`, `:multistage`, `:optimized` |
| Python  | 5000 | 8281 → 8283 | `tp-docker/python:simple`, `:optimized`, `:multistage` |
| Node.js | 3000 | 8381 → 8383 | `tp-docker/nodejs:simple`, `:optimized`, `:multistage` |
| Go      | 8080 | 8481 → 8484 | `tp-docker/golang:simple`, `:slim`, `:scratch`, `:distroless` |

Ports du challenge « 4 applications simultanées » : 8081 (Java), 8082 (Python), 8083 (Node.js), 8084 (Go).

## Rendu

Depuis la page d'énoncé, utilisez le bouton **Exporter mes réponses (JSON)** et transmettez le fichier
obtenu au formateur, avec votre dossier `starter-candidat/` complété.
