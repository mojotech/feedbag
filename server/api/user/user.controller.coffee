"use strict"

User = require("./user.model")
passport = require("passport")
config = require("../../config/environment")
jwt = require("jsonwebtoken")

validationError = (res, err) -> res.json 422, err


###
Get list of users
###
exports.index = (req, res) ->
  User.find {}, (err, users) ->
    return res.send(500, err)  if err
    res.json 200, users


###
Get a single user
###
exports.show = (req, res, next) ->
  userId = req.params.id
  User.findById userId, (err, user) ->
    return next(err)  if err
    return res.send(401)  unless user
    res.json user.profile

###
Deletes a user
###
exports.destroy = (req, res) ->
  # TODO: build out to remove token from github app
  User.findByIdAndRemove req.params.id, (err, user) ->
    return res.send(500, err)  if err
    res.send 204


###
Get my info
###
exports.me = (req, res, next) ->
  userId = req.user._id
  User.findOne
    _id: userId
  , (err, user) ->
    return next(err)  if err
    return res.json(401)  unless user
    res.json user


###
Authentication callback
###
exports.authCallback = (req, res, next) -> res.redirect "/"
