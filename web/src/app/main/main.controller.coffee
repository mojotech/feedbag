angular.module "feedBag"
  .controller "MainCtrl", ($scope, $http, socket) ->
    $scope.Activities = []

    activity = $http
      method: "get",
      url: "api/activity",
      params:
        action: "get"

    activity.success (activities, status) ->
      $scope.Activities = activities

    socket.syncUpdates "activity", $scope.Activities
