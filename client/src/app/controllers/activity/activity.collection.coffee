
class ActivityModel extends Backbone.Model

class ActivityCollection extends Backbone.Collection
  url: "api/activity"
  model: ActivityModel

module.exports = new ActivityCollection
