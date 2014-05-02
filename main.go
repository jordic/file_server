package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var dir string
var port string

const MAX_MEMORY = 1 * 1024 * 1024

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
    body { font-family:Arial; font-size:12px }
    .glyphicon { margin-right:5px; color:grey; font-size:14px }
    #file_tree { line-height:18px }
    .dir { color:green }
    </style>
    <title>Http File server: {{.Title}}</title>
  </head>
  <body>
  <div class="container">
    <div class="row">
        <div class="col-md-6">
        <h3>Listing files</h3>
        </div>
    </div>
    <div class="row">
    <div class="col-md-6">
    <form action="" role="form" method="POST" class="form-inline" enctype="multipart/form-data">
        <input type="hidden" value="upload" name="action" />
        <div class="form-group">
            <label for="ff">Upload a file</label>
            <input type="file" class="form-control" id="ff" placeholder="Choose your file">
        </div>
        <div class="form-group">
            <button type="submit" class="btn btn-primary" >Upload</button>
        </div>
    </form>
    </div>
    </div>
    <div class="row" style="margin-top:20px" id="file_tree">
        <div class="col-md-12">
        {{.Listing}}
        </div>
    </div>
    </div>
  </body>
</html>
`

func main() {

	flag.StringVar(&dir, "dir", ".", "Specify a directory to server files from.")
	flag.StringVar(&port, "port", ":8080", "Port to bind the file server")

	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handleReq))
	log.Printf("Listening.....")
	http.ListenAndServe(port, mux)

}

func handleReq(w http.ResponseWriter, r *http.Request) {

	//act := r.Values['action']
	log.Printf("Request: %s", r.FormValue("action"))
	if r.FormValue("action") == "upload" {
		log.Printf("Uploading file")
	}

	if strings.HasSuffix(r.URL.Path, "/") {
		handleDir(w, r)
	} else {
		log.Printf("descargando archivo %s", path.Clean(dir+r.URL.Path))
		http.ServeFile(w, r, path.Clean(dir+r.URL.Path))
		//http.ServeContent(w, r, r.URL.Path)
		//w.Write([]byte("this is a test inside file handler"))

	}

}

func handleDir(w http.ResponseWriter, r *http.Request) {

	var d string = "."

	log.Printf("len %d,, %s", len(r.URL.Path), dir)
	if len(r.URL.Path) == 1 {
		// handle root dir
		d = dir

	} else {
		// @todo convert pahts to absolutes
		d += r.URL.Path
	}

	thedir, err := os.Open(d)
	if err != nil {
		// not a directory, handle a 404
		http.Error(w, "Page not found", 404)
		return
	}
	defer thedir.Close()

	finfo, err := thedir.Readdir(-1)
	if err != nil {
		return
	}

	out := ""
	for _, fi := range finfo {
		//log.Println(fi.Name())
		class := "file glyphicon glyphicon-file"
		name := fi.Name()
		if fi.IsDir() {
			class = "dir glyphicon glyphicon-folder-open"
			name += "/"
		}
		out += fmt.Sprintf("<a href='%s'><span class='%s'></span> %s</a><br />", name, class, name)
	}

	t := template.Must(template.New("listing").Parse(templateList))

	v := map[string]interface{}{
		"Title":   d,
		"Listing": template.HTML(out),
	}

	t.Execute(w, v)

	//w.Write([]byte("this is a test inside dir handle"))
}
