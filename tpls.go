package main

const templateList = `
<!DOCTYPE html>
<html lang="en" ng-app="fMgr">
  <head>
    <meta charset="utf-8">
    <link href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css" 
        rel="stylesheet" />
    
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
    p.bg-success { padding:5px; color:green; }

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
    #fmessage { position:absolute; z-index:50; width:30%;  top:0px; display:inline-block; right:15px; }


    </style>
<script>
var fMgr = angular.module('fMgr', ['tableSort', 'ui.bootstrap']);

fMgr.config(['$locationProvider', function($locationProvider){
    $locationProvider.html5Mode(true);
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



fMgr.controller("ListCtr", function($scope, $http, $location, $document, $window, $timeout){

    $scope.Path = "[% .Path %]"
    
    
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
        })
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
         if( $scope.Path != newl ) {
            $scope.Path = newl
            get_data()
         }
         
    })

    $scope.DeleteFile = function(item) {
        var res = confirm("Are you sure?")
        if(res) {
            //console.log( $scope.Path + item )
            $http.get("/", {params:{
                "ajax": "true",
                "action": "delete",
                "file": $scope.Path + item
            }}).then(function(d){
                if(d.data == "ok") {
                    Flash_Message("bg-success", "File deleted")
                    //$location.path( $scope.Path )
                    get_data()
                } else {
                    Flash_Message("bg-danger", d.data, 5000)
                }
            })
        }
    }

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

    $scope.CreateFolder = function() {
        var folder = $scope.folder_filename
        $scope.folder_filename = undefined
        $scope.folder_popover = undefined

        if(!folder) {

            Flash_Message("bg-danger", "Provide a folder name")
            return
        }  

        $http.get("/", {params:{
                "ajax": "true",
                "action": "create_folder",
                "file": folder,
                "path": $scope.Path
            
            }}).then(function(d){
                if(d.data == "ok") {
                    Flash_Message("bg-success", "Folder Created")
                    //$location.path( $scope.Path )
                    get_data()
                } else {
                    Flash_Message("bg-danger", d.data, 5000)
                }
            })

    }

    function Flash_Message(type, msg, time) {

        t = 3000
        if(time) t = time;

        $scope.flash = { type:type, message:msg }
        $timeout(function(){ 
            $scope.flash = undefined
         }, t)
    }


    $scope.AddFiles = function() {

        var uploadFiles = function(evt){
            var files = evt.target.files;
            var formData = new FormData( document.getElementById('upload_files') );
           
            //formData.append("action", "upload")
            //formData.append("is_ajax", "true")
            var xhr = new XMLHttpRequest();
            xhr.open('PUT', $scope.Path, true);
            //xhr.setRequestHeader("Content-Type","multipart/form-data;")
            xhr.onload = function(e) {
                //console.log(e)
                document.getElementById('file_upload').removeEventListener('change', uploadFiles);
                Flash_Message("bg-success", "File uploaded")
                get_data()
             };

             xhr.onerror = function(e) {
                console.log(e)
             }

            xhr.upload.onprogress = function(e) {
                console.log(e)
                if (e.lengthComputable) {
                    console.log( (e.loaded / e.total) * 100 )
                }
             }

             xhr.send(formData)

        };


        document.getElementById('file_upload').addEventListener('change', uploadFiles, false)
        document.getElementById('file_upload').click();
        //console.log("file upload")
    }


    $scope.RenameFile = function(f) {
        var old_path =  $scope.Path +  f
        var res = prompt("Rename/Move File?", f)
        if(res ) {
                


            Flash_Message("bg-success", "File "+ res + " renamed")
            return
        } else {
            return
        }
        
    }


})
    



</script>
    <title>FileManager</title>
  </head>
  <body ng-controller="ListCtr">

  <div class="container" style="position:relative">
    <div class="row">
        <div class="col-md-6">
        <h3 id="title">FileManager</h3>
        </div>
    </div>
    <p ng-show="flash.message" class="{{ flash.type }} text-center" id="fmessage">{{ flash.message }}</p>
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



<div class="row">
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


<div class="row list" style="margin-top:15px">
<div class="col-md-12">
    <table class="table" ts-wrapper>
        <thead>
            <tr>
                <td ts-criteria="IsDir" class="f">+</td>
                <td ts-criteria="Name" ts-default="Ascending">Filename</td>
                <td ts-criteria="Size|parseFloat">Size</td>
                <td ts-criteria="ModTime">Last Modified</td>
                <td>Actions</td>
            </tr>
        </thead>
        <tbody>
            <tr ng-repeat="item in Files|filter:query|filter:{'IsDir':ff}|filter:{'IsHidden':hidden}" ts-repeat>
                <td width="20"><span class="glyphicon glyphicon-folder-open" ng-show="item.IsDir"></span>
                <span ng-hide="item.IsDir" class="glyphicon glyphicon-file"></span></td>
                <td><a href="{{ Path }}{{ item.Name }}" target="_self" ng-if="!item.IsDir">{{ item.Name }}</a><a href="{{ item.Name }}/" ng-if="item.IsDir" class="dir">{{ item.Name }}</span></td>
                <td width="100">{{ item.Size/1024|number:0 }}Kb</td>
                <td width="140">{{ item.ModTime|date:'dd/MM/yyyy HH:mm:ss' }}</td>
                <td width="60">
                    <span ng-click="RenameFile(item.Name)" class="glyphicon glyphicon-pencil delete"> </span>
                    <span ng-click="DeleteFile(item.Name)" class="glyphicon glyphicon-trash delete"> </span>
                    
                </td>
            </tr>
        </tbody>

    </table>
</div>
</div>

</div>

<div class="container">
    <div class="row" style="margin-top:200px; border-top:1px solid #eaeaea; padding-top:20px; font-size:10px">
        <p><a href="http://github.com/jordic/file_server">http://github.com/jordic/file_server</a> -- v.[% .version %]
        </p>
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









</script>

</html>
`
