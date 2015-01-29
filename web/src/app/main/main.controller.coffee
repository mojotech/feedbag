angular.module "feedBag"
  .controller "MainCtrl", ($scope, $http, socket) ->
    $scope.Activities = []

    activity = $http
      method: "get",
      url: "api/activity",

    activity.success (activities, status) ->
      $scope.Activities = activities

    socket.syncUpdates "activity", $scope.Activities
