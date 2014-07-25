function is_mob() {
   if(window.innerWidth <= 800) {
     return true;
   } else {
     return false;
   }
}


var fMgr = angular.module('fMgr', ['tableSort', 'ui.bootstrap', 
        'ui.codemirror', 'tempoModule', 'ui.select', 'ngSanitize', 'ngDialog']);

fMgr.config(['$locationProvider', function($locationProvider){
    $locationProvider.html5Mode(true);
}]).config(['$httpProvider', function($httpProvider) {
    $httpProvider.defaults.headers.common["angular"] = "true";
}]);

fMgr.directive('tfocus', function($timeout){
   
    return {
        replace: true,
        restrict: 'A',
        link: function(scope, element, attr){
            scope.$watch(attr.tfocus, function(value){
               //element.css('display', value ? '' : 'none');
               if(value == true)
                    $timeout(function(){
                        element[0].focus()
                        element[0].select()
                    }, 0)
                    
               
            });
        }
    };
})

fMgr.filter('bytes', function() {
    return function(bytes, precision) {
        if (isNaN(parseFloat(bytes)) || !isFinite(bytes)) return '-';
        if (bytes==0) return '0 bytes';
        if (typeof precision === 'undefined') precision = 1;
        var units = ['bytes', 'kB', 'MB', 'GB', 'TB', 'PB'],
            number = Math.floor(Math.log(bytes) / Math.log(1024));
        return (bytes / Math.pow(1024, Math.floor(number))).toFixed(precision) +  ' ' + units[number];
    }
});

fMgr.service('Flash', function($timeout){

    var $scope;
    var duration = 3000

    function message(type, msg, dur) {

        if(dur) {
            d = dur
        } else {
            d = duration
        }

        $scope.flash = { type:type, message:msg }
        $timeout(function(){ 
            $scope.flash = undefined
         }, d)
    }

    this.scope = function(sc) {
        $scope = sc
        return this
    }

    this.duration = function(d) {
        duration = d
        return this
    }

    this.success = function(msg) {
        message("bg-success", msg)
    }

    this.error = function(msg) {
        message("bg-danger", msg)        
    }

}).service('Path', function(){
    this.path = ''
    this.Get = function() {
        return this.path
    }
    this.Set = function(p) {
        this.path = p
    }
})

fMgr.factory('ServerCommand', function($http, $q, Flash){

    var opts = {
        'new_file': function(p1) {
            return {
                action: 'save',
                params: {
                    file: p1,
                    content: " " 
                }
            }
        },
        createFolder: function(p1) {
          return {
                action: 'createFolder',
                params: {
                    source: p1,
                }
            }  
        }
    }


    var queryServer = function(params) {
        return $http.post("/", params)
    }

    var on_error = function(data) {
        //console.log(data)
        Flash.error("server error")
    }

    return {
        get_raw: function(params) {
            return queryServer(params)
        },
        get: function(params, ok, refresh) {
            queryServer(params).then(function(d){
                if(d.data.status == 0) {
                    Flash.success(ok)
                    if(refresh)
                        refresh()
                } else {
                    Flash.error(d.data.message, 5000)
                }
            }, on_error)
        },
        
        prepare: function(type, p1, p2, p3) {
            return opts[type](p1, p2, p3)
        }

    }
})

fMgr.factory('Downloader', function($q) {
    return {
        Async: function(file) {
            var deferred = $q.defer();
            var xhr = new XMLHttpRequest();
            xhr.open('GET', file, true);
            xhr.onreadystatechange = function(e) {
                if(this.readyState == 4 && this.status==200) {
                    deferred.resolve(this.responseText)
                }
            }
            xhr.send();
            return deferred.promise;
        },
        Stream: function(file) {

        }
    }
})


// Directives

var tempoModule = angular.module('tempoModule', ['ui.bootstrap']);

