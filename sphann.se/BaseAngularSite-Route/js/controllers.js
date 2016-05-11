var testApp = angular.module('testApp', ['ngRoute', 'testAppControllers']);
var testAppControllers = angular.module('testAppControllers', []);

var recepts = [];
var recept = {
	name: "Sill",
	notes: "Note"
}
recepts[] = recept;

testAppControllers.controller('StartCtrl', ['$scope', '$http',
  function ($scope, $http) {
  	alert('hejsan');
  	$scope.recept = recept;
}]);

testAppControllers.controller('HejsanCtrl', ['$scope', '$http',
  function ($scope, $http) {
  	alert('hejssdfsdfsdfsdan');
}]);

