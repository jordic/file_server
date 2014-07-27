package cmdwebstream

import (
	"bufio"
	"io"
	"net/http"
	"os/exec"
)

// Cmd starts a system command, redirecting, stdout and stdin
// the the browser request..
// Its possible to implement a tail -f... but i don't know when a browser
// will kill a request..
type Cmd struct {
	Command *exec.Cmd
}

func (self *Cmd) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/octet-stream")

	fw := flushWriter{w: w}
	if f, ok := w.(http.Flusher); ok {
		fw.f = f
	}

	stdout, err := self.Command.StdoutPipe()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	stderr, err := self.Command.StderrPipe()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bfin := bufio.NewReader(stdout)
	bfer := bufio.NewReader(stderr)

	self.Command.Start()
	gr := func(buf io.Reader) {
		for {
			b := make([]byte, 8)
			_, err := buf.Read(b)
			if err != nil {
				break
			}
			fw.Write(b)
		}
	}

	go gr(bfin)
	go gr(bfer)

	if err := self.Command.Wait(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

// Implementing flsuh writer, because we want a stream output...
// flushed every write.
type flushWriter struct {
	f http.Flusher
	w io.Writer
}

func (fw *flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if fw.f != nil {
		fw.f.Flush()
	}
	return
}
