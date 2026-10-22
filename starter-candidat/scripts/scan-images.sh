#!/usr/bin/env bash
# =============================================================================
# Challenge 9.8 - Scanner automatiquement toutes les images du TP avec Trivy
# Squelette a completer par le candidat.
# =============================================================================
set -euo pipefail

# TODO 1 : variables : dossier de sortie (reports/trivy), severites (HIGH,CRITICAL),
#          version de l'image Trivy si le binaire n'est pas installe

# TODO 2 : verifier que Docker repond, sinon quitter avec un message clair

# TODO 3 : definir une fonction qui appelle "trivy" si installe,
#          sinon l'image aquasec/trivy (socket Docker + cache montes)

# TODO 4 : lister toutes les images tp-docker/* presentes localement

# TODO 5 : pour chaque image :
#            - rapport TXT filtre HIGH,CRITICAL
#            - rapport JSON complet
#            - compter CRITICAL et HIGH (jq ou python3)
#            - recuperer l'utilisateur configure (docker image inspect)

# TODO 6 : produire un fichier summary.csv (image,user,critical,high)

# TODO 7 : code retour non nul si au moins une CVE CRITICAL est detectee
#          (option activable par variable d'environnement)
