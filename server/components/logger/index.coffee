'use strict'

winston = require "winston"
expressWinston = require "express-winston"
mongoWinston = require('winston-mongodb').MongoDB

config = require "../../config/environment"


module.exports.Logger = do ->
  logTransports = []

  logTransports.push new (winston.transports.Console)
      colorize: true

  logTransports.push new (winston.transports.MongoDB)
      db: config.mongo.name
      host: config.mongo.host
      collection: 'logs'
      safe: false
      port: config.mongo.port
      username: config.mongo.username
      passsword: config.mongo.password
      level: 'info'


  new (winston.Logger)
    transports: logTransports


module.exports.RoutesInfo = (app) ->
  # express winston request logger

  logTransports = []

  logTransports.push new (winston.transports.MongoDB)
      db: config.mongo.name
      host: config.mongo.host
      collection: 'requests'
      safe: false
      port: config.mongo.port
      username: config.mongo.username
      passsword: config.mongo.password
      level: 'info' # log all routes status 200+. Use 'warn' for status 400+

  app.use expressWinston.logger
    transports: logTransports
    meta: true
    msg: "{{req.method}} {{req.url}} {{res.statusCode}} {{res.responseTime}}ms"
    statusLevels: true


module.exports.RoutesError = (app) ->
  # express winston error logger

  logTransports = []

  logTransports.push new (winston.transports.MongoDB)
      db: config.mongo.name
      host: config.mongo.host
      collection: 'errors'
      safe: false
      port: config.mongo.port
      username: config.mongo.username
      passsword: config.mongo.password

  app.use expressWinston.errorLogger
    transports: logTransports
    meta: true
    msg: "{{req.method}} {{req.url}} {{res.statusCode}} {{res.responseTime}}ms"
