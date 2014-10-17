mongoose = require("mongoose")
Schema = mongoose.Schema

Creation = (schema, options) ->
  options = options or {}

  schema.add created_at:
    type: Date
    default: Date.now

  schema.add updated_at:
    type: Date
    default: Date.now

  schema.set "toJSON",
    virtuals: false
    transform: (user, json, options) ->

      #Global Transforms on defaults
      delete json.__v

      json

  schema.pre "save", (next) ->
    self = this
    self.updated_at = new Date()
    next()

  schema.pre "remove", (next) ->

    #TODO: Remove all references if has any child documents.
    next()

module.exports = exports = Creation
