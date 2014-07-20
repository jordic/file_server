package main

import (
	//"errors"
	//"io"
	"io/ioutil"
	"os"
	"os/exec"
	//"strconv"
	"fmt"
	"github.com/opesun/copyrecur"
	"path/filepath"
	"strings"
)

//type params map[string]string
// Command represents an action on system.
type Command struct {
	Name       string `json:"name"`
	Params     map[string]string
	ParamsList []string
	Path       string
	run        func(c *Command) int
	Stdout     string
	Stderr     string
	status     int
}

func (c *Command) Run() int {
	return c.run(c)
}

// Returns a path ending with "/"
func (c *Command) GetPath() string {
	return strings.TrimRight(c.Path, "/") + "/"
}

// GetPParam gets a path param, trimming, the first /
func (c *Command) GetPathParam(s string) string {
	return strings.Trim(c.Params[s], "/")
}

func (c *Command) Status() int {
	return c.status
}

// Returns a command to exec in a given path ( dir )
func GetCommand(cmd string, dir string) *Command {
	if val, ok := commands[cmd]; ok {
		val.Path = dir
		return val
	}
	return nil
}

var commands = map[string]*Command{
	"save":         saveCommand,
	"delete":       deleteCommand,
	"createFolder": createfolderCommand,
	"rename":       renameCommand,
	"copy":         copyCommand,
	"compress":     compressCommand,
	"mv":           mvCommand,
	//"sys":          sysCommand,
}

var copyCommand = &Command{
	Name: "Copy",
	run:  copy_file,
}

func copy_file(c *Command) int {

	source := c.GetPath() + c.GetPathParam("source")
	dest := c.GetPath() + c.GetPathParam("dest")

	fi, err := os.Stat(source)
	if err != nil {
		c.Stderr = err.Error()
		c.status = 1
		return 1
	}

	if fi.IsDir() {
		err := copyrecur.CopyDir(source, dest)
		if err != nil {
			c.Stderr = err.Error()
			c.status = 1
			return 1
		}
		return 0
	} else {
		err := copyrecur.CopyFile(source, dest)
		if err != nil {
			c.Stderr = err.Error()
			c.status = 1
			return 1
		}
		return 0
	}
	return 0
}

// rename command
var renameCommand = &Command{
	Name: "Rename",
	run:  rename_file,
}

func rename_file(c *Command) int {

	fo := c.GetPathParam("source")
	fn := c.GetPathParam("dest")
	fo = strings.Trim(fo, "../")
	fn = strings.Trim(fn, "../")

	err := os.Rename(c.GetPath()+fo, c.GetPath()+fn)
	if err != nil {
		c.Stderr = err.Error()
		c.status = 1
		return 1
	}

	return 0
}

// Create a dir
var createfolderCommand = &Command{
	Name: "Create Folder",
	run:  create_folder,
}

func create_folder(c *Command) int {

	folder := c.Params["source"]
	file := c.GetPath() + strings.Trim(folder, "/")
	err := os.Mkdir(file, 0777)
	if err != nil {
		c.Stderr = err.Error()
		c.status = 1
		return 1
	}
	return 0
}

// Save a File
var saveCommand = &Command{
	Name: "Save File",
	run:  save_file,
}

func save_file(c *Command) int {
	data := []byte(c.Params["content"])
	file := c.GetPath() + c.GetPathParam("file")
	//@todo parametrize mask
	err := ioutil.WriteFile(file, data, 0644)
	if err != nil {
		c.status = 1
		c.Stderr = err.Error()
		return 1
	}
	return 0
}

// Delete a file, or list of files
var deleteCommand = &Command{
	Name: "Delete File",
	run:  delete_file,
}

func delete_file(c *Command) int {

	files := c.ParamsList
	errs := make([]string, 0)

	has_errors := false
	for k := range files {
		f := c.GetPath() + strings.Trim(files[k], "/")
		err := os.Remove(f)
		if err != nil {
			errs = append(errs, err.Error())
			has_errors = true
		}
	}

	if has_errors == true {
		c.status = 1
		c.Stderr = strings.Join(errs, ",")
		return 1
	}
	return 0

}

var compressCommand = &Command{
	Name: "Compress Tar/gz",
	run:  compress_file,
}

func compress_file(c *Command) int {
	// source will be .. /abc/abc/abc/
	// source will contain .. Absolute path to dir... /xxx/a/b/c
	source := c.GetPath() + c.GetPathParam("source")
	// base will contain c
	base := filepath.Base(source)
	dir := filepath.Dir(source)
	fname := fmt.Sprintf("%s.tar.gz", base)

	cmd := exec.Command("tar", "cvfz", fname, base)
	cmd.Dir = dir

	out, err := cmd.CombinedOutput()
	if err != nil {
		c.Stderr = string(out)
		c.status = 1
		return 1
	}

	c.Stdout = string(out)
	return 0
}

var mvCommand = &Command{
	Name: "Mv Command",
	run:  mv_file,
}

func mv_file(c *Command) int {

	source := c.GetPath() + c.GetPathParam("source")
	dest := c.GetPath() + c.GetPathParam("dest")
	cmd := exec.Command("mv", source, dest)

	out, err := cmd.CombinedOutput()
	if err != nil {
		c.Stderr = string(out)
		c.status = 1
		return 1
	}

	c.Stdout = string(out)
	return 0
}

/*

POST...

    {
        'command': 'save'
        'params': {
            'file': 'xxx'
            'content': 'xxxx'
        }
        'paramList': ['xxx', 'xxx']
        'path'
    }


    command: command_name
    params: hash

    // AjaxApiHandler
    type WebCommand struct {
        command string
        params map[string]string
        parmList []string
        pat string
    }

    type OutputCommand {
        Out string
        Err string
    }

    wc: = &WebCommand{}
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&wc)

    command := GetCommand(wc)
    if command == nil {
        http.Error(w, "Command Not Found", http.StatusInternalServerError)
        return
    }

    err = command.run()

    if err != nil {
        log.Error(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Sprint(w, Outputcommand)


*/
