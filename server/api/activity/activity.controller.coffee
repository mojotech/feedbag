"use strict"

_ = require "lodash"
Activity = require "./activity.model"
User = require "../user/user.model"
rabbit = require '../../components/rabbitmq'

###
Root Events Endpoint
###
exports.index = (req, res) ->
  activity = Activity.find {}

  activity.sort "-date"

  activity.populate "user"

  activity.exec (err, activity) ->
    res.json activity
