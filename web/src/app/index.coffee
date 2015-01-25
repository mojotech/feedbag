angular.module "feedBag", ['ngAnimate', 'ngSanitize', 'ngRoute', 'ngMaterial']
  .config ($routeProvider) ->
    $routeProvider
      .when "/",
        templateUrl: "app/main/main.html"
        controller: "MainCtrl"
      .otherwise
        redirectTo: "/"

