"use strict"

angular.module("FeedBagApp")
.controller "ActivityCtrl", ($scope, $http, socket, Activity) ->

  $scope.Activity = []

  Activity.query (Activity) ->
    $scope.Activity = Activity
    socket.syncUpdates "activity", $scope.Activity

  $scope.$on "$destroy", ->
    socket.unsyncUpdates "activity"
