package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var dir string
var port string
var logging bool
var store = sessions.NewCookieStore([]byte("keysecret"))

const MAX_MEMORY = 1 * 1024 * 1024

func main() {

	flag.StringVar(&dir, "dir", ".", "Specify a directory to server files from.")
	flag.StringVar(&port, "port", ":8080", "Port to bind the file server")
	flag.BoolVar(&logging, "log", true, "Enable Log (true/false)")

	flag.Parse()

	if logging == false {
		log.SetOutput(ioutil.Discard)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handleReq))
	log.Printf("Listening.....")
	http.ListenAndServe(port, mux)

}

func handleReq(w http.ResponseWriter, r *http.Request) {

	//act := r.Values['action']
	//log.Printf("Request: %s", r.FormValue("action"))
	if r.FormValue("action") == "upload" {
		log.Printf("Uploading file")
		upload_file(w, r, r.URL.Path[1:])
		http.Redirect(w, r, r.URL.Path, http.StatusFound)
		return
	}

	if strings.HasSuffix(r.URL.Path, "/") {
		log.Printf("Index dir")
		handleDir(w, r)
	} else {
		log.Printf("downloading file %s", path.Clean(dir+r.URL.Path))
		http.ServeFile(w, r, path.Clean(dir+r.URL.Path))
		//http.ServeContent(w, r, r.URL.Path)
		//w.Write([]byte("this is a test inside file handler"))

	}

}

func handleDir(w http.ResponseWriter, r *http.Request) {

	var d string = "."

	//log.Printf("len %d,, %s", len(r.URL.Path), dir)
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

	// get flash messages?
	session, err := store.Get(r, "flash-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fm := session.Flashes("message")
	session.Save(r, w)
	//fmt.Fprintf(w, "%v", fm[0])

	t := template.Must(template.New("listing").Parse(templateList))
	v := map[string]interface{}{
		"Title":   d,
		"Listing": template.HTML(out),
		"Path":    r.URL.Path,
		"notroot": len(r.URL.Path) > 1,
		"message": fm,
	}

	t.Execute(w, v)

	//w.Write([]byte("this is a test inside dir handle"))
}

func upload_file(w http.ResponseWriter, r *http.Request, p string) {
	if err := r.ParseMultipartForm(MAX_MEMORY); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	for key, value := range r.MultipartForm.Value {
		//fmt.Fprintf(w, "%s:%s", key, value)
		log.Printf("%s:%s", key, value)
	}

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			log.Println(fileHeader.Filename)
			p := p + fileHeader.Filename
			buf, _ := ioutil.ReadAll(file)
			e := ioutil.WriteFile(p, buf, os.ModePerm)
			if e != nil {
				http.Error(w, e.Error(), http.StatusForbidden)
			}
		}
	}

	// flash message
	session, err := store.Get(r, "flash-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	session.AddFlash("File successfull uploaded", "message")
	session.Save(r, w)

}
