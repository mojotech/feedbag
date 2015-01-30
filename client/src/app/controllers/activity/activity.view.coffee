templates = require("../templates")

class ItemView extends Backbone.View
  template: => templates.get(@model.get("template_id"))
  className: "widget"

  initialize: -> @render()

  render: ->
    @$el.html(@template()(@model.get("event")))
    this

module.exports = ItemView
