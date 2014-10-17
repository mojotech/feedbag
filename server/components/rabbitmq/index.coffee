"use strict"

jackrabbit = require 'jackrabbit'
async = require 'async'
logger = require('../logger').Logger

module.exports =

  # Where rabbit server instance will live.
  server: undefined

  connect: (uri) ->
    # Establish connection with Rabbitmq via jackrabbit module.
    @server = jackrabbit(uri)

    # On connected callback, setup connection with all rooms.
    @server.on "connected", =>

      # Kickoff room creation in parallel and log any errors
      # to the logger service.
      async.parallel
        github: (cb) =>
          @server.create "jobs.github", cb

        events: (cb) =>
          @server.create "events.create", cb

        activity: (cb) =>
          @server.create "activity.create", cb

      , (err, results) ->
          logger.error "Error creating rabbitmq rooms", err if err


