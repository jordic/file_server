package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

// data_ds_store reads file data from disk. It returns an error on failure.
func data_ds_store() ([]byte, error) {
	return bindata_read(
		"/Users/jordi/Documents/projectes/go/src/github.com/jordic/file_server/data/.DS_Store",
		"data/.DS_Store",
	)
}

// data_app_css reads file data from disk. It returns an error on failure.
func data_app_css() ([]byte, error) {
	return bindata_read(
		"/Users/jordi/Documents/projectes/go/src/github.com/jordic/file_server/data/app.css",
		"data/app.css",
	)
}

// data_app_js reads file data from disk. It returns an error on failure.
func data_app_js() ([]byte, error) {
	return bindata_read(
		"/Users/jordi/Documents/projectes/go/src/github.com/jordic/file_server/data/app.js",
		"data/app.js",
	)
}

// data_libs_js reads file data from disk. It returns an error on failure.
func data_libs_js() ([]byte, error) {
	return bindata_read(
		"/Users/jordi/Documents/projectes/go/src/github.com/jordic/file_server/data/libs.js",
		"data/libs.js",
	)
}

// data_main_html reads file data from disk. It returns an error on failure.
func data_main_html() ([]byte, error) {
	return bindata_read(
		"/Users/jordi/Documents/projectes/go/src/github.com/jordic/file_server/data/main.html",
		"data/main.html",
	)
}

// data_styles_css reads file data from disk. It returns an error on failure.
func data_styles_css() ([]byte, error) {
	return bindata_read(
		"/Users/jordi/Documents/projectes/go/src/github.com/jordic/file_server/data/styles.css",
		"data/styles.css",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"data/.DS_Store": data_ds_store,
	"data/app.css": data_app_css,
	"data/app.js": data_app_js,
	"data/libs.js": data_libs_js,
	"data/main.html": data_main_html,
	"data/styles.css": data_styles_css,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
func AssetDir(name string) ([]string, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	pathList := strings.Split(cannonicalName, "/")
	node := _bintree
	for _, p := range pathList {
		node = node.Children[p]
		if node == nil {
			return nil, fmt.Errorf("Asset %s not found", name)
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"data": &_bintree_t{nil, map[string]*_bintree_t{
		".DS_Store": &_bintree_t{data_ds_store, map[string]*_bintree_t{
		}},
		"app.css": &_bintree_t{data_app_css, map[string]*_bintree_t{
		}},
		"app.js": &_bintree_t{data_app_js, map[string]*_bintree_t{
		}},
		"libs.js": &_bintree_t{data_libs_js, map[string]*_bintree_t{
		}},
		"main.html": &_bintree_t{data_main_html, map[string]*_bintree_t{
		}},
		"styles.css": &_bintree_t{data_styles_css, map[string]*_bintree_t{
		}},
	}},
}}
