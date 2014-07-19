package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Mod time will be init, every time, you start the server,
// Becasuse there is not an easy way of determining the timestamp of compilation
// or I don't know it.
var ModTime = time.Now()

type ServeStaticFromBinary struct {
	MountPoint string
	DataDir    string // data/ with end slash
}

func (s *ServeStaticFromBinary) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	file := r.URL.Path[len(s.MountPoint):]
	// requires go-bindata assets generation
	data, err := Asset(s.DataDir + file)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.ServeContent(w, r, file, ModTime, NewAssetDownload(data))
}

// AssetDownload for implementing reader, and closer
type AssetDownload struct {
	*bytes.Reader
	io.Closer
}

func NewAssetDownload(a []byte) *AssetDownload {
	return &AssetDownload{
		bytes.NewReader(a),
		ioutil.NopCloser(nil),
	}
}
