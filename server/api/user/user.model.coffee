"use strict"

mongoose = require("mongoose")
Schema = mongoose.Schema

CreationPlugin = require("../../components/plugins/creation.plugin")

UserSchema = new Schema
  name:
    type: String
    trim: true

  email:
    type: String
    lowercase: true
    trim: true

  role:
    type: String
    default: "user"

  gravatarHash:
    type: String

  last_modified_event:
    type: Date

  githubAccessToken:
    type: String

  github:
    id:
      type: String

    username:
      type: String

    access_token:
      type: String

    events:
      last_modified:
        type: Date

    raw: {}

###
Virtuals
###
# Non-sensitive info we'll be putting in the token
UserSchema.virtual("token").get ->
  _id: @_id
  role: @role

###
Pre-save hook
###
UserSchema.pre "save", (next) ->
  self = this
  self.gravatarHash = crypto.createHash("md5").update(self.email).digest("hex")
  next()

UserSchema.plugin CreationPlugin
module.exports = mongoose.model("User", UserSchema)
