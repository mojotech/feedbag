###
Express configuration
###

"use strict"

express = require("express")
favicon = require("static-favicon")
morgan = require("morgan")
compression = require("compression")
bodyParser = require("body-parser")
methodOverride = require("method-override")
cookieParser = require("cookie-parser")
errorHandler = require("errorhandler")
path = require("path")
config = require("./environment")
passport = require("passport")

module.exports = (app) ->

  env = app.get("env")

  app.set "views", config.root + "/server/views"
  app.set "view engine", "jade"

  app.use compression()
  app.use bodyParser.urlencoded(extended: false)
  app.use bodyParser.json()
  app.use methodOverride()
  app.use cookieParser()
  app.use passport.initialize()

  if "production" is env
    app.use favicon(path.join(config.root, "public", "favicon.ico"))
    app.use express.static(path.join(config.root, "public"))
    app.set "appPath", config.root + "/public"
    app.use morgan("dev")

  if "development" is env or "test" is env
    app.use favicon(path.join(config.root, "client", "favicon.ico"))
    app.use require("connect-livereload")()
    app.use express.static(path.join(config.root, ".tmp"))
    app.use express.static(path.join(config.root, "client"))
    app.set "appPath", "client"
    app.use morgan("dev")
    app.use errorHandler() # Error handler - has to be last
