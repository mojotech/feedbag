activityCollection = require("../activity/activity.collection")
activityItem = require("../activity/activity.view")
templates = require("../templates")

class AppView extends Backbone.Marionette.CollectionView
  el: "#feedbag-app"
  childView: activityItem

module.exports = new AppView(collection: activityCollection)
