const app = require('express')()
const { loadNuxt, build } = require('nuxt')
// Import and Set Nuxt.js options
// let config = require('../nuxt.config.ts')
const isDev = process.env.NODE_ENV !== 'production'
const apiRouter = require('./middle')

/* eslint-disable */

async function start() {
  // Init Nuxt.js
  const nuxt = await loadNuxt(isDev ? 'dev' : 'start')

  const { host, port } = nuxt.options.server

  app.use((err, req, res, next) => {
    console.error(err)
    log.error(err.stack)
    return res.status(err.statusCode || 500).send(err.message)
  })
  app.use('/api', apiRouter)
  // Render every route with Nuxt
  app.use(nuxt.render)

  // Build only in dev mode with hot-reloading
  if (isDev) {
    build(nuxt)
  }
  // Listen the server
  app.listen(port, host)
  // Listen the server
  console.log(`Server listening on http://${host}:${port}`)
}

start()
