package main

import (
	"encoding/json"
	"github.com/jordic/file_server/cmdwebstream"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type WebCommand struct {
	Command    string            `json:"action"`
	Params     map[string]string `json:"params"`
	ParamsList []string          `json:"paramslist"`
}

type WebCommandResp struct {
	Status int    `json:"status"`
	Msg    string `json:"message"`
}

func WebCommandHandler(w http.ResponseWriter, r *http.Request) {

	wc := &WebCommand{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&wc)
	if err != nil {
		http.Error(w, "Error decoding json", http.StatusInternalServerError)
		log.Printf("e decoding %s", err)
		return
	}

	if wc.Command == "syscmd" && disable_sys_command == true {
		http.NotFound(w, r)
		log.Printf("commands explicity disabled by commandline %s", wc.Command)
		return
	}

	if wc.Command == "syscmd" {
		//HandlerStreamCommand(w, wc)
		path := strings.TrimRight(dir, "/") + "/"
		source := strings.Trim(wc.Params["source"], "/")
		args := wc.ParamsList

		cmd := exec.Command(wc.Params["command"], args...)
		cmd.Dir = path + source

		cw := &cmdwebstream.Cmd{
			Command: cmd,
		}

		cw.ServeHTTP(w, r)
		//cmdwebstream.Handler(w, r, cmd)
		return

	}

	cmd := GetCommand(wc.Command, dir)
	if cmd == nil {
		//http.Error(w, "Command Not found", http.StatusInternalServerError)
		http.NotFound(w, r)
		log.Printf("command not found %s", wc.Command)
		return
	}

	resp := &WebCommandResp{}

	cmd.Params = wc.Params
	cmd.ParamsList = wc.ParamsList
	res := cmd.Run()
	if res != 0 {
		resp.Status = res
		resp.Msg = cmd.Stderr
	} else {
		resp.Status = 0
		resp.Msg = cmd.Stdout
	}

	out, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
		log.Printf("e encoding json %s", err)
		return
	}

	log.Printf("command executed %s", wc.Command)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

	return
}
