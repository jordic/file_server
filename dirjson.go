package main

import (
	"encoding/json"
	"github.com/jordic/file_server/util"
	"io"
	"os"
	"time"
)

type File struct {
	Name    string
	Size    int64
	ModTime time.Time
	IsDir   bool
	IsText  bool
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
