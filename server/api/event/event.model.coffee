"use strict"

mongoose = require("mongoose")
Schema = mongoose.Schema

EventSchema = new Schema

  github_id:
    type: String
    unique: true

  created_at:
    type: Date

  user:
    type: Schema.ObjectId
    ref: "User"

  type:
    type: String

  repo:
    name:
      type: String
    url:
      type: String

  payload: {}

module.exports = mongoose.model("Event", EventSchema)
