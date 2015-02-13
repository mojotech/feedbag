class ActivityModel extends Backbone.Model
  idAttribute: "id"

class ActivityCollection extends Backbone.Collection
  model: ActivityModel
  url: "api/activity"
  comparator: (model) -> new Date(model.get("event_time")).getTime() * -1

module.exports = new ActivityCollection