tempoModule.directive('inlineEdit', function($log, $location, $position, 
    ServerCommand, $timeout){

    return {
        templateUrl: "/inline-edit.html",
        link: link,
        scope: {
            item: "=item",
            Path: "=path"
        },
        replace: true,
        restrict: 'AEC',
    }

    
    function link($scope, element, attrs) {
        
        var old_path = $scope.Path + $scope.item.Name

        $scope.Show = function(e) {

            if( window.is_mob() ) 
                return

            if(e.target.nodeName != "TD") {
                e.stopPropagation()
                return
            }
            $scope.newName = $scope.item.Name
            $scope.showrename = true

            //var ele = angular.element('#file_' + $scope.item.Name )
            //ele.select()

        }

        $scope.CloseInline = function() {

            $scope.showrename = false
                //angular.element(window).off("click", closeInline)
            }
    
        $scope.SaveItem = function() {
            
            //var old_path = $scope.Path + old_name
            $timeout(function(){
                $scope.item.Name = $scope.newName    
            }, 0)
            
            var res = $scope.Path + $scope.item.Name

            if( old_path == res ) return;

            ServerCommand.get({
                    action: 'rename', 
                    params: {
                        source: old_path,
                        dest: res
                    }
                }, "File renamed", null)

            $scope.showrename = false;

        }

        $scope.Go = function(route) {
            $location.url( $scope.Path + route + "/" )
        }

    }

    

})

tempoModule.directive('ngEnter', function () {
    return function (scope, element, attrs) {
        element.bind("keydown keypress", function (event) {
            if(event.which === 13) {
                scope.$apply(function (){
                    scope.$eval(attrs.ngEnter);
                });

                event.preventDefault();
            }
        });
    };
});

tempoModule.directive('ngEsc', function () {
    return function (scope, element, attrs) {
        element.bind("keydown keypress", function (event) {
            //console.log(event.which)
            if(event.which === 27) {
                scope.$apply(function (){
                    scope.$eval(attrs.ngEsc);
                });

                event.preventDefault();
            }
        });
    };
});


tempoModule.directive('unbindable', function(){
    return {
        scope: true
    };
});

tempoModule.directive('inlineModal', function($log, $position, ServerCommand){

    return {
        templateUrl: "/inline-modal.html",
        link: link,
        restrict: 'AEC',
        scope: {
            action: '@action',
            btntext: '@',
            icon: '@',
            title: '@title',
            handler: '&',
            Path: '=path',
            message: '@'
        }
    }

    function link($scope, element, attrs) {

        var content = angular.element( '.popover', element )
        
        var res = $position.position(element)

        element.css('position', 'relative')
        //content.css('position', 'absolute')

        //console.log(res)

        res.display = 'block'
        res.width = '200px'
        

        if( is_mob() ) {
            res.left = 0
            res.top = 20
        } else {
            res.left = 0
            res.top = 25
        }


        
        
        //$log.log(res)
        content.css({
            'top': res.top, 
            'left': res.left,
            'width': '200px'
        });

        $scope.folder_popover = false

        $scope.ClickButton = function() {
            //alert('click')
            if( $scope.folder_popover == true ) {
                $scope.folder_popover = false
            } else {
                $scope.folder_popover = true
            }

            var closePopOver = function(e) {
                var child = angular.element("*", element)
                if ( angular.element(e.target).is(child) ) {
                    return false
                }
                angular.element(window).off("click", closePopOver)
                $scope.folder_popover = undefined
                $scope.$apply()
            }

            $scope.Close = function() {
                $scope.folder_popover = undefined
            }

            if($scope.folder_popover === true) {
                angular.element(window).on("click", closePopOver)
            } 
        }

        $scope.Process = function() {
            var obj = ServerCommand.prepare($scope.action, $scope.Path + $scope.filename )
            var msg = $scope.message
            //console.log(res)
            //$scope.handler()
            ServerCommand.get(obj, msg, $scope.handler)

            $scope.ClickButton()
            $scope.filename = ""
            //$scope.$apply()
        }
  
    }

    
})