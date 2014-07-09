package main

import (
	"archive/zip"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	//	"runtime/pprof"
	_ "net/http/pprof"
	"strings"
	"time"
)

var dir string
var port string
var logging bool
var store = sessions.NewCookieStore([]byte("keysecret"))

//var cpuprof string

const MAX_MEMORY = 1 * 1024 * 1024
const VERSION = "0.91a"

type File struct {
	Name    string
	Size    int64
	ModTime time.Time
	IsDir   bool
}

/*
@ TODO
download zip folders
https://bitbucket.org/kardianos/staticserv/src/5a536ebb8016d795187138ad99881533e14a59ef/main.go?at=default
*/

func main() {

	//fmt.Println(len(os.Args), os.Args)
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		fmt.Println("Version " + VERSION)
		os.Exit(0)
	}

	flag.StringVar(&dir, "dir", ".", "Specify a directory to server files from.")
	flag.StringVar(&port, "port", ":8080", "Port to bind the file server")
	flag.BoolVar(&logging, "log", true, "Enable Log (true/false)")
	//flag.StringVar(&cpuprof, "cpuprof", "", "write cpu and mem profile")

	flag.Parse()

	/*
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
		if cpuprof != "" {
			f, err := os.Create(cpuprof)
			if err != nil {
				log.Fatal(err)
			}
			pprof.StartCPUProfile(f)
			defer pprof.StopCPUProfile()
		}*/

	if logging == false {
		log.SetOutput(ioutil.Discard)
	}
	// If no path is passed to app, normalize to path formath
	if dir == "." {
		dir, _ = filepath.Abs(dir)
	}

	if _, err := os.Stat(dir); err != nil {
		log.Fatalf("Directory %s not exist", dir)
	}

	// normalize dir, ending with... /
	if strings.HasSuffix(dir, "/") == false {
		dir = dir + "/"
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handleReq))
	log.Printf("Listening on port %s .....", port)
	http.ListenAndServe(port, mux)

}

func handleReq(w http.ResponseWriter, r *http.Request) {

	// get dir / app>template
	//	get dir /.zip
	//  get dir /.json

	//fmt.Printf("ajax: %s\n", r.Header.Get("angular"))
	//fmt.Printf("ajax: %s\n", r.Header.Get("Accept"))
	//Is_Ajax := strings.Contains(r.Header.Get("Accept"), "application/json")
	//fmt.Printf("is_ajax %s\n", )

	if r.Method == "PUT" {
		AjaxUpload(w, r)
		return
	}

	if r.Method == "POST" {
		SaveFile(w, r)
		return
	}

	if r.FormValue("ajax") == "true" {
		AjaxActions(w, r)
		return
	}

	if strings.HasSuffix(r.URL.Path, "/") {
		log.Printf("Index dir %s", r.URL.Path)
		handleDir(w, r)
	} else {
		log.Printf("downloading file %s", path.Clean(dir+r.URL.Path))
		http.ServeFile(w, r, path.Clean(dir+r.URL.Path))
		//http.ServeContent(w, r, r.URL.Path)
		//w.Write([]byte("this is a test inside file handler"))

	}

}

func SaveFile(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t map[string]string
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//fmt.Printf("%#v", t)
	//fmt.Print(t["file"])
	f := strings.Trim(t["file"], "/")
	data := []byte(t["content"])
	err = ioutil.WriteFile(dir+f, data, 0644)
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "ok")
	return

}

func AjaxActions(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("action") == "save" {
		f := strings.Trim(r.FormValue("file"), "/")
		data := []byte(r.FormValue("content"))

		err := ioutil.WriteFile(dir+f, data, 0644)
		if err != nil {
			fmt.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, "ok")
		return
	}

	if r.FormValue("action") == "delete" {
		f := strings.Trim(r.FormValue("file"), "/")
		err := os.Remove(dir + f)
		if err != nil {
			fmt.Fprint(w, fmt.Sprintf("Error %s", err))
			return
		}
		fmt.Fprint(w, fmt.Sprintf("ok"))
		return
	}

	if r.FormValue("action") == "deleteList" {
		var l = make([]string, 0)
		list := r.FormValue("files")
		err := json.Unmarshal([]byte(list), &l)
		if err != nil {
			fmt.Fprint(w, err)
		}

		//fmt.Printf("list %#v %#v", l, list)
		for k := range l {
			f := strings.Trim(l[k], "/")
			fmt.Println(l[k])
			err := os.Remove(dir + f)
			if err != nil {
				fmt.Fprint(w, fmt.Sprintf("Error %s", err))
			}
		}

		fmt.Fprint(w, "ok")
		return
	}

	if r.FormValue("action") == "create_folder" {
		f := strings.Trim(r.FormValue("path"), "/") + "/" + r.FormValue("file")
		err := os.Mkdir(dir+f, 0777)
		if err != nil {
			fmt.Fprint(w, fmt.Sprintf("Error %s", err))
			return
		}
		fmt.Fprint(w, fmt.Sprintf("ok"))
		return
	}

	// @todo make some test cases of trim renames...
	if r.FormValue("action") == "rename" {
		fo := strings.Trim(r.FormValue("file"), "/")
		fn := strings.Trim(r.FormValue("new"), "/")
		fo = strings.Trim(fo, "../")
		fn = strings.Trim(fn, "../")
		//fmt.Printf("Old %s new %s", fo, fn)

		if dir != "." {
			fo = dir + fo
			fn = dir + fn
		}

		err := os.Rename(fo, fn)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		fmt.Fprint(w, "ok")
		return
	}

}

