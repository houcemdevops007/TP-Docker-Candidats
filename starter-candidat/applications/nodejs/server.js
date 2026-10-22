'use strict';
/**
 * TP Docker DevSecOps - Application 3 : API REST Node.js / Express
 * Auteur : Khlifi Houcem - Formateur DevSecOps et Cloud
 *
 * Endpoints :
 *   GET /             message de bienvenue
 *   GET /health       etat de sante
 *   GET /api/info     informations d'execution
 *   GET /api/version  version applicative
 */
const os = require('node:os');
const express = require('express');
const pkg = require('./package.json');

const PORT = Number.parseInt(process.env.PORT || '3000', 10);
const APP_NAME = process.env.APP_NAME || pkg.name;
const APP_VERSION = process.env.APP_VERSION || pkg.version;
const APP_ENV = process.env.APP_ENV || process.env.NODE_ENV || 'development';

const app = express();
app.disable('x-powered-by'); // ne pas divulguer la technologie serveur

app.get('/', (req, res) => {
  res.json({
    message: "Bienvenue sur l'API Node.js Express du TP Docker",
    application: APP_NAME,
    endpoints: ['/', '/health', '/api/info', '/api/version'],
  });
});

app.get('/health', (req, res) => {
  res.json({ status: 'UP' });
});

app.get('/api/info', (req, res) => {
  const { username, uid } = os.userInfo();
  res.json({
    application: APP_NAME,
    version: APP_VERSION,
    environment: APP_ENV,
    language: 'JavaScript',
    runtime: `Node.js ${process.version}`,
    hostname: os.hostname(),
    platform: `${process.platform}/${process.arch}`,
    user: `${username} (uid=${uid})`,
    timestamp: new Date().toISOString(),
  });
});

app.get('/api/version', (req, res) => {
  res.json({ application: APP_NAME, version: APP_VERSION });
});

// 404 au format JSON pour toute autre route
app.use((req, res) => {
  res.status(404).json({ error: 'Not Found', path: req.path });
});

module.exports = app;

// Demarrage du serveur uniquement si le fichier est execute directement
if (require.main === module) {
  const server = app.listen(PORT, '0.0.0.0', () => {
    console.log(`[tp-docker-nodejs] ecoute sur 0.0.0.0:${PORT} (env=${APP_ENV})`);
  });

  // Arret propre : docker stop envoie SIGTERM au PID 1
  const shutdown = (signal) => {
    console.log(`[tp-docker-nodejs] ${signal} recu, arret en cours...`);
    server.close(() => process.exit(0));
    setTimeout(() => process.exit(1), 10000).unref();
  };
  process.on('SIGTERM', () => shutdown('SIGTERM'));
  process.on('SIGINT', () => shutdown('SIGINT'));
}
