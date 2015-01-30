"use strict"

layoutView = require('./controllers/layout')
templates = require('./controllers/templates')
activity = require('./controllers/activity/activity.collection')

class Feedbag
  constructor: -> @initialize()

  initialize: ->
    req = templates.fetch()
    req.success => @render()
    req.fail -> console.error "Failed to load templates"

  recentActivity: ->
    activity.fetch()

  render: ->
    @recentActivity()
    layoutView.render()

window.Feedbag = new Feedbag()
