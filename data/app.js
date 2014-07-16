



function is_mob() {
   if(window.innerWidth <= 800) {
     return true;
   } else {
     return false;
   }
}


var fMgr = angular.module('fMgr', ['tableSort', 'ui.bootstrap', 
        'ui.codemirror', 'tempoModule', 'ui.select', 'ngSanitize']);

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

    function message(type, msg) {
    
        $scope.flash = { type:type, message:msg }
        $timeout(function(){ 
            $scope.flash = undefined
         }, duration)
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





fMgr.controller("ListCtr", function($scope, $http, $location, 
        $document, $window, $timeout, ServerCommand, Flash){


    $scope.IsMobile = window.is_mob()

    // Config flash
    Flash.scope( $scope );


    $scope.Path = APP_PATH;
    $scope.view = 'main'
    
    var get_data = function() {
        $http.get($scope.Path + "?format=json")
            .then(function(res){
                //console.log(res)
                if(res.data=="null")
                    $scope.Files = []
                else
                    $scope.Files = res.data
              

                angular.forEach($scope.Files, function(item){
                    item.ModTime = new Date(item.ModTime)
                    //console.log( item.Name.indexOf(".") )
                    if(item.Name.indexOf(".") === 0) item.IsHidden = true;
                    else item.IsHidden = false
                })
                $scope.Rutas = GetRuta()
                $scope.selected = 0

            }, function(error){
                Flash.duration(10000).error( "Server Disconnected" )
            })
        $scope.view = 'main'
    }
    $scope.get_data = get_data
    get_data()

    //Flash_Message('bg-success', "Data Loaded")

    $scope.hidden = false

    $scope.filter = function(tipo) {
        if(tipo=="all") {  
            $scope.ff = undefined
        }
        if(tipo == "files") {
            $scope.ff = false
        }
        if(tipo == "folder") {
            $scope.ff = true   
        }

        if(tipo == "hidden") {
            if($scope.hidden == true )
                $scope.hidden = false
            else
                $scope.hidden = true   
        }
    }

    $scope.check_tipo = function(tipo) {
        var s = 0;
        if( tipo == "all" && $scope.ff == undefined) s = 1;
        if( tipo == "folder" && $scope.ff == true) s = 1;
        if( tipo == "files" && $scope.ff == false) s = 1;
        if( tipo == "hidden" && $scope.hidden == true) s = 1;
        res = ['btn-default', 'btn-info']
        return res[s]
    }

    $scope.load_dir = function(name) {
        //alert(name)
        var path = $scope.Path + name + "/"
        //get_data()
        $location.path( path ) 
        return false
    }

    GetRuta = function() {
        var parts = $scope.Path.split("/")
        result = []

        function get_route(index) {
            var p = "/"
            for(var i=1; i<=index; i++) {
                p += parts[i] + "/"
            }
            return p
        }


        for(var i=1; i<parts.length-1; i++) {
            var obj = { 
                name: parts[i],
                url: get_route(i)
            }


            result.push(obj)
        }
        return result
    }

    $scope.$on('$locationChangeSuccess', function(){
        //console.log('ara?')
        var newl = $location.path()
        $scope.OldPath = $scope.Path
         if( $scope.Path != newl ) {
            $scope.Path = newl
            get_data()
         }
         
    })

    
    $scope.RenameFile = function(f) {
        var old_path = $scope.Path +  f
        var res = prompt("Rename/Move File?", old_path)
        if(res ) {
            if(res == old_path) return
            ServerCommand.get({
                action: 'rename', 
                params: {
                    source: old_path,
                    dest: res
                }
            }, "File renamed", get_data)
        }   
    }

    $scope.CopyFile = function(f) {
        var old_path = $scope.Path +  f
        var res = prompt("Copy To:", old_path)
        if(res ) {
            if(res == old_path) return
            ServerCommand.get({
                action: 'copy', 
                params: {
                    source: old_path,
                    dest: res
                }
            }, "File copied", get_data)
        }   
    }


    

    $scope.Compress = function(item) {

        ServerCommand.get({
            action: 'compress',
            params: {
                source: $scope.Path + item
            }
        }, "File_compressed", get_data)

    }


    $scope.AddFiles = function() {

        var uploadFiles = function(evt){
            var files = evt.target.files;
            var formData = new FormData( document.getElementById('upload_files') );
           
            var xhr = new XMLHttpRequest();
            xhr.open('PUT', $scope.Path, true);
            
            xhr.onload = function(e) {
                //console.log(e)
                document.getElementById('file_upload').removeEventListener('change', uploadFiles);
                document.getElementById("file_upload").value = "";
                Flash.success("File uploaded")
                get_data()
             };

             xhr.onerror = function(e) {
                Flash.error("Error uploading file")
             }

            xhr.upload.onprogress = function(e) {
                //console.log(e)
                if (e.lengthComputable) {
                    //console.log( (e.loaded / e.total) * 100 )
                    $scope.flash.message = "Uploading... " + parseInt((e.loaded / e.total) * 100) + "%"
                    $scope.$apply()
                }
             }

             xhr.send(formData)
             $scope.flash = { type:"bg-success", message: "Uploading file... " }

        };


        document.getElementById('file_upload').addEventListener('change', uploadFiles, false)
        document.getElementById('file_upload').click();
        //console.log("file upload")
    }


    

    // Editor
    $scope.editorOptions = {
        lineWrapping : true,
        lineNumbers: true,
        keyMap: 'sublime',
        //theme: 'monokai'
    };

    $scope.EditFile = function(item) {
        $scope.currentEditedFile = $scope.Path + item
        var file = $scope.Path + item
        var noJsonTransform = function(data) { return data; };

        $http.get( file, {
            transformResponse: noJsonTransform}).then(function(d){
            //console.log(d)
            $scope.EditorCurrentContent = d.data
            $scope.EditorOldContent = $scope.EditorCurrentContent
            $scope.EditorRefresh = true
            $scope.view = 'edit'
            $scope.$apply()
            $scope.EditorInstance.refresh()

        })
        
    }

    $scope.codemirrorLoaded = function(_editor){
        // Editor part
        $scope.EditorInstance = _editor;
    };


    $scope.ToView = function(view) {

        if( $scope.EditorOldContent != $scope.EditorCurrentContent ) {
            var res = confirm("Loose changes?")
            if(!res) {
                return
            }
        }

        $scope.view = view
        $scope.EditorRefresh = false
    }

    $scope.SaveFile = function(exit) {
        
        var onSave = function(){
            $scope.EditorOldContent = $scope.EditorCurrentContent
               if(exit==true) 
                    $scope.ToView('main')
        }

        ServerCommand.get({
            action: 'save',
                params: {
                    file: $scope.currentEditedFile,
                    content: $scope.EditorCurrentContent
                }
            }, "File Saved", onSave)
    }


    // Multiselect
    $scope.selected = 0
    var lastChecked = null
    $scope.CheckboxToggle = function(item, evt) {
        
        var list = angular.element('input[type="checkbox"')        
        if(evt.shiftKey) {
            var start = list.index(evt.target)
            var end = list.index(lastChecked)
            list.slice( Math.min(start,end), Math.max(start,end)+ 1 ).attr('checked', lastChecked.checked)
        }

        lastChecked = evt.target

        $scope.selected = 0
        list.each(function(i, item){
            if( item.checked == true ) {
                $scope.selected++
            }
        })

    }

    $scope.DeleteFile = function(item) {
        var res = confirm("Are you sure?")
        if(res) {
            ServerCommand.get({
                action: 'delete',
                paramslist: [$scope.Path + item]
            }, "File deleted!", get_data)
        }
    }


    $scope.DeleteSelected = function() {
        var items = []
        angular.element('input[type="checkbox"]').each(function(i, item){
            if( item.checked == true ) {
                items.push( $scope.Path + item.value )
            }
        })

        ServerCommand.get({
            action: 'delete',
            paramslist: items
        }, "Files deleted!", get_data)
    }

    
})
    

fMgr.controller("FinderCtrl", function($scope, $http, $location){
    
    $scope.dirs = []
    $scope.item = {}
    $scope.getData = function(params) {
        //console.log(params)
        var q = params
        if(q.length < 3)
            return
        
        return $http.get("/-/api/dirs?q=" + q, {}).then(
            function(response) {
                $scope.dirs = response.data
            });
    }

    
    /*$scope.dselected = ''*/
    $scope.$watch('item.selected', function(nvalue, ovalue) {
        //console.log(old, nl)
        if(nvalue)
            $location.path("/" + nvalue.path + "/")
    })

})




// Directives

var tempoModule = angular.module('tempoModule', ['ui.bootstrap']);

tempoModule.directive('inlineEdit', function($log, $location, $position, ServerCommand){

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
            //console.log(e.target.nodeName)
            //console.log(element)

            if( window.is_mob() ) 
                return

            if(e.target.nodeName != "TD") {
                e.stopPropagation()
                return
            }

            //console.log("passa")

            if( $scope.showrename == true ) {
                $scope.showrename = false
            } else {
                $scope.showrename = true
            }
        }
    
        $scope.SaveItem = function() {
            
            //var old_path = $scope.Path + old_name
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