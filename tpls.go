package main

const templateList = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <link href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css" 
        rel="stylesheet" />
    
    <script src="https://code.jquery.com/jquery-1.11.0.min.js"> </script>
    <script src="//netdna.bootstrapcdn.com/bootstrap/3.1.1/js/bootstrap.min.js"> </script>

    <style>
    body { font-family:Arial; font-size:14px }
    .glyphicon { margin-right:5px; color:grey; font-size:14px }
    #file_tree { line-height:18px }
    .dir { color:green }
    .breadcrumb>li+li:before { content:" "; }
    #file_list li { padding-bottom:8px; margin-bottom:8px; list-style: none; border-bottom: 1px solid #eaeaea; }
    p.bg-success { padding:5px; color:green; }
    </style>
    <script>
    $(document).ready(function(){
    $('.delete').click(function(){
        if(confirm("Are you sure?"))
            return true
        return false
    })
    })
    $('.bg-success').click(function(){
        $(this).hide()
    })

    </script>
    <title>{{ .Path }}</title>
  </head>
  <body>
  <div class="container">
    <div class="row">
        <div class="col-md-6">
        <h3>Mini WebFileManager</h3>
        </div>
    </div>
    <div class="row">
    <ol class="breadcrumb">
        <li><a href="/"><span class="glyphicon glyphicon-home"> </span></a></li>
        <li>{{ .Path }}</li>
        {{if .notroot }}
        <li><a href="../">..</a></li>
        {{ end }}
    </ol>
    </div>

    {{if .message }}
    <div class="row">
        <div class="col-md-12" role="main">
            <p class="bg-success">{{ .message }}</p>
        </div>
    </div>
    {{ end }}

    <div class="row">
    <div class="col-md-8">
    <form action="" role="form" method="POST" class="form-inline" enctype="multipart/form-data">
        <input type="hidden" value="upload" name="action" />
        <div class="form-group">
            <label for="ff">Upload a file</label>
            <input type="file" name="file" class="form-control" id="ff" placeholder="Choose your file">
        </div>
        <div class="form-group">
            <button type="submit" class="btn btn-primary" >Upload</button>
        </div>
    </form>
    </div>
    </div>
    <div class="row" style="margin-top:20px" id="file_tree">



        <div class="col-md-12" id="file_list">
        {{.Listing}}
        </div>
    </div>
    

    <div class="row" style="margin-top:200px; border-top:1px solid #eaeaea; padding-top:20px; font-size:10px">
        <p><a href="http://github.com/jordic/file_server">http://github.com/jordic/file_server</a> -- v.{{ .version }}
        </p>
    </div>
    </div>

  </body>
</html>
`
