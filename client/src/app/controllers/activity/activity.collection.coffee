
class ActivityModel extends Backbone.Model

class ActivityCollection extends Backbone.Collection
  url: "api/activity"
  model: ActivityModel
  comparator: "created_at"

module.exports = new ActivityCollection
