"use strict"

_ = require "lodash"
async = require 'async'
Event = require "./event.model"
Activity = require "../activity/activity.model"
User = require "../user/user.model"
rabbit = require '../../components/rabbitmq'

###
Root Events Endpoint
###
exports.index = (req, res) ->
  events = Event.find {}

  events.sort "-created_at"

  events.populate "user"

  events.exec (err, events) ->
    res.json(events)


###
Clear Activity and Events
###
exports.clear = (req, res) ->

  removeFunc = (item, cb) -> item.remove(cb)

  async.parallel
    clearEvents: (cb) ->
      Event.find {}, (err, events) ->
        async.each events, removeFunc, cb

    clearActivity:(cb) ->
      Activity.find {}, (err, activity) ->
        async.each activity, removeFunc, cb

  , (err, results) ->
    return res.status(400).json(error: "problem clearing database") if err
    res.json task: "complete"

###
Test RabbitMQ
###
exports.test_populate_events = (req, res) ->

  User.find {}, (err, users) ->
    # Find all users and publish a github fetch job for them
    # RabbitMQ will take care of the rest.
    _.each users, (user) ->
      console.log "Schedualing job for #{user.email}"
      rabbit.server.publish "jobs.github", user_id: user._id

  res.json
    status: "I think it worked. Go Check!"
