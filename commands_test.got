package main

import (
	//"fmt"
	"os"
	//("os/exec"
	"io/ioutil"
	"path/filepath"
	"testing"
)

// +build ignore

var dataDir string = "./test_data/"
var p string

func init() {

	var err error
	p, err = filepath.Abs(dataDir)
	if err != nil {
		panic(err)
	}
	// clean data dir for start testing...
	clean_up()
}

func clean_up() {
	err := os.RemoveAll(p)
	if err != nil {
		panic(err)
	}

	err = os.Mkdir(p, 0777)
	if err != nil {
		panic(err)
	}

}

func TestGetPath(t *testing.T) {

	cm := GetCommand("save", p)
	//fmt.Print(cm.Path)
	if cm.Path != p {
		t.Errorf("%s!=%s", cm.Path, p)
	}

	if cm.GetPath() != p+"/" {
		t.Errorf("%s!=%s", cm.GetPath(), p+"/")
	}

}

func TestFile(t *testing.T) {

	clean_up()
	cm := GetCommand("save", p)
	cm.Params = map[string]string{
		"file":    "test.txt",
		"content": "test",
	}
	res := cm.Run()
	if res == 1 {
		t.Errorf("Terror", cm.Stderr)
	}

	d, err := ioutil.ReadFile(p + "/test.txt")
	if err != nil {
		t.Error(err)
	}
	if string(d) != "test" {
		t.Error("Wrong content")
	}

	cm = GetCommand("rename", p)
	cm.Params = map[string]string{
		"source": "test.txt",
		"dest":   "test2.txt",
	}

	res = cm.Run()
	if res == 1 {
		t.Errorf("Terror", cm.Stderr)
	}

	cm = GetCommand("delete", p)
	cm.ParamsList = []string{"test2.txt"}
	cm.Path = p
	res = cm.Run()
	if res == 1 {
		t.Error(cm.Stderr)
	}

}

func TestCopyFile(t *testing.T) {
	clean_up()
	cm := GetCommand("save", p)
	cm.Params = map[string]string{
		"file":    "test.txt",
		"content": "test",
	}
	res := cm.Run()
	if res == 1 {
		t.Errorf("Terror", cm.Stderr)
	}

	cm = GetCommand("copy", p)
	cm.Params = map[string]string{
		"source": "test.txt",
		"dest":   "test2.txt",
	}
	res = cm.Run()
	if res == 1 {
		t.Error(cm.Stderr)
	}

	d, err := ioutil.ReadFile(p + "/test.txt")
	if err != nil {
		t.Error(err)
	}
	r, err := ioutil.ReadFile(p + "/test.txt")
	if err != nil {
		t.Error(err)
	}

	if string(d) != string(r) {
		t.Errorf("%s!=%s", d, r)
	}

	//@todo copy dir

}

func TestCreateFolder(t *testing.T) {

	clean_up()

	cm := GetCommand("createFolder", p)
	cm.Params = map[string]string{
		"source": "/test",
	}
	res := cm.Run()
	if res == 1 {
		t.Errorf("TError %s", cm.Stderr)
	}

	if _, err := os.Stat(p + "/test"); err != nil {
		t.Error("File not found")
	}

	cm = GetCommand("createFolder", p)
	cm.Params = map[string]string{
		"source": "/test",
	}

	res = cm.Run()
	if res == 0 {
		t.Errorf("TError %s", cm.Stderr)
	}

	clean_up()

	cm = GetCommand("createFolder", p)
	cm.Params = map[string]string{
		"source": "test",
	}

	res = cm.Run()
	if res == 1 {
		t.Errorf("TError %s %#v", cm.Stderr, cm)
	}

}

func TestCompressMv(t *testing.T) {

	clean_up()
	cm := GetCommand("save", p)
	cm.Params = map[string]string{
		"file":    "test.txt",
		"content": "test",
	}
	res := cm.Run()
	if res == 1 {
		t.Errorf("Terror", cm.Stderr)
	}

	cm = GetCommand("compress", p)
	cm.Params = map[string]string{
		"source": "test.txt",
	}
	cm.Run()
	if _, err := os.Stat(p + "/test.txt.tar.gz"); err != nil {
		t.Error("File not found")
	}

	cm = GetCommand("mv", p)
	cm.Params = map[string]string{
		"source": "test.txt.tar.gz",
		"dest":   "test2.tar.gz",
	}

	cm.Run()
	if _, err := os.Stat(p + "/test2.tar.gz"); err != nil {
		t.Error("File not found")
	}

}
