const express = require('express')

const app = express()
const { loadNuxt, build } = require('nuxt')
const log4js = require('log4js')

// process.env.DEBUG = 'nuxt:*'
// Import and Set Nuxt.js options
// let config = require('../nuxt.config.ts')
const isDev = process.env.NODE_ENV !== 'production'

const { systemLogger } = require('./logger')

app.use(log4js.connectLogger(systemLogger, { level: 'auto' }))

/* eslint-disable */

async function start() {
  // Init Nuxt.js
  const nuxt = await loadNuxt(isDev ? 'dev' : 'start')

  const { host, port } = nuxt.options.server

  app.use(express.json()) // req.body에 json형식으로 저장하기 위해
  app.post('/api/error', (req, res) => {
    systemLogger.error(req.body)
    res.json({ message: 'Success Error log' })
  })
  // Render every route with Nuxt
  app.use(nuxt.render)

  // Build only in dev mode with hot-reloading
  if (isDev) {
    build(nuxt)
  }
  // Listen the server
  app.listen(port, host)
  // Listen the server
  systemLogger.info(`Server listening on http://${host}:${port}`)
}

start()
