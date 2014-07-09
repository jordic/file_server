package main

const templateList = `
<!DOCTYPE html>
<html lang="en" ng-app="fMgr">
  <head>
    <meta charset="utf-8">
    
    <link rel="stylesheet" href="//cdn.jsdelivr.net/codemirror/4.3.0/codemirror.css">
    <link rel="stylesheet" href="//cdn.jsdelivr.net/codemirror/4.3.0/theme/monokai.css">

    <link href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css" 
        rel="stylesheet" />
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
    <script src="//netdna.bootstrapcdn.com/bootstrap/3.1.1/js/bootstrap.min.js"> </script>
    <script src="//ajax.googleapis.com/ajax/libs/angularjs/1.2.19/angular.min.js"></script>
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


    </style>
<script>
var fMgr = angular.module('fMgr', ['tableSort', 'ui.bootstrap', 'ui.codemirror']);

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

    var queryServer = function(params) {
        return $http.post("/", params)
    }

    var on_error = function(data) {
        console.log(data)
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
        }

    }
})



fMgr.controller("ListCtr", function($scope, $http, $location, 
        $document, $window, $timeout, ServerCommand, Flash){


    // Config flash
    Flash.scope( $scope );


    $scope.Path = "[% .Path %]"
    $scope.view = 'main'
    
    function get_data() {
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

    $scope.AddFolder = function() {

        if( $scope.folder_popover == true )
            $scope.folder_popover = undefined
        else $scope.folder_popover = true

        // This is for handling close when clicks outside popover
        var closePopOver = function(e) {
            if (angular.element(e.target).is('#folder_pop *, #folder_pop')) {
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


    $scope.CreateFolder = function() {
        var folder = $scope.folder_filename
        $scope.folder_filename = undefined
        $scope.folder_popover = undefined

        if(!folder) {
            Flash.error("Provide a folder name")
            return
        }  

        ServerCommand.get({
            action: 'createFolder',
            params: { source: $scope.Path + folder }
        }, "Folder Created", get_data)
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
    






</script>
    <title>FileManager</title>
  </head>
  <body ng-controller="ListCtr">
  

  <div class="container"  style="position:relative">
    <div class="row">
        <div class="col-md-6">
        <h3 id="title">FileManager</h3>
        </div>
    </div>
    
    <div  ng-show="flash.message" id="fmessage" class="{{ flash.type }}" >
    <p>{{ flash.message }}</p>
  </div>

  </div>
<!-- controller -->
<div class="container">


<div class="row">

<div class="col-md-12">
    <ol class="breadcrumb">
        <li><a href="/"><span class="glyphicon glyphicon-home"> </span></a></li>
        <li ng-repeat="item in Rutas"><a href="{{ item.url }}">{{ item.name }}</a></li> 
    </ol>

</div>

</div>



<div class="row" ng-show="view=='main'">
<form id="upload_files" style="display:none" enctype="multipart/form-data">
<input type="file" id="file_upload" name="files" multiple style="display:none" >
</form>
 <form role="form">
    <div class="col-md-6">
        <span id="folder_pop">
            <button type="button" class="btn btn-info btn-sm" ng-click="AddFolder()"><span class="glyphicon glyphicon-folder-open"> </span> &nbsp; Add Folder</button>
            <div ng-show="folder_popover" class="popover bottom am-flip-x" style="top: 20px; left: -5px; display: block; width:200px">
                <div class="arrow"></div>
                <h3 class="popover-title">Folder Name</h3>
                <div class="popover-content">
                    <input type="text" class="form-control input-sm" 
                        ng-model="folder_filename" id="folder_field" 
                        tfocus="folder_popover">
                    <button type="button" ng-click="CreateFolder()" 
                        class="btn btn-info btn-sm pull-right" 

                        style="margin-top:5px; margin-bottom:5px">Add</button>
                </div>
            </div>
        </span>
        <button type="button" class="btn btn-info btn-sm" type="file" ng-click="AddFiles()"
        tooltip-placement="top" tooltip="Upload Multiple Files"><span class="glyphicon glyphicon-plus"> </span> &nbsp; Upload</button>
        
        <a class="btn btn-info btn-sm" target="_self" href="?format=zip" 
            tooltip-placement="top" tooltip="Download as Zip"><span class="glyphicon glyphicon-download-alt"> </span></a>
        &nbsp;&nbsp;&nbsp;
        <button class="btn btn-danger btn-sm" ng-show="selected>0" ng-click="DeleteSelected()"
            tooltip-placement="top" tooltip="Delete selected"><span class="glyphicon glyphicon-trash"> </span></button>


    </div>
    <div class="col-md-6 pull-right text-right">
    
            
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
            
            <input ng-model="query" class="form-control input-sm" placeholder="Filter" style="width:45%; float:right; margin-right:15px" >

    </div>
    </form>
</div>


<div class="row list" style="margin-top:15px" ng-show="view=='main'">
<div class="col-md-12">
    <table class="table" ts-wrapper>
        <thead>
            <tr>
                <td>&nbsp</td>
                <td ts-criteria="IsDir" class="f">+</td>
                <td ts-criteria="Name" ts-default="Ascending">Filename</td>
                <td ts-criteria="Size|parseFloat">Size</td>
                <td ts-criteria="ModTime">Last Modified</td>
                <td>Actions</td>
            </tr>
        </thead>
        <tbody>
            <tr ng-repeat="item in Files|filter:query|filter:{'IsDir':ff}|filter:{'IsHidden':hidden}" ts-repeat>
                <td width="20"><input type="checkbox" name="checkboxs[]" value="{{ item.Name }}" 
                    ng-click="CheckboxToggle(this, $event)" /></td>
                <td width="20"><span class="glyphicon glyphicon-folder-open" ng-show="item.IsDir"></span>
                <span ng-hide="item.IsDir" class="glyphicon glyphicon-file"></span></td>
                <td><a href="{{ Path }}{{ item.Name }}" target="_self" ng-if="!item.IsDir">{{ item.Name }}</a><a href="{{ item.Name }}/" ng-if="item.IsDir" class="dir">{{ item.Name }}</span></td>
                <td width="100">{{ item.Size/1024|number:0 }}Kb</td>
                <td width="140">{{ item.ModTime|date:'dd/MM/yyyy HH:mm:ss' }}</td>
                <td width="90">
                
                    
                    

                    <!--<a href="{{ item.Name }}/?format=zip" target="_self" ng-if="item.IsDir" 
                        class="glyphicon glyphicon-download-alt delete" tooltip-placement="top" tooltip="Donwload as Zip"> </a>-->
                    
                    <div class="btn-group" dropdown>
                        <a href class="dropdown-toggle">
                            <span class="glyphicon glyphicon-cog"> </span>
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

                    <span ng-if="item.IsText" ng-click="EditFile(item.Name)" class="glyphicon glyphicon-pencil delete" tooltip-placement="top" tooltip="Edit"> </span>

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
    <div class="col-md-6">
        <h4>Editing {{ currentEditedFile }}</h4>
    </div>
    <div class="col-md-6 pull-right text-right">
        <button type="button" class="btn btn-info btn-sm" ng-click="ToView('main')"><span class="glyphicon glyphicon-arrow-left"> </span> &nbsp; Back</button>
        <button type="button" class="btn btn-info btn-sm" ng-click="SaveFile()"><span class="glyphicon glyphicon-floppy-disk"> </span> &nbsp; Save</button>
        
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

<script>

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








</script>




</html>
`
