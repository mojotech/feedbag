'use strict'

angular.module 'FeedBagApp'
.controller 'LoginCtrl', ($scope, Auth, $location, $window) ->
  $scope.loginOauth = (provider) ->
    $window.location.href = '/auth/' + provider
