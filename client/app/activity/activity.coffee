"use strict"
angular.module("FeedBagApp").config ($stateProvider) ->
  $stateProvider.state "activity",
    url: "/"
    templateUrl: "app/activity/activity.html"
    controller: "ActivityCtrl"

