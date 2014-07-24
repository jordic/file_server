



var ToView;





fMgr.controller("ListCtr", function($scope, $http, $location, 
        $document, $window, $timeout, ServerCommand, Flash, 
        ngDialog, Path){


    $scope.IsMobile = window.is_mob()
    $scope.SysCmdDisabled = DISABLE_SYS_COMMAND

    // Config flash
    Flash.scope( $scope );


    $scope.Path = APP_PATH;
    $scope.view = 'main'
    
    var get_data = function() {

        // @todo 
        $scope.data_loading = true

        $http.get($scope.Path + "?format=json")
            .then(function(res){
                $scope.data_loading = false
                //console.log(res)
                if(res.data=="null")
                    $scope.Files = []
                else
                    $scope.Files = res.data
              
                $scope.iFiles = {}
                $scope.idFiles = {}
                var idc = 0;
                angular.forEach($scope.Files, function(item){
                    
                    $scope.iFiles[item.Name] = item
                    item.ModTime = new Date(item.ModTime)
                    if(item.Name.indexOf(".") === 0) item.IsHidden = true;
                    else item.IsHidden = false

                    item.Id = idc++;
                    $scope.idFiles[item.Id] = item

                })
                $scope.Rutas = GetRuta()
                $scope.selected = 0


            }, function(error){
                Flash.duration(10000).error( "Server Disconnected" )
            })
        $scope.view = 'main'
        $scope.query = ''
        
        Path.Set($scope.Path);
        $scope.Cursor = 0


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


    
    var ieditor

    // Editor
    $scope.editorOptions = {
        lineWrapping : true,
        lineNumbers: true,
        keyMap: 'sublime',
        //theme: 'default'
        theme: 'monokai',
        extraKeys: {
            Esc: function() {
                //console.log('key', $scope)
                //$scope.ToView('main')
                $('#editor_back').trigger('click')
            },
            "Ctrl-S": function(){
                $('#editor_save').trigger('click')
            }
        }
        //mode: 'javascript'
    };

    $scope.EditFile = function(item) {
        $scope.currentEditedFile = $scope.Path + item
        var file = $scope.Path + item
        var noJsonTransform = function(data) { return data; };
        
        var m = determine_editor_mode(item, $scope.Path)
        //ieditor.setOption("mode", m)

        //console.log(CodeMirror.modes)

        $http.get( file, {
            transformResponse: noJsonTransform}).then(function(d){
            //console.log(d)
            $scope.EditorCurrentContent = d.data
            $scope.EditorOldContent = $scope.EditorCurrentContent
            $scope.EditorRefresh = true
            $scope.view = 'edit'
            
            //console.log(mode)
            //$scope.editorOptions.mode = mode
            $scope.EditorInstance.setOption("mode", m)
            //$scope.$apply()
            $timeout(function() {
                $scope.EditorInstance.refresh()
                $scope.EditorInstance.focus()
            }, 0)

            
        })
        
    }
    var _ed
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

    //ToView = $scope.ToView;

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

    $scope.ShowCommands = function() {
        ngDialog.open({ template: '/exec_command.html' });
    }


    // Key bindings
    $('#finder input').bind('blur', function(e){
        console.log('bluf finder')
        angular.element('body').trigger('click')
    })


    angular.element($window).on('keydown', function(e) {
        
            //console.log(e)
            if (e.target.nodeName == "INPUT")
                return

            //console.log(e.keyCode)
            switch(e.keyCode) {
                case 74: //j
                    $scope.Cursor++
                    break;
                case 75: //k
                    $scope.Cursor--
                    break;
                case 79: //o
                    var item = ItemFromCursor()
                    if(item.IsDir)
                        $location.path( $scope.Path + item.Name + "/" )      
                    break;
                case 69: // e
                    var item = ItemFromCursor()
                    if(item.IsText)
                        $scope.EditFile( item.Name )
                    break;
                case 88: // x
                    var item = ItemFromCursor()
                    var n = item.Id
                    angular.element('#chk_'+n).trigger('click')
                    break;
                case 68: //d
                    if($scope.selected > 0) {
                        $scope.DeleteSelected()
                    }
                    break;
                case 83: // s
                    angular.element('#finder button').trigger('click')
                    break;    
                case 72: // s
                    $location.path("/")
                    break;    

            }
            var total = angular.element('tr').length - 2
            if($scope.Cursor == total)
                $scope.Cursor = 0
            if($scope.Cursor == -1)
                $scope.Cursor = total - 1
            $scope.$apply()
            //$scope.MoveCursor()
    
    });
    $scope.Cursor = 0
    

    var ItemFromCursor = function() {
        var f = angular.element('#r'+$scope.Cursor),
        name = f.attr("_name")
        return $scope.iFiles[name]
    }

    
})

fMgr.controller("ExecCtrl", function($scope, $http, $location, Path, 
        ServerCommand){
        
        $scope.Path = Path.Get()
        $scope.command = ''
        $scope.params = ''

        $scope.Execute = function() {

            $scope.is_executing = true

            if( $scope.command == "") {
                alert("Must supply a command")
                return
            }

            var p = {
                action: 'syscmd', 
                params: {
                    'source': $scope.Path,
                    'command': $scope.command
                },
                paramslist: $scope.params.split(" ")
            }   

            ServerCommand.get_raw(p).then(function(d){
                $scope.is_executing = false
                $scope.output = d.data.message
            })


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



function determine_editor_mode( file, sp ) {

    if(file.indexOf('.js') != -1 || file.indexOf('.json') != -1 )
        return "javascript";

    //if(file.indexOf('.php') != -1) return "php";
    //if(file.indexOf('.htm') != -1) return "htmlmixed";
    if(file.indexOf('.css') != -1) return "css";
    if(file.indexOf('.php') != -1) return "php";
    if(file.indexOf('.html') != -1) return "htmlmixed";
    if(file.indexOf('.md') != -1) return "markdown";

    var test = sp + file
    if(test.indexOf('nginx') != -1) return "nginx";

}


