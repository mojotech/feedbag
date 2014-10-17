'use strict'

angular.module 'FeedBagApp'
.config ($stateProvider) ->
  $stateProvider
  .state 'login',
    url: '/login'
    templateUrl: 'app/profile/login/login.html'
    controller: 'LoginCtrl'
