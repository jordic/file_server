package main

import (
	"encoding/json"
	"fmt"
	"github.com/jordic/fuzzyfs"
	"log"
	"net/http"
	"time"
)

var dirlist *fuzzyfs.DirList

func init() {

	dirlist = fuzzyfs.NewDirList()
	dirlist.MaxDepth = 4
	dirlist.PathSelect = fuzzyfs.DirsAndSymlinksAsDirs

}

// Build index, starts when app start...
func Build_index(path string) {
	startTime := time.Now()
	log.Printf("Building index .. %s", path)
	err := dirlist.Populate(path, nil)
	if err != nil {
		panic(err)
	}
	endTime := time.Now()
	log.Printf("%d entries. time index .. %s", dirlist.Length(), endTime.Sub(startTime))

}

func SearchHandle(w http.ResponseWriter, r *http.Request) {

	var query string
	query = r.FormValue("q")

	if len(query) < 3 {
		fmt.Fprint(w, "[]")
		return
	}

	res := dirlist.Query(query, 2)
	//fmt.Printf("%#v", res)
	cut := len(res)
	if cut > 10 {
		cut = 9
	}

	out, err := json.Marshal(res[0:cut])

	if err != nil {
		log.Printf("error encoding results... ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	//fmt.Fprint(w, out)
	w.Write(out)
	return

}
