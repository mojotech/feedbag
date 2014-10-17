'use strict'

angular.module 'FeedBagApp'
.factory 'Activity', ($resource) ->
  $resource '/api/activity/:id', id: '@_id'

