"use strict"

angular.module("FeedBagApp", [
  "ngCookies"
  "ngResource"
  "ngSanitize"
  "ngAnimate"
  "ui.bootstrap"
  "btford.socket-io"
  "ui.router"
  "ui.scrollfix"
  "feedbag.templates"
]).config(($stateProvider, $urlRouterProvider, $locationProvider, $httpProvider) ->
  $urlRouterProvider.otherwise "/"
  $locationProvider.html5Mode true
  $httpProvider.interceptors.push "authInterceptor"

).factory("authInterceptor", ($rootScope, $q, $cookieStore, $location) ->
  request: (config) ->
    config.headers = config.headers or {}
    config.headers.Authorization = "Bearer " + $cookieStore.get("token")  if $cookieStore.get("token")
    config

  responseError: (response) ->
    if response.status is 401
      $location.path "/login"
      $cookieStore.remove "token"
      $q.reject response
    else
      $q.reject response

).run ($rootScope, $location, Auth) ->

  # Redirect to login if route requires auth and you're not logged in
  $rootScope.$on "$stateChangeStart", (event, next) ->
    $location.path "/login"  if next.authenticate and not Auth.isLoggedIn()
