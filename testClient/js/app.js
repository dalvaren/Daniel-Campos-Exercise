// Ionic Starter App

// angular.module is a global place for creating, registering and retrieving Angular modules
// 'starter' is the name of this angular module example (also set in a <body> attribute in index.html)
// the 2nd parameter is an array of 'requires'
angular.module('application', [])
  .config(function($httpProvider) {
    $httpProvider.defaults.useXDomain = true;
  })
  .controller('AppCtrl', ['$scope', '$http',
    function($scope, $http) {
      var baseUrl = "http://localhost:8081"
      $scope.taskList = [];
      $scope.openModel = false;

      $http.defaults.headers.put = {
        'Access-Control-Allow-Origin': '*',
        'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS',
        'Access-Control-Allow-Headers': 'Content-Type, X-Requested-With',
      };
      $http.defaults.useXDomain = true;

      $scope.getTasks = function () {
        $http({
          url: baseUrl + '/task/',
          method: "GET",
          withCredentials: false,
          headers: {
            'Content-Type': 'application/json; charset=utf-8'
          }
        }).success(function(response){
          $scope.taskList = response.items;
        });
      };

      $scope.getTasks();

      $scope.eraseTaskFields = function(){
        $scope.ID = '';
        $scope.Title = '';
        $scope.Description = '';
        $scope.Priority = '';
      };

      $scope.editTask = function(taskId){
        $scope.eraseTaskFields();
        $scope.openModel = true;
        if (taskId != undefined) {
          $scope.taskList.forEach(function(task){
            if (task.ID == taskId) {
              $scope.ID = task.ID;
              $scope.Title = task.Title;
              $scope.Description = task.Description;
              $scope.Priority = task.Priority;
            }
          });
        }
      };

      $scope.closeTask = function() {
        $scope.openModel = false;
        $scope.eraseTaskFields();
      };

      $scope.deleteTask = function() {
        if($scope.ID != undefined && $scope.ID != '') {
          $http({
            url: baseUrl + '/task/' + $scope.ID,
            method: "DELETE",
            withCredentials: false,
            headers: {
              'Content-Type': 'application/json; charset=utf-8'
            }
          }).success(function(response){
            $scope.getTasks();
            $scope.closeTask();
          }).error(function(response){
            alert(response.message);
          });
        }
      };

      $scope.saveTask = function() {
        if($scope.ID == undefined || $scope.ID == '') {
          $http({
            url: baseUrl + '/task/',
            method: "POST",
            withCredentials: false,
            data    : {
              title: $scope.Title,
              description: $scope.Description,
              priority: $scope.Priority
            },
            headers: {
              'Content-Type': 'application/json; charset=utf-8'
            }
          }).success(function(response){
            $scope.getTasks();
            $scope.closeTask();
          }).error(function(response){
            alert(response.message);
          });
        } else {
          $http({
            url: baseUrl + '/task/' + $scope.ID,
            method: "PUT",
            withCredentials: false,
            data    : {
              title: $scope.Title,
              description: $scope.Description,
              priority: $scope.Priority
            },
            headers: {
              'Content-Type': 'application/json; charset=utf-8'
            }
          }).success(function(response){
            $scope.getTasks();
            $scope.closeTask();
          }).error(function(response){
            alert(response.message);
          });
        }
      };

      $scope.completeTask = function () {
        $http({
          url: baseUrl + '/task/' + $scope.ID,
          method: "PUT",
          withCredentials: false,
          data    : {
            title: $scope.Title,
            description: $scope.Description,
            priority: $scope.Priority.toString(),
            completed: true
          },
          headers: {
            'Content-Type': 'application/json; charset=utf-8'
          }
        }).success(function(response){
          $scope.getTasks();
          $scope.closeTask();
        }).error(function(response){
          alert(response.message);
        });
      };

  }]);
