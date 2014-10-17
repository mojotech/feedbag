###
  Module:       Feedbag Templates
  Description:  Templating system for feedbag activity items. Uses a hash map
                to dynamically load in the template required by the activity
                event.
###


app = angular.module("feedbag.templates", [])
app.directive "item", ($compile, $http, $templateCache) ->

  templateHash =
    banner: "components/templates/banner.html"
    blank: "components/templates/blank.html"

  linker = (scope, element, attrs) ->
    template = $templateCache.get(templateHash[scope.data.template])

    $http.get(templateHash[scope.data.template], cache: $templateCache).success (html) ->
      element.html(html).show()
      $compile(element.contents()) scope

  restrict: "E"
  link: linker
  transclude: true
  scope:
    data: "="
