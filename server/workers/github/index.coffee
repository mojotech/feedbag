_ = require 'lodash'
async = require 'async'
path = require 'path'
mongoose = require 'mongoose'
throng = require 'throng'
jackrabbit = require 'jackrabbit'
GitHubApi = require 'github'
{EventEmitter} = require 'events'

Config = require '../../config/environment'
Logger = require('../../components/logger').Logger

# Mongoose Models
User = require "../../api/user/user.model"
Event = require "../../api/event/event.model"
Activity = require "../../api/activity/activity.model"


class Worker extends EventEmitter

  constructor: (opts) ->
    {@scriptsDir, @rabbitUri, @mongoUri, @debug} = opts

    @rabbitRooms = ["jobs.github", "events.create", "activity.create"]
    @github = null
    @rabbit = null
    @Users = null
    @Activity = null

    @initializer()

    console.log "Spinning up worker #{process.pid}"

    process.on 'SIGTERM', @handleProcessDeath
    process.on 'SIGINT', @handleProcessDeath
    process.on 'SIGQUIT', @handleProcessDeath



  # Set listeners and kickoff processing when everything is ready.
  initializer: ->
    @on "rabbit:connected", @establishRabbitRooms
    @on "rabbit:roomsCreated", @setHandlers

    # Connect to required services
    @connectGithub()
    @connectMongoose(@mongoUri)
    @connectRabbit(@rabbitUri)
    @loadActivityScripts(@scriptsDir)



  processJob: (job, ack) =>
    async.waterfall [
      async.apply(@findUser, job.user_id)
      @setGithubAuth,
      @getGithubEvents,
      @createEvents,
      @processActivity,
      @saveActivity
    ], (err, results) ->
      Logger.error "Issue processing github worker job", err if err

      console.log "Finished processing for #{job.user_id}"

      # Call the ack when all methods have finished processing
      ack()


  findUser: (id, cb) -> User.findById id, cb

  setGithubAuth: (user, cb) =>
    # Set the github authentication headers
    @github.authenticate
      type: "oauth"
      token: user.github.access_token

    # pass on the user to the next waterfall function
    cb(null, user)

  getGithubEvents: (user, cb) =>
    githubFetchOptions =
      user: user.github.username
      # headers:
        # 'If-Modified-Since': new Date(user.github.events.last_modified).toGMTString() if user.github.events.last_modified?

    @github.events.getFromUser githubFetchOptions, (err, events) ->
      Logger.error "Could not fetch github events", err if err

      err = new Error "No Events from Github" unless events?.length

      # pass the user and events to the next method
      cb(err, user, events)

  createEvents: (user, events, cb) ->
    eventMap = _.map events, (githubEvent) ->
      new Event
        github_id: githubEvent.id
        user: user._id
        type: githubEvent.type
        repo: githubEvent.repo
        created_at: githubEvent.created_at
        payload: githubEvent.payload

    saveFunc = (event, asyncCB) -> event.save (err, event) ->
        return asyncCB(null, event) unless err

        # Call the callback even if there was an error to continue the
        # processing
        asyncCB()

    async.map eventMap, saveFunc, (err, events) ->
      Logger.error "Issue while saving the events", err if err

      events = _.compact events
      cb(err, user, events)



  processActivity: (user, events, cb) =>
    opts =
      user: user
      events: events
      Event: Event

    scriptPrep = _.map @scripts, (script) -> async.apply(script, opts)

    async.parallel scriptPrep, (err, activityArray) ->
      Logger.error "An activity script returned an error", err if err

      activityArray = _.compact _.flatten activityArray

      if activityArray?.length
        activityMap = _.map activityArray, (activity) ->
          # set the defauls for new activity events
          _.defaults activity,
            template: "banner"
            description: ""
            icon: "fa-ban"
            body: ""
            user: user._id

          new Activity activity

      cb(null, user, activityMap)


  saveActivity: (user, activityMap, cb) =>
    return cb(null) unless activityMap?.length

    saveFunc = (activity, asyncCB) => activity.save (err, activity) =>
      @rabbit.publish "activity.create", doc: activity unless err
      asyncCB(err, activity)

    async.each activityMap, saveFunc, (err, activities) ->
      Logger.error "Problem saving the activities", err if err
      cb()




  establishRabbitRooms: =>
    createRooms = (roomName, cb) => @rabbit.create roomName, cb
    async.each @rabbitRooms, createRooms, (err, results) =>
      Logger.error "Worker couldn't connect to RabbitMQ room", err if err
      @emit "rabbit:roomsCreated"


  setHandlers: -> @rabbit.handle "jobs.github", @processJob


  loadActivityScripts: (scripts) ->
    @scripts = _.map Config.getGlobbedFiles(scripts + "/*"), (routePath) ->
      require(path.resolve(routePath))


  connectGithub: ->
    @github = new GitHubApi
      version: "3.0.0"
      debug: @debug or false
      protocol: "https"


  connectMongoose: (uri) ->
    mongoose.connect uri


  connectRabbit: (uri) ->
    @rabbit = jackrabbit(uri).on "connected", => @emit "rabbit:connected"


  handleProcessDeath: ->
    console.log "Processing going down..."
    process.exit()



module.exports = Worker
