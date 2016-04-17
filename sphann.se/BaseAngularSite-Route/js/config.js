var testApp = angular.module('testApp', ['ngRoute', 'testAppControllers']);

testApp.config(function($routeProvider) {
	$routeProvider.
		when('/', {
			template: '<h1>Home</h1>',
			controller: 'StartCtrl'
		}).
		when('/hejsan', {
			template: '<h1>Tjosan</h1>',
			controller: 'HejsanCtrl'					
		}).
		otherwise ({
			redirectTo: '/'
		});
});