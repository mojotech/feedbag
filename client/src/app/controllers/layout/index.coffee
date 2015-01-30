activityCollection = require("../activity/activity.collection")
activityItem = require("../activity/activity.view")
templates = require("../templates")

class AppView extends Backbone.View
  el: "#feedbag-app"
  childView: activityItem
  _viewPointers: {}

  renderCollection: ->
    els = _.map @_viewPointers, (view) -> view.el
    $(els)

  rendersCollectively: ->
    @_viewPointers = {}
    @collection.each @addOne

  addOne: (model) ->
    view = new @childView(model: model)
    @_viewPointers[model.cid] = view

  removeOne: (model) ->
    @_viewPointers[model.cid].remove()
    delete @_viewPointers[model.cid]

  initialize: ({@collection}) ->
    @rendersCollectively()

    @listenTo @collection, "reset", =>
      @rendersCollectively()
      @render()

    @listenTo @collection, "add", (model) =>
      @addOne(model)
      @render()

    @listenTo @collection, "remove", (model) =>
      @removeOne(model)
      @render()

  render: ->
    @$el.html(@renderCollection())
    this

module.exports = new AppView(collection: activityCollection)
