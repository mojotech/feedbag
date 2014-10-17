"use strict"
angular.module("FeedBagApp").controller "NavbarCtrl", ($scope, $location, Auth, $window) ->

  $scope.isCollapsed = true
  $scope.isLoggedIn = Auth.isLoggedIn
  $scope.isAdmin = Auth.isAdmin

  $scope.getCurrentUser = Auth.getCurrentUser

  $scope.logout = ->
    Auth.logout()
    $location.path "/login"

  $scope.isActive = (route) -> route is $location.path()

  $scope.loginOauth = (provider) -> $window.location.href = '/auth/' + provider


