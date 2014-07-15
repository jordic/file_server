package main

import (
	"archive/zip"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jordic/file_server/util"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var (
	dir     string
	port    string
	logging bool
	depth   int
	// directory indexin

)

//var store = sessions.NewCookieStore([]byte("keysecret"))

//var cpuprof string

const MAX_MEMORY = 1 * 1024 * 1024
const VERSION = "0.93b"

type File struct {
	Name    string
	Size    int64
	ModTime time.Time
	IsDir   bool
	IsText  bool
}

func main() {

	//fmt.Println(len(os.Args), os.Args)
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		fmt.Println("Version " + VERSION)
		os.Exit(0)
	}

	flag.StringVar(&dir, "dir", ".", "Specify a directory to server files from.")
	flag.StringVar(&port, "port", ":8080", "Port to bind the file server")
	flag.BoolVar(&logging, "log", true, "Enable Log (true/false)")
	//flag.IntVar(&depth, "depth", 5, "Depth directory crwaler")

	//flag.StringVar(&cpuprof, "cpuprof", "", "write cpu and mem profile")

	flag.Parse()

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

	// build index files in background
	go Build_index(dir)

	mux := http.NewServeMux()
	//mux.Handle("/-/libs.js", http.ServeFile(w, r, Asset("libs.js")))

	mux.Handle("/-/assets/", http.HandlerFunc(serve_statics))

	mux.Handle("/-/api/dirs", makeGzipHandler(http.HandlerFunc(SearchHandle)))

	mux.Handle("/", makeGzipHandler(http.HandlerFunc(handleReq)))
	log.Printf("Listening on port %s .....", port)
	http.ListenAndServe(port, mux)

}

func serve_statics(w http.ResponseWriter, r *http.Request) {

	file := r.URL.Path[10:]
	by, err := Asset("data/" + file)

	if strings.Contains(file, "css") == true {
		w.Header().Set("Content-Type", "text/css")
	}

	if strings.Contains(file, "js") == true {
		w.Header().Set("Content-Type", "text/javascript")
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(by)
	return
}

func handleReq(w http.ResponseWriter, r *http.Request) {

	//Is_Ajax := strings.Contains(r.Header.Get("Accept"), "application/json")
	if r.Method == "PUT" {
		AjaxUpload(w, r)
		return
	}
	if r.Method == "POST" {
		WebCommandHandler(w, r)
		return
	}

	if strings.HasSuffix(r.URL.Path, "/") {
		log.Printf("Index dir %s", r.URL.Path)
		handleDir(w, r)
	} else {
		log.Printf("downloading file %s", path.Clean(dir+r.URL.Path))
		r.Header.Del("If-Modified-Since")
		http.ServeFile(w, r, path.Clean(dir+r.URL.Path))
		//http.ServeContent(w, r, r.URL.Path)
		//w.Write([]byte("this is a test inside file handler"))

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

		w.Header().Set("Content-Type", "application/json")
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
	template_file, err := Asset("data/main.html")
	if err != nil {
		log.Fatalf("Cant load template main")
	}

	t := template.Must(template.New("listing").Delims("[%", "%]").Parse(string(template_file)))
	v := map[string]interface{}{
		"Path":    r.URL.Path,
		"version": VERSION,
	}
	w.Header().Set("Content-Type", "text/html")
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
			false,
		}
		// detect is if text file
		if fi.IsDir() == false {
			xf.IsText = util.IsTextFile(t.d + fi.Name())
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
