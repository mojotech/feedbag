###
Main application file
###

'use strict'

# Set default node environment to development
process.env.NODE_ENV = process.env.NODE_ENV or 'development'

express = require 'express'
mongoose = require 'mongoose'
rabbit = require './components/rabbitmq'
config = require './config/environment'
path = require 'path'
errors = require './components/errors'

# Connect to database
mongoose.connect config.mongo.uri, config.mongo.options

rabbit.connect config.rabbit.uri

# Setup server
app = express()
server = require('http').createServer(app)
socketio = require('socket.io') server,
  serveClient: !(config.env is 'production')
  path: '/socket.io-client'


require('./config/socketio')(socketio)
require('./config/express')(app)

# Log all requests to server
require('./components/logger').RoutesInfo(app)

config.getGlobbedFiles("./server/routes/**/*").forEach (routePath) ->
  require(path.resolve(routePath)) app

# Log error routes
require('./components/logger').RoutesError(app)

# Final catchall for errors
app.use errors[500]

# Start server
server.listen(config.port, config.ip, () ->
  console.log "Express server listening on #{config.port}, in #{app.get('env')} mode"
)

# Expose app
exports = module.exports = app
