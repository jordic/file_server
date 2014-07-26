package cmdwebstream

import (
	"bufio"
	"io"
	"net/http"
	"os/exec"
)

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

func Handler(w http.ResponseWriter, r *http.Request, cmd *exec.Cmd) {

	w.Header().Set("Content-Type", "application/octet-stream")

	fw := flushWriter{w: w}
	if f, ok := w.(http.Flusher); ok {
		fw.f = f
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bfin := bufio.NewReader(stdout)
	bfer := bufio.NewReader(stderr)

	cmd.Start()
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

	if err := cmd.Wait(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
