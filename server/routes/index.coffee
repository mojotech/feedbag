###
Main application routes
###
"use strict"

errors = require "../components/errors"

module.exports = (app) ->
  app.use "/auth", require("../auth")

  # All undefined asset or api routes should return a 404
  app.route("/:url(api|auth|components|app|bower_components|assets)/*").get errors[404]

  # All other routes should redirect to the index.html
  app.route("/*").get (req, res) -> res.sendfile app.get("appPath") + "/index.html"


  # Use middleware to parse and handle spacific routing errors
  app.use errors[401]
