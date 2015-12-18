package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// For handling compress dir structures...
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
