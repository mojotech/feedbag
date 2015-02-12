class ActivityModel extends Backbone.Model
  idAttribute: "id"

class ActivityCollection extends Backbone.Collection
  model: ActivityModel
  url: "api/activity"
  comparator: (model) -> -1 * model.get('id')

module.exports = new ActivityCollection
