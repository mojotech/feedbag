###
Broadcast updates to client when the model changes
###
"use strict"

activity = require("./activity.model")
rabbit = require '../../components/rabbitmq'

onSave = (socket, doc, cb) ->
  socket.emit "activity:save", doc

onRemove = (socket, doc, cb) ->
  socket.emit "activity:remove", doc

exports.register = (socket) ->

  # Create handler for new events. This activity is called
  # every time a unique activity is added to the events
  # collection.
  rabbit.server.handle "activity.create", (job, ack) ->
    onSave socket, job.doc
    ack()


  activity.schema.post "save", (doc) ->
    onSave socket, doc

  activity.schema.post "remove", (doc) ->
    onRemove socket, doc
