var testApp = angular.module('testApp', ['ngRoute', 'testAppControllers']);
var testAppControllers = angular.module('testAppControllers', []);


testAppControllers.controller('StartCtrl', ['$scope', '$http',
  function ($scope, $http) {
  	alert('hejsan');
}]);

testAppControllers.controller('HejsanCtrl', ['$scope', '$http',
  function ($scope, $http) {
  	alert('hejssdfsdfsdfsdan');
}]);
