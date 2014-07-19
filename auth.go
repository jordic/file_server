package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func BasicAuth(fn http.HandlerFunc, cred string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if cred == "" {
			fn(w, r)
			return
		}

		if checkAuth(w, r, cred) {
			fn(w, r)
			return
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="Authenticate"`)
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))

	}
}

func checkAuth(w http.ResponseWriter, r *http.Request, cred string) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}

	user := strings.Split(cred, ":")
	return pair[0] == user[0] && pair[1] == user[1]

}
