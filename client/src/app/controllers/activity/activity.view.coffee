templates = require("../templates")

class ItemView extends Backbone.View
  template: => templates.get(@model.get("template_id"))

  className: ->
    "widget #{@template()?.size or 'small'}"

  initialize: ->
    @render()

  render: ->
    @$el.html(@template()?.template(@model.get("event")))
    this

module.exports = ItemView
