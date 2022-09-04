const { Router } = require('express')

const testRouter = Router()

testRouter.get('/gwak', (req, res) => {
  console.log(req)
  res.json({
    message: 'test',
  })
})

module.exports = testRouter
