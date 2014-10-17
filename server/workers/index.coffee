throng = require 'throng'
config = require '../config/environment'

Worker = require './github'

workerOpts =
  scriptsDir: "server/scripts" # from the server directory root
  mongoUri: config.mongo.uri
  rabbitUri: config.rabbit.uri
  debug: false


# Instanciate new worker child processes
throng ->
  new Worker(workerOpts)
,
  workers: null
  lifetime: Infinity
