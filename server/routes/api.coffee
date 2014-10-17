###
Main application routes
###
"use strict"

errors = require("../components/errors")

module.exports = (app) ->

  # Insert routes below
  app.use "/api/users", require("../api/user")
  app.use "/api/events", require("../api/event")
  app.use "/api/activity", require("../api/activity")
