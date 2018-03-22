var myApp = angular.module('restCall', []);
var myApp = angular.module('restCall2', []);


angular.module('restCall', [])
.controller('CallController', function($scope, $http) {
    $http.get('http://localhost:8080/products').
        then(function(response) {
            $scope.products = response.data;
			var log = [];
			var sum = 0;
			angular.forEach(response.data, function(value, key) {
			  
			  sum = sum + Number(value.count);
			  $scope.summ = sum;			  
			}, log);
			
        });
});