func handleDir(w http.ResponseWriter, r *http.Request) {

	var d string = ""

	//log.Printf("len %d,, %s", len(r.URL.Path), dir)
	if len(r.URL.Path) == 1 {
		// handle root dir
		d = dir
	} else {
		d += dir + r.URL.Path[1:]
	}

	// handle json format of dir...
	if r.FormValue("format") == "json" {
		result := &DirJson{w, d}
		err := result.Get()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.FormValue("format") == "zip" {
		result := &DirZip{w, d}
		err := result.Get()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// If we dont receive json param... we are asking, for genric app ui...
	t := template.Must(template.New("listing").Delims("[%", "%]").Parse(templateList))
	v := map[string]interface{}{
		"Path":    r.URL.Path,
		"version": VERSION,
	}

	t.Execute(w, v)

}

func AjaxUpload(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pa := r.URL.Path[1:]

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		var ff string
		if dir != "." {
			ff = dir + pa + part.FileName()
		} else {
			ff = pa + part.FileName()
		}

		dst, err := os.Create(ff)
		defer dst.Close()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprint(w, "ok")
	return
}

// DirJson handle dir listings in json format
type DirJson struct {
	w io.Writer
	d string
}

// Get Writes to stdout, json string...
func (t *DirJson) Get() error {

	//fmt.Print(t.d)
	thedir, err := os.Open(t.d)
	if err != nil {
		return err
	}
	defer thedir.Close()

	finfo, err := thedir.Readdir(-1)
	if err != nil {
		return err
	}

	var aout []*File

	for _, fi := range finfo {
		xf := &File{
			fi.Name(),
			fi.Size(),
			fi.ModTime(),
			fi.IsDir(),
		}

		// if is a symlink ... follow it to test if is a real dir...
		if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
			//fmt.Printf("%s is a symlink", xf.Name)
			fx, err := os.Readlink(t.d + fi.Name())
			if err != nil {
				continue
			}
			fxi, err := os.Stat(fx)
			if err != nil {
				continue
			}
			//fmt.Printf("%s is a dir %#v\n", t.d+fi.Name(), fxi.IsDir())
			// If is a dir, populate object with it.
			if fxi.IsDir() {
				xf.IsDir = true
			}

		}

		aout = append(aout, xf)
	}

	xo, err := json.Marshal(aout)
	if err != nil {
		return err
	}
	t.w.Write(xo)
	return nil

}

type DirZip struct {
	w http.ResponseWriter
	d string
}

func (t *DirZip) Get() error {

	zipFileName := fmt.Sprintf("%s.zip", filepath.Base(t.d))
	t.w.Header().Set("Content-Type", "application/zip")
	t.w.Header().Set("Content-Disposition", `attachment; filename="`+zipFileName+`"`)

	zw := zip.NewWriter(t.w)
	defer zw.Close()

	filepath.Walk(t.d, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}
		zipPath := path[len(t.d):]
		zipPath = strings.TrimLeft(strings.Replace(zipPath, `\`, "/", -1), `/`)
		ze, err := zw.Create(zipPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot create zip entry <%s>: %s\n", zipPath, err)
			return err
		}
		file, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open file <%s>: %s\n", path, err)
			return err
		}
		defer file.Close()

		io.Copy(ze, file)
		return nil

	})

	return nil

}

/*
// commands
supervisorctl restart xxxx
service nginx restart
service mysql restart
mv $1 $2
tar xvfz name.tar.gz asdf/
git pull origin master
git push origin master
git commit -am "message"
bin/sqldumpr -xxxx


*/
