"use strict"

mongoose = require("mongoose")
Schema = mongoose.Schema

CreationPlugin = require("../../components/plugins/creation.plugin")

ActivitySchema = new Schema

  name:
    type: String

  template:
    type: String

  description:
    type: String

  icon:
    type: String

  date:
    type: Date

  body:
    type: String

  user:
    type: Schema.ObjectId
    ref: "User"

ActivitySchema.plugin CreationPlugin

module.exports = mongoose.model("Activity", ActivitySchema)
