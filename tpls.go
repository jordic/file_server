package main

const templateList = `
<!DOCTYPE html>
<html lang="en" ng-app="fMgr">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    
    <link rel="stylesheet" href="//cdn.jsdelivr.net/codemirror/4.3.0/codemirror.css">
    <link rel="stylesheet" href="//cdn.jsdelivr.net/codemirror/4.3.0/theme/monokai.css">

    <link href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css" 
        rel="stylesheet" />
    <link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/select2/3.5.0/select2.min.css" >
    <!--
    <script src="//cdn.jsdelivr.net/g/codemirror"></script>
    -->
    <script src="//cdn.jsdelivr.net/codemirror/4.3.0/codemirror.js"></script>
    <script src="//cdn.jsdelivr.net/codemirror/4.3.0/addon/selection/active-line.js"></script>
    <script src="//cdn.jsdelivr.net/codemirror/4.3.0/keymap/sublime.js"></script>
    <script src="//cdn.jsdelivr.net/codemirror/4.3.0/addon/display/rulers.js"></script>
    <script src="//cdn.jsdelivr.net/codemirror/4.3.0/mode/css/css.js"></script>
    <script src="//cdn.jsdelivr.net/codemirror/4.3.0/mode/javascript/javascript.js"></script>
    <script src="//cdn.jsdelivr.net/codemirror/4.3.0/mode/markdown/markdown.js"></script>

    <script src="//code.jquery.com/jquery-1.11.0.min.js"> </script>
    <script src="//cdnjs.cloudflare.com/ajax/libs/select2/3.5.0/select2.min.js"> </script>
    <script src="//netdna.bootstrapcdn.com/bootstrap/3.1.1/js/bootstrap.min.js"> </script>
    <script src="//ajax.googleapis.com/ajax/libs/angularjs/1.2.19/angular.min.js"></script>
    <script src="//ajax.googleapis.com/ajax/libs/angularjs/1.2.19/angular-sanitize.js"></script>
    <script src="//cdn.jsdelivr.net/angular.bootstrap/0.11.0/ui-bootstrap-tpls.js"> </script>
    <script src="//cdn.jsdelivr.net/angular.bootstrap/0.11.0/ui-bootstrap-tpls.min.js"> </script>
    
    <style>
    body { font-family:Arial; font-size:12px }
    .list .glyphicon { margin-right:5px; color:grey; font-size:12px }
    .list .glyphicon-folder-open { color:#2a8dc4; }
    #title { color:#2a8dc4; font-size:16px; }

    #file_list li { padding-bottom:8px; margin-bottom:8px; list-style: none; border-bottom: 1px solid #eaeaea; }
    
    /* table sort */
    th.tablesort-sortable {-webkit-user-select: none; -khtml-user-select: none; -moz-user-select: none; -o-user-select: none; user-select: none; cursor: pointer; }
    table .tablesort-sortable:after{content:""; float:right; margin-top:7px; visibility:hidden; border-left:4px solid transparent; border-right:4px solid transparent; border-top:none; border-bottom:4px solid #000; }
    table .tablesort-desc:after{border-top:4px solid #000; border-bottom:none; } 
    table .tablesort-asc,table .tablesort-desc{ background-color:rgba(141, 192, 219, 0.25); }
    table .tablesort-sortable:hover:after, table .tablesort-asc:after, table .tablesort-desc:after {visibility:visible; }
    .showIfLast {display: none; }
    /* Only show it if it is also the last row of the table. */
    .showIfLast:last-child {display: table-row; }
    .showIfLast td {text-align: center; }
    .showIfLast td:after {content: "No data"; }

    thead td { font-weight:bold; }
    .dir, .delete { cursor:pointer; }
    .btn:focus {outline: none; }
    #fmessage { position: absolute; z-index: 50; top: 10px; display: block; right: 10px; box-shadow: 0 2px 10px 0 rgba(0, 0, 0, 0.26);
            background: #134166; min-height: 38px; min-width: 288px; padding: 8px 24px 0px; color: #666; border-radius: 2px;
        opacity: 1; font-size: 14px; }

    #fmessage.bg-danger { background-color:#f2dede; margin:0px; }
    #fmessage.bg-success { background-color:#dff0d8; margin:0px; }


    #editor { min-height:400px; height:80%; }
    .mtop { margin-top:15px; }

    .CodeMirror { font-family: 'Menlo', 'Iconsolata', 'Monaco', monospace; 
        font-size:12px;  height: 600px; line-height:16px; 
        border:1px solid #eaeaea;  }

    .rename { font-size:11px; color:green; display:block; 
        float:right; margin-right:10px;  }
    .tdrename { cursor:pointer; position:relative; }
    .tdrename:hover {  }
    .renameinput { position:absolute;z-index:5; background-color:#fff; top:3px; left:-3px;  }
    .renameinput input:focus { outline: none }
    .renameinput input { border:1px solid #eaeaea;
        width:80%; height:20px; background-color:#eaeaea; margin-top:5px; }
    .renameinput button { position:absolute; left:81%; top:4px; }
    .renameinput button.btn-default { position:absolute; left:88%; top:4px; }

    .list { margin-top:15px; }
    .table tr>td.actions { width:40px; }

    #finder button, #finder input  { background-color:#DFF0F5; 
        border-color:#DFF0F5; color:#000; }
    #finder .pt { position:absolute; top:35%; right:15px; }
    #finder input, #finder button, #finder li { font-size:12px; color:#047FC4; }
    .ui-select-highlight { color:#005E4B; }

    @media(max-width:767px){ 
        td.actions { width:30px; }
        thead td { font-size:11px;}
        #filter { display:none; }
        .breadcrumb { margin-bottom:5px; background-color:#fff;  }
        ol { background-color:#fff; }
        .form-control { font-size:12px }
        #finder { width:100%; margin-bottom:5px; }
        #finder .btn { font-size:12px;   }
        #list .dropdown-menu { left:-145px; }
        #list .dropdown-menu li { font-size:12px; }
        .renameinput { width:100%; }
        .renameinput input { width: 55%; }
        .renameinput button { left:58% }
        .renameinput button.btn-default { left:72%; }
    }


 /*!
 * ui-select
 * http://github.com/angular-ui/ui-select
 * Version: 0.3.1 - 2014-07-12T16:26:10.171Z
 * License: MIT
 */.ui-select-highlight{font-weight:700}.ui-select-offscreen{clip:rect(0 0 0 0)!important;width:1px!important;height:1px!important;border:0!important;margin:0!important;padding:0!important;overflow:hidden!important;position:absolute!important;outline:0!important;left:0!important;top:0!important}.ng-dirty.ng-invalid>a.select2-choice{border-color:#D44950}.selectize-input.selectize-focus{border-color:#007FBB!important}.selectize-control>.selectize-dropdown,.selectize-control>.selectize-input>input{width:100%}.ng-dirty.ng-invalid>div.selectize-input{border-color:#D44950}.btn-default-focus{color:#333;background-color:#EBEBEB;border-color:#ADADAD;text-decoration:none;outline:-webkit-focus-ring-color auto 5px;outline-offset:-2px;box-shadow:inset 0 1px 1px rgba(0,0,0,.075),0 0 8px rgba(102,175,233,.6)}.input-group>.ui-select-bootstrap.dropdown{position:static}.input-group>.ui-select-bootstrap>input.ui-select-search.form-control{border-radius:4px 0 0 4px}.ui-select-bootstrap>.ui-select-match{text-align:left}.ui-select-bootstrap>.ui-select-match>.caret{position:absolute;top:45%;right:15px}.ui-select-bootstrap>.ui-select-choices{width:100%;height:auto;max-height:200px;overflow-x:hidden}.ui-select-bootstrap .ui-select-choices-row>a{display:block;padding:3px 20px;clear:both;font-weight:400;line-height:1.42857143;color:#333;white-space:nowrap}.ui-select-bootstrap .ui-select-choices-row>a:focus,.ui-select-bootstrap .ui-select-choices-row>a:hover{text-decoration:none;color:#262626;background-color:#f5f5f5}.ui-select-bootstrap .ui-select-choices-row.active>a{color:#fff;text-decoration:none;outline:0;background-color:#5bc0de}.ui-select-match.ng-hide-add,.ui-select-search.ng-hide-add{display:none!important}.ui-select-bootstrap.ng-dirty.ng-invalid>button.btn.ui-select-match{border-color:#D44950}

    </style>
<script>

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


    $scope.Path = "[% .Path %]"
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




</script>
    <title>FileManager</title>
  </head>
  <body ng-controller="ListCtr">
  

  <div class="container"  style="position:relative">
    <div class="row hidden-xs">
        <div class="col-md-6">
        <h3 id="title">FileManager</h3>
        </div>
    </div>
    
    <div  ng-show="flash.message" id="fmessage" class="{{ flash.type }}" >
    <p>{{ flash.message }}</p>
  </div>

  </div>
<!-- controller -->
<div class="container" id="nav">


<div class="row" id="header">

<div class="col-md-8" id="breadcrumb">
    <ol class="breadcrumb">
        <li><a href="/"><span class="glyphicon glyphicon-home"> </span></a></li>
        <li ng-repeat="item in Rutas"><a href="{{ item.url }}">{{ item.name }}</a></li> 
    </ol>
</div>

<div class="col-md-4" ng-controller="FinderCtrl" id="finder">
    <form role="form">
        <ui-select type="text" ui-select2="s2opts" 
            width="100%" ng-model="item.selected"
            >
        <ui-select-match placeholder="Directory quick finder">/{{$select.selected.path}}</ui-select-match>
        <ui-select-choices repeat="item in dirs"
             refresh="getData($select.search)"
             refresh-delay="0">
             <span class="glyphicon glyphicon-folder-close dir"></span> &nbsp; /<span ng-bind-html="item.path | highlight: $select.search"> </span>
        </ui-select-choices>
      <!--<div ng-bind-html="$select.search"></div>-->
    </ui-select-choices>
    </ui-select>
    </form>
</div>

</div>

</div>

<div class="container" id="list">

<div class="row" ng-show="view=='main'">
<form id="upload_files" style="display:none" enctype="multipart/form-data">
<input type="file" id="file_upload" name="files" multiple style="display:none" >
</form>
 <form role="form">
    <div class="col-md-6 col-xs-6">
        
        <inline-modal action="createFolder" btntext="Add Folder" 
            icon="glyphicon-folder-open" title="Folder Name:" handler="get_data()" 
            message="Folder created Created!" path="Path"></inline-modal>
        
        <button type="button" class="btn btn-info btn-sm" type="file" ng-click="AddFiles()"
        tooltip-placement="top" tooltip="Upload Multiple Files"><span class="glyphicon glyphicon-plus"> </span> <span class="hidden-xs">&nbsp; Upload</span></button>

        
        <inline-modal action="new_file" btntext="Add File" 
            icon="glyphicon-file" title="Filename:" handler="get_data()" 
            message="File Created!" path="Path"></inline-modal>

        &nbsp;&nbsp;&nbsp;
        
        <button class="btn btn-danger btn-sm" ng-show="selected>0" ng-click="DeleteSelected()"
            tooltip-placement="top" tooltip="Delete selected"><span class="glyphicon glyphicon-trash"> </span></button>

        

    </div>
    <div class="col-md-6 pull-right text-right  col-xs-6">
    
            
            <div class="btn-group-sm" style="float:right">
                <button type="button" class="btn ng-class: check_tipo('all');"  
                    tooltip-placement="top" tooltip="Show all"
                    ng-click="filter('all')">All</button>
                <button type="button" 
                    tooltip-placement="top" tooltip="Show only folders"
                    class="btn ng-class: check_tipo('folder');" ng-click="filter('folder')"><span class="glyphicon glyphicon-folder-open"> </span></button>
                <button type="button" 
                    tooltip-placement="top" tooltip="Show only files"
                    class="btn ng-class: check_tipo('files');" ng-click="filter('files')"><span class="glyphicon glyphicon-file"> </span></button>
                <button type="button" class="btn ng-class: check_tipo('hidden');"  ng-click="filter('hidden')" 
                    tooltip-placement="top" tooltip="Show Hidden"><span class="glyphicon glyphicon-eye-open"> </span></button>
            </div>
            
            <input ng-model="query" class="form-control input-sm" placeholder="Filter" style="width:45%; float:right; margin-right:15px" id="filter">

    </div>
    </form>
</div>


<div class="row list" ng-show="view=='main'" id="list">
<div class="col-md-12">
    <table class="table" ts-wrapper>
        <thead>
            <tr>
                <td  class="hidden-xs">&nbsp</td>
                <td ts-criteria="IsDir" class="f">+</td>
                <td ts-criteria="Name" ts-default="Ascending">Filename</td>
                <td ts-criteria="Size|parseFloat" class="hidden-xs">Size</td>
                <td ts-criteria="ModTime"  class="hidden-xs">Last Modified</td>
                <td><span class="hidden-xs">Actions</span></td>
            </tr>
        </thead>
        <tbody>
            <tr ng-repeat="item in Files|filter:query|filter:{'IsDir':ff}|filter:{'IsHidden':hidden}" ts-repeat>
                <td width="20" class="hidden-xs"><input type="checkbox" name="checkboxs[]" value="{{ item.Name }}" 
                    ng-click="CheckboxToggle(this, $event)" /></td>
                <td width="20"><span class="glyphicon glyphicon-folder-open" ng-show="item.IsDir"></span>
                        <span ng-hide="item.IsDir" class="glyphicon glyphicon-file"></span></td>
                
                <td inline-edit item="item" path="Path"> </td>
                
                <td width="100" class="hidden-xs">{{ item.Size|bytes }}</td>
                <td width="140" class="hidden-xs">{{ item.ModTime|date:'dd/MM/yyyy HH:mm:ss' }}</td>
                <td class="actions">
                
                    
                    

                    <!--<a href="{{ item.Name }}/?format=zip" target="_self" ng-if="item.IsDir" 
                        class="glyphicon glyphicon-download-alt delete" tooltip-placement="top" tooltip="Donwload as Zip"> </a>-->
                    
                    <div class="btn-group" dropdown>
                        <a href class="dropdown-toggle">
                            <span class="glyphicon glyphicon-align-justify"> </span>
                        </a>
                        <ul class="dropdown-menu" role="menu">
                            <li><a ng-click="CopyFile(item.Name)" href="#">Copy</a></li>
                            <li><a ng-click="DeleteFile(item.Name)" href="#">Delete</a></li>
                            <li ng-if="!item.IsDir"><a ng-click="EditFile(item.Name)" href="#">Edit</a></li>
                            <li><a ng-click="RenameFile(item.Name)" href="#">Rename</a></li>
                            <li><a ng-click="Compress(item.Name)" href="#">Compress</a></li>
                            <li ng-if="item.IsDir"><a href="{{ item.Name }}/?format=zip">Download as Zip</a></li>
                            
                        </ul>

                    </div>

                    <span ng-if="item.IsText" ng-click="EditFile(item.Name)" class="glyphicon glyphicon-pencil delete hidden-xs" tooltip-placement="top" tooltip="Edit" > </span>

                </td>
            </tr>
        </tbody>

    </table>
</div>
</div>

</div>


<div ng-show="view=='edit'" class="container">

<div class="row">
    <form role="form">
    <div class="col-md-6 col-xs-6">
        <p>{{ currentEditedFile }}</p>
    </div>
    <div class="col-md-6 pull-right text-right col-xs-6">
        <button type="button" class="btn btn-info btn-sm" ng-click="ToView('main')"><span class="glyphicon glyphicon-arrow-left"> </span><sapn class="hidden-xs"> &nbsp; Back</sapn></button>
        <button type="button" class="btn btn-info btn-sm" ng-click="SaveFile()"><span class="glyphicon glyphicon-floppy-disk"> </span><span class="hidden-xs"> &nbsp; Save</span></button>
        
    </div>
    </form>
</div>

<div class="row" class="mtop" style="margin-top:20px">
    <div class="col-md-12">
        <textarea id="editor" ui-codemirror="{onLoad:codemirrorLoaded}"
            ui-codemirror-opts="editorOptions"
            ng-model="EditorCurrentContent"
            ui-refresh='EditorRefresh'>
        </textarea>
    </div>
</div>

</div>


<div class="container">
    <div class="row" style="margin-top:50px; border-top:1px solid #eaeaea; padding-top:20px; font-size:10px">
        <div class="col-md-12">
        <p><a href="http://github.com/jordic/file_server">http://github.com/jordic/file_server</a> -- v.[% .version %]
        </p></div>
    </div>
</div>
    

  </body>

<script type="text/ng-template" id="/inline-modal.html">
    <button class="btn btn-info btn-sm" ng-click="ClickButton()">
        <span class="glyphicon {{ icon }}"> </span><span class="hidden-xs">&nbsp; {{ btntext }}</span></button>

    <div ng-show="folder_popover==true" class="popover bottom am-flip-x" 
        style="display: block; width:200px">
                <div class="arrow"></div>
                <h3 class="popover-title">{{ title }}</h3>
                <div class="popover-content">
                    <input type="text" class="form-control input-sm" 
                        ng-model="filename" id="folder_field" 
                        tfocus="folder_popover">
                    <button type="button" ng-click="Process()" 
                        class="btn btn-info btn-sm pull-right" 
                            style="margin-top:5px; margin-bottom:5px">Add</button>
                </div>
        </div>


</script>

<script type="text/ng-template" id="/inline-edit.html">
    <td ng-mouseenter="roll=true" ng-mouseleave="roll=false" class="tdrename"
        ng-click="Show($event)">    
        <a href="{{ Path }}{{ item.Name }}" target="_self" 
            ng-if="!item.IsDir">{{ item.Name }}</a>
        <a  ng-click="Go(item.Name)" ng-if="item.IsDir" class="dir">{{ item.Name }}</a><span class="visible-xs small">{{ item.ModTime|date:'dd/MM/yyyy HH:mm:ss' }} | {{ item.Size|bytes }}</span>
    

    <div class="col-md-12 renameinput inline-form" ng-show="showrename==true">
        <input type="text" class="" name="value" 
            ng-model="item.Name" tfocus="showrename" />
        <button class="btn btn-info btn-xs" ng-click="SaveItem()">SAVE</button>
        <button class="btn btn-default btn-xs" ng-click="showrename=false">CANCEL</button>
    </div>
    </td>
</script>

<script>


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
        res.display = 'block'
        res.width = '200px'
        res.left -= 50
        res.top += 40
        
        //$log.log(res)
        content.css({
            'top': 35, 
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



/*
 angular-tablesort v1.0.4
 (c) 2013 Mattias Holmlund, http://mattiash.github.io/angular-tablesort
 License: MIT
*/

var tableSortModule = angular.module( 'tableSort', [] );

tableSortModule.directive('tsWrapper', function( $log, $parse ) {
    'use strict';
    return {
        scope: true,
        controller: function($scope) {
            $scope.sortExpression = [];
            $scope.headings = [];

            var parse_sortexpr = function( expr ) {
                return [$parse( expr ), null, false];
            };

            this.setSortField = function( sortexpr, element ) {
                var i;
                var expr = parse_sortexpr( sortexpr );
                if( $scope.sortExpression.length === 1
                    && $scope.sortExpression[0][0] === expr[0] ) {
                    if( $scope.sortExpression[0][2] ) {
                        element.removeClass( "tablesort-desc" );
                        element.addClass( "tablesort-asc" );
                        $scope.sortExpression[0][2] = false;
                    }
                    else {
                        element.removeClass( "tablesort-asc" );
                        element.addClass( "tablesort-desc" );
                        $scope.sortExpression[0][2] = true;
                    }
                }
                else {
                    for( i=0; i<$scope.headings.length; i=i+1 ) {
                        $scope.headings[i]
                            .removeClass( "tablesort-desc" )
                            .removeClass( "tablesort-asc" );
                    }
                    element.addClass( "tablesort-asc" );
                    $scope.sortExpression = [expr];
                }
            };

            this.addSortField = function( sortexpr, element ) {
                var i;
                var toggle_order = false;
                var expr = parse_sortexpr( sortexpr );
                for( i=0; i<$scope.sortExpression.length; i=i+1 ) {
                    if( $scope.sortExpression[i][0] === expr[0] ) {
                        if( $scope.sortExpression[i][2] ) {
                            element.removeClass( "tablesort-desc" );
                            element.addClass( "tablesort-asc" );
                            $scope.sortExpression[i][2] = false;
                        }
                        else {
                            element.removeClass( "tablesort-asc" );
                            element.addClass( "tablesort-desc" );
                            $scope.sortExpression[i][2] = true;
                        }
                        toggle_order = true;
                    }
                }
                if( !toggle_order ) {
                    element.addClass( "tablesort-asc" );
                    $scope.sortExpression.push( expr );
                }
            };

            this.registerHeading = function( headingelement ) {
                $scope.headings.push( headingelement );
            };

            $scope.sortFun = function( a, b ) {
                var i, aval, bval, descending, filterFun;
                for( i=0; i<$scope.sortExpression.length; i=i+1 ){
                    aval = $scope.sortExpression[i][0](a);
                    bval = $scope.sortExpression[i][0](b);
                    filterFun = b[$scope.sortExpression[i][1]];
                    if( filterFun ) {
                        aval = filterFun( aval );
                        bval = filterFun( bval );
                    }
                    if( aval === undefined ) {
                        aval = "";
                    }
                    if( bval === undefined ) {
                       bval = "";
                    }
                    descending = $scope.sortExpression[i][2];
                    if( aval > bval ) {
                        return descending ? -1 : 1;
                    }
                    else if( aval < bval ) {
                        return descending ? 1 : -1;
                    }
                }
                return 0;
            };
        }
    };
} );

tableSortModule.directive('tsCriteria', function() {
    return {
        require: "^tsWrapper",
        link: function(scope, element, attrs, tsWrapperCtrl) {
            var clickingCallback = function(event) {
                scope.$apply( function() {
                    if( event.shiftKey ) {
                        tsWrapperCtrl.addSortField(attrs.tsCriteria, element);
                    }
                    else {
                        tsWrapperCtrl.setSortField(attrs.tsCriteria, element);
                    }
                } );
            };
            element.bind('click', clickingCallback);
            element.addClass('tablesort-sortable');
            if( "tsDefault" in attrs && attrs.tsDefault !== "0" ) {
                tsWrapperCtrl.addSortField( attrs.tsCriteria, element );
                if( attrs.tsDefault == "descending" ) {
                    tsWrapperCtrl.addSortField( attrs.tsCriteria, element );
                }
            }
            tsWrapperCtrl.registerHeading( element );
        }
    };
});

tableSortModule.directive("tsRepeat", function($compile) {
    return {
        terminal: true,
        require: "^tsWrapper",
        priority: 1000000,
        link: function(scope, element) {
            var clone = element.clone();
            var tdcount = element[0].childElementCount;
            var repeatExpr = clone.attr("ng-repeat");
            repeatExpr = repeatExpr.replace(/^\s*([\s\S]+?)\s+in\s+([\s\S]+?)(\s+track\s+by\s+[\s\S]+?)?\s*$/,
                "$1 in $2 | tablesortOrderBy:sortFun$3");

            element.html("<td colspan='"+tdcount+"'></td>");
            element[0].className += " showIfLast";
            clone.removeAttr("ts-repeat");

            clone.attr("ng-repeat", repeatExpr);
            var clonedElement = $compile(clone)(scope);
            element.after(clonedElement);
        }
    };
} );

tableSortModule.filter( 'tablesortOrderBy', function(){
    return function(array, sortfun ) {
        if(!array) return;
        var arrayCopy = [];
        for ( var i = 0; i < array.length; i++) { arrayCopy.push(array[i]); }
        return arrayCopy.sort( sortfun );
    };
} );

tableSortModule.filter( 'parseInt', function(){
    return function(input) {
        return parseInt( input );
    };
} );

tableSortModule.filter( 'parseFloat', function(){
    return function(input) {
        return parseFloat( input );
    };
} );


'use strict';

/**
 * Binds a CodeMirror widget to a <textarea> element.
 */
angular.module('ui.codemirror', [])
  .constant('uiCodemirrorConfig', {})
  .directive('uiCodemirror', ['uiCodemirrorConfig', function (uiCodemirrorConfig) {

    return {
      restrict: 'EA',
      require: '?ngModel',
      priority: 1,
      compile: function compile() {

        // Require CodeMirror
        if (angular.isUndefined(window.CodeMirror)) {
          throw new Error('ui-codemirror need CodeMirror to work... (o rly?)');
        }

        return function postLink(scope, iElement, iAttrs, ngModel) {


          var options, opts, codeMirror, value;

          value = iElement.text();

          if (iElement[0].tagName === 'TEXTAREA') {
            // Might bug but still ...
            codeMirror = window.CodeMirror.fromTextArea(iElement[0], {
              value: value
            });
          } else {
            iElement.html('');
            codeMirror = new window.CodeMirror(function(cm_el) {
              iElement.append(cm_el);
            }, {
              value: value
            });
          }

          options = uiCodemirrorConfig.codemirror || {};
          opts = angular.extend({}, options, scope.$eval(iAttrs.uiCodemirror), scope.$eval(iAttrs.uiCodemirrorOpts));

          function updateOptions(newValues) {
            for (var key in newValues) {
              if (newValues.hasOwnProperty(key)) {
                codeMirror.setOption(key, newValues[key]);
              }
            }
          }

          updateOptions(opts);

          if (iAttrs.uiCodemirror) {
            scope.$watch(iAttrs.uiCodemirror, updateOptions, true);
          }


          if (ngModel) {
            // CodeMirror expects a string, so make sure it gets one.
            // This does not change the model.
            ngModel.$formatters.push(function (value) {
              if (angular.isUndefined(value) || value === null) {
                return '';
              } else if (angular.isObject(value) || angular.isArray(value)) {
                throw new Error('ui-codemirror cannot use an object or an array as a model');
              }
              return value;
            });


            // Override the ngModelController $render method, which is what gets called when the model is updated.
            // This takes care of the synchronizing the codeMirror element with the underlying model, in the case that it is changed by something else.
            ngModel.$render = function () {
              //Code mirror expects a string so make sure it gets one
              //Although the formatter have already done this, it can be possible that another formatter returns undefined (for example the required directive)
              var safeViewValue = ngModel.$viewValue || '';
              codeMirror.setValue(safeViewValue);
            };


            // Keep the ngModel in sync with changes from CodeMirror
            codeMirror.on('change', function (instance) {
              var newValue = instance.getValue();
              if (newValue !== ngModel.$viewValue) {
                // Changes to the model from a callback need to be wrapped in $apply or angular will not notice them
                scope.$apply(function() {
                  ngModel.$setViewValue(newValue);
                });
              }
            });
          }


          // Watch ui-refresh and refresh the directive
          if (iAttrs.uiRefresh) {
            scope.$watch(iAttrs.uiRefresh, function (newVal, oldVal) {
              // Skip the initial watch firing
              if (newVal !== oldVal) {
                codeMirror.refresh();
              }
            });
          }


          // Allow access to the CodeMirror instance through a broadcasted event
          // eg: $broadcast('CodeMirror', function(cm){...});
          scope.$on('CodeMirror', function(event, callback) {
            if (angular.isFunction(callback)) {
              callback(codeMirror);
            } else {
              throw new Error('the CodeMirror event requires a callback function');
            }
          });


          // onLoad callback
          if (angular.isFunction(opts.onLoad)) {
            opts.onLoad(codeMirror);
          }

        };
      }
    };
  }]);


/*!
 * ui-select
 * http://github.com/angular-ui/ui-select
 * Version: 0.3.1 - 2014-07-12T16:26:10.166Z
 * License: MIT
 */
!function(){"use strict";void 0===angular.element.prototype.querySelectorAll&&(angular.element.prototype.querySelectorAll=function(e){return angular.element(this[0].querySelectorAll(e))}),angular.module("ui.select",[]).constant("uiSelectConfig",{theme:"bootstrap",placeholder:"",refreshDelay:1e3}).service("uiSelectMinErr",function(){var e=angular.$$minErr("ui.select");return function(){var t=e.apply(this,arguments),c=t.message.replace(new RegExp("\nhttp://errors.angularjs.org/.*"),"");return new Error(c)}}).service("RepeatParser",["uiSelectMinErr",function(e){var t=this;t.parse=function(t){if(!t)throw e("repeat","Expected 'repeat' expression.");var c=t.match(/^\s*([\s\S]+?)\s+in\s+([\s\S]+?)(?:\s+track\s+by\s+([\s\S]+?))?\s*$/);if(!c)throw e("iexp","Expected expression in form of '_item_ in _collection_[ track by _id_]' but got '{0}'.",t);var s=c[1],l=c[2],i=c[3];if(c=s.match(/^(?:([\$\w]+)|\(([\$\w]+)\s*,\s*([\$\w]+)\))$/),!c)throw e("iidexp","'_item_' in '_item_ in _collection_' should be an identifier or '(_key_, _value_)' expression, but got '{0}'.",s);return{lhs:s,rhs:l,trackByExp:i}},t.getGroupNgRepeatExpression=function(){return"($group, $items) in $select.groups"},t.getNgRepeatExpression=function(e,t,c,s){var l=e+" in "+(s?"$items":t);return c&&(l+=" track by "+c),l}}]).controller("uiSelectCtrl",["$scope","$element","$timeout","RepeatParser","uiSelectMinErr",function(e,t,c,s,l){function i(){o.resetSearchInput&&(o.search=a,o.selected&&o.items.length&&(o.activeIndex=o.items.indexOf(o.selected)))}function n(e){var t=!0;switch(e){case h.Down:o.activeIndex<o.items.length-1&&o.activeIndex++;break;case h.Up:o.activeIndex>0&&o.activeIndex--;break;case h.Tab:case h.Enter:o.select(o.items[o.activeIndex]);break;case h.Escape:o.close();break;default:t=!1}return t}function r(){var e=t.querySelectorAll(".ui-select-choices-content"),c=e.querySelectorAll(".ui-select-choices-row");if(c.length<1)throw l("choices","Expected multiple .ui-select-choices-row but got '{0}'.",c.length);var s=c[o.activeIndex],i=s.offsetTop+s.clientHeight-e[0].scrollTop,n=e[0].offsetHeight;i>n?e[0].scrollTop+=i-n:i<s.clientHeight&&(o.isGrouped&&0===o.activeIndex?e[0].scrollTop=0:e[0].scrollTop-=s.clientHeight-i)}var o=this,a="";o.placeholder=void 0,o.search=a,o.activeIndex=0,o.items=[],o.selected=void 0,o.open=!1,o.focus=!1,o.focusser=void 0,o.disabled=void 0,o.resetSearchInput=void 0,o.refreshDelay=void 0;var u=t.querySelectorAll("input.ui-select-search");if(1!==u.length)throw l("searchInput","Expected 1 input.ui-select-search but got '{0}'.",u.length);o.activate=function(e){o.disabled||(i(),o.open=!0,c(function(){o.search=e||o.search,u[0].focus()}))},o.parseRepeatAttr=function(t,c){function i(t){o.groups={},angular.forEach(t,function(t){var s=e.$eval(c),l=angular.isFunction(s)?s(t):t[s];o.groups[l]?o.groups[l].push(t):o.groups[l]=[t]}),o.items=[],angular.forEach(Object.keys(o.groups).sort(),function(e){o.items=o.items.concat(o.groups[e])})}function n(e){o.items=e}var r=s.parse(t),a=c?i:n;o.isGrouped=!!c,o.itemProperty=r.lhs,e.$watchCollection(r.rhs,function(e){if(void 0===e||null===e)o.items=[];else{if(!angular.isArray(e))throw l("items","Expected an array but got '{0}'.",e);a(e)}})};var d;o.refresh=function(t){void 0!==t&&(d&&c.cancel(d),d=c(function(){e.$eval(t)},o.refreshDelay))},o.setActiveItem=function(e){o.activeIndex=o.items.indexOf(e)},o.isActive=function(e){return o.items.indexOf(e[o.itemProperty])===o.activeIndex},o.select=function(e){o.selected=e,o.close()},o.close=function(){o.open&&(i(),o.open=!1,o.focusser[0].focus())};var h={Enter:13,Tab:9,Up:38,Down:40,Escape:27};u.on("keydown",function(t){if(o.items&&o.items.length>=0){var c=t.which;switch(e.$apply(function(){var e=n(c);e&&(t.preventDefault(),t.stopPropagation())}),c){case h.Down:case h.Up:r()}}}),e.$on("$destroy",function(){u.off("keydown")})}]).directive("uiSelect",["$document","uiSelectConfig","uiSelectMinErr","$compile",function(e,t,c,s){return{restrict:"EA",templateUrl:function(e,c){var s=c.theme||t.theme;return s+"/select.tpl.html"},replace:!0,transclude:!0,require:["uiSelect","ngModel"],scope:!0,controller:"uiSelectCtrl",controllerAs:"$select",link:function(t,l,i,n,r){function o(e){var c=!1;c=window.jQuery?window.jQuery.contains(l[0],e.target):l[0].contains(e.target),c||(a.close(),t.$digest())}var a=n[0],u=n[1],d=angular.element("<input ng-disabled='$select.disabled' class='ui-select-focusser ui-select-offscreen' type='text' aria-haspopup='true' role='button' />");s(d)(t),a.focusser=d,l.append(d),d.bind("focus",function(){t.$evalAsync(function(){a.focus=!0})}),d.bind("blur",function(){t.$evalAsync(function(){a.focus=!1})}),d.bind("keydown",function(e){e.which===h.TAB||h.isControl(e)||h.isFunctionKey(e)||e.which===h.ESC||((e.which==h.DOWN||e.which==h.UP||e.which==h.ENTER||e.which==h.SPACE)&&(e.preventDefault(),e.stopPropagation(),a.activate()),t.$digest())}),d.bind("keyup input",function(e){e.which===h.TAB||h.isControl(e)||h.isFunctionKey(e)||e.which===h.ESC||e.which==h.ENTER||(a.activate(d.val()),d.val(""),t.$digest())});var h={TAB:9,ENTER:13,ESC:27,SPACE:32,LEFT:37,UP:38,RIGHT:39,DOWN:40,SHIFT:16,CTRL:17,ALT:18,PAGE_UP:33,PAGE_DOWN:34,HOME:36,END:35,BACKSPACE:8,DELETE:46,isArrow:function(e){switch(e=e.which?e.which:e){case h.LEFT:case h.RIGHT:case h.UP:case h.DOWN:return!0}return!1},isControl:function(e){var t=e.which;switch(t){case h.SHIFT:case h.CTRL:case h.ALT:return!0}return e.metaKey?!0:!1},isFunctionKey:function(e){return e=e.which?e.which:e,e>=112&&123>=e}};i.$observe("disabled",function(){a.disabled=void 0!==i.disabled?i.disabled:!1}),i.$observe("resetSearchInput",function(){var e=t.$eval(i.resetSearchInput);a.resetSearchInput=void 0!==e?e:!0}),t.$watch("$select.selected",function(e){u.$viewValue!==e&&u.$setViewValue(e)}),u.$render=function(){a.selected=u.$viewValue},e.on("click",o),t.$on("$destroy",function(){e.off("click",o)}),r(t,function(e){var t=angular.element("<div>").append(e),s=t.querySelectorAll(".ui-select-match");if(1!==s.length)throw c("transcluded","Expected 1 .ui-select-match but got '{0}'.",s.length);l.querySelectorAll(".ui-select-match").replaceWith(s);var i=t.querySelectorAll(".ui-select-choices");if(1!==i.length)throw c("transcluded","Expected 1 .ui-select-choices but got '{0}'.",i.length);l.querySelectorAll(".ui-select-choices").replaceWith(i)})}}}]).directive("uiSelectChoices",["uiSelectConfig","RepeatParser","uiSelectMinErr","$compile",function(e,t,c,s){return{restrict:"EA",require:"^uiSelect",replace:!0,transclude:!0,templateUrl:function(t){var c=t.parent().attr("theme")||e.theme;return c+"/choices.tpl.html"},compile:function(l,i){var n=t.parse(i.repeat),r=i.groupBy;return function(l,i,o,a,u){if(r){var d=i.querySelectorAll(".ui-select-choices-group");if(1!==d.length)throw c("rows","Expected 1 .ui-select-choices-group but got '{0}'.",d.length);d.attr("ng-repeat",t.getGroupNgRepeatExpression())}var h=i.querySelectorAll(".ui-select-choices-row");if(1!==h.length)throw c("rows","Expected 1 .ui-select-choices-row but got '{0}'.",h.length);h.attr("ng-repeat",t.getNgRepeatExpression(n.lhs,"$select.items",n.trackByExp,r)).attr("ng-mouseenter","$select.setActiveItem("+n.lhs+")").attr("ng-click","$select.select("+n.lhs+")"),u(function(e){var t=i.querySelectorAll(".ui-select-choices-row-inner");if(1!==t.length)throw c("rows","Expected 1 .ui-select-choices-row-inner but got '{0}'.",t.length);t.append(e),s(i)(l)}),a.parseRepeatAttr(o.repeat,r),l.$watch("$select.search",function(){a.activeIndex=0,a.refresh(o.refresh)}),o.$observe("refreshDelay",function(){var t=l.$eval(o.refreshDelay);a.refreshDelay=void 0!==t?t:e.refreshDelay})}}}}]).directive("uiSelectMatch",["uiSelectConfig",function(e){return{restrict:"EA",require:"^uiSelect",replace:!0,transclude:!0,templateUrl:function(t){var c=t.parent().attr("theme")||e.theme;return c+"/match.tpl.html"},link:function(t,c,s,l){s.$observe("placeholder",function(t){l.placeholder=void 0!==t?t:e.placeholder})}}}]).filter("highlight",function(){function e(e){return e.replace(/([.?*+^$[\]\\(){}|-])/g,"\\$1")}return function(t,c){return c&&t?t.replace(new RegExp(e(c),"gi"),'<span class="ui-select-highlight">$&</span>'):t}})}(),angular.module("ui.select").run(["$templateCache",function(e){e.put("bootstrap/choices.tpl.html",'<ul class="ui-select-choices ui-select-choices-content dropdown-menu" role="menu" aria-labelledby="dLabel" ng-show="$select.items.length > 0"><li class="ui-select-choices-group"><div class="divider" ng-show="$index > 0"></div><div ng-show="$select.isGrouped" class="ui-select-choices-group-label dropdown-header">{{$group}}</div><div class="ui-select-choices-row" ng-class="{active: $select.isActive(this)}"><a href="javascript:void(0)" class="ui-select-choices-row-inner"></a></div></li></ul>'),e.put("bootstrap/match.tpl.html",'<button type="button" class="btn btn-sm form-control ui-select-match" tabindex="-1" ng-hide="$select.open" ng-disabled="$select.disabled" ng-class="{\'btn-sm-focus\':$select.focus}" ;="" ng-click="$select.activate()"><span ng-hide="$select.selected !== undefined" class="text-muted">{{$select.placeholder}}</span> <span ng-show="$select.selected !== undefined" ng-transclude=""></span> <span class="pt glyphicon glyphicon-search"></span></button>'),e.put("bootstrap/select.tpl.html",'<div class="ui-select-bootstrap dropdown" ng-class="{open: $select.open}"><div class="ui-select-match"></div><input type="text" autocomplete="off" tabindex="-1" class="form-control ui-select-search" placeholder="{{$select.placeholder}}" ng-model="$select.search" ng-show="$select.open"><div class="ui-select-choices"></div></div>'),e.put("select2/choices.tpl.html",'<ul class="ui-select-choices ui-select-choices-content select2-results"><li class="ui-select-choices-group" ng-class="{\'select2-result-with-children\': $select.isGrouped}"><div ng-show="$select.isGrouped" class="ui-select-choices-group-label select2-result-label">{{$group}}</div><ul class="select2-result-sub"><li class="ui-select-choices-row" ng-class="{\'select2-highlighted\': $select.isActive(this)}"><div class="select2-result-label ui-select-choices-row-inner"></div></li></ul></li></ul>'),e.put("select2/match.tpl.html",'<a class="select2-choice ui-select-match" ng-class="{\'select2-default\': $select.selected === undefined}" ng-click="$select.activate()"><span ng-hide="$select.selected !== undefined" class="select2-chosen">{{$select.placeholder}}</span> <span ng-show="$select.selected !== undefined" class="select2-chosen" ng-transclude=""></span> <span class="select2-arrow"><b></b></span></a>'),e.put("select2/select.tpl.html",'<div class="select2 select2-container" ng-class="{\'select2-container-active select2-dropdown-open\': $select.open,\n                \'select2-container-disabled\': $select.disabled,\n                \'select2-container-active\': $select.focus }"><div class="ui-select-match"></div><div class="select2-drop select2-with-searchbox select2-drop-active" ng-class="{\'select2-display-none\': !$select.open}"><div class="select2-search"><input type="text" autocomplete="off" autocorrect="off" autocapitalize="off" spellcheck="false" class="ui-select-search select2-input" ng-model="$select.search"></div><div class="ui-select-choices"></div></div></div>'),e.put("selectize/choices.tpl.html",'<div ng-show="$select.open" class="ui-select-choices selectize-dropdown single"><div class="ui-select-choices-content selectize-dropdown-content"><div class="ui-select-choices-group optgroup"><div ng-show="$select.isGrouped" class="ui-select-choices-group-label optgroup-header">{{$group}}</div><div class="ui-select-choices-row" ng-class="{active: $select.isActive(this)}"><div class="option ui-select-choices-row-inner" data-selectable=""></div></div></div></div></div>'),e.put("selectize/match.tpl.html",'<div ng-hide="$select.open || $select.selected === undefined" class="ui-select-match" ng-transclude=""></div>'),e.put("selectize/select.tpl.html",'<div class="selectize-control single"><div class="selectize-input" ng-class="{\'focus\': $select.open, \'disabled\': $select.disabled, \'selectize-focus\' : $select.focus}" ng-click="$select.activate()"><div class="ui-select-match"></div><input type="text" autocomplete="off" tabindex="-1" class="ui-select-search" placeholder="{{$select.placeholder}}" ng-model="$select.search" ng-hide="$select.selected && !$select.open" ng-disabled="$select.disabled"></div><div class="ui-select-choices"></div></div>')}]);





</script>




</html>
`
