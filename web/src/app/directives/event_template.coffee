angular.module("feedBag").directive "widget", ($compile, $http, $templateCache) ->

  linker = (scope, element, attrs) ->
    angular.extend(scope, scope.data.event)
    template = $templateCache.get(scope.data.template_id)
    element.html(template)
    $compile(element.contents())(scope)

  restrict: "E"
  link: linker
  transclude: true
  scope:
    data: "="
