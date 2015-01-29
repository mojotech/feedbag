angular.module "feedBag", [
  'ngAnimate',
  'ngRoute',
  'ui.bootstrap',
  'btford.socket-io'
]
  .config ($routeProvider) ->
    $routeProvider
      .when "/",
        templateUrl: "app/main/main.html"
        controller: "MainCtrl"
      .otherwise
        redirectTo: "/"

angular.module("feedBag").run ($http, $templateCache) ->
  templates = $http
    method: "get"
    url: "templates"

  templates.success (templates) ->
    templates.forEach (res) ->
      $templateCache.put(res.id, res.template)
