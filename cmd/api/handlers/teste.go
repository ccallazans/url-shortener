package handlers

import "net/http"

func Teste(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(503)
	w.Write([]byte("bad"))
}
