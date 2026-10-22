'use strict';
// Tests integres a Node.js (node:test) : aucune dependance supplementaire.
const { test, before, after } = require('node:test');
const assert = require('node:assert/strict');
const app = require('../server');

let server;
let baseUrl;

before(async () => {
  await new Promise((resolve) => {
    server = app.listen(0, '127.0.0.1', resolve);
  });
  baseUrl = `http://127.0.0.1:${server.address().port}`;
});

after(() => new Promise((resolve) => server.close(resolve)));

test('GET /health retourne UP', async () => {
  const res = await fetch(`${baseUrl}/health`);
  assert.equal(res.status, 200);
  assert.deepEqual(await res.json(), { status: 'UP' });
});

test('GET /api/version retourne la version du package', async () => {
  const res = await fetch(`${baseUrl}/api/version`);
  assert.equal(res.status, 200);
  const body = await res.json();
  assert.equal(body.version, '1.0.0');
});

test('GET /inconnu retourne 404', async () => {
  const res = await fetch(`${baseUrl}/inconnu`);
  assert.equal(res.status, 404);
});
