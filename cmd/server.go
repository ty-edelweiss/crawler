package main

import (
	"net/http"

	"github.com/ty-edelweiss/crawler"
)

func main() {
	s := server.NewServer("Server start")
	http.HandleFunc("/test", s.Handler)
	http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello Go!"))
	})

	http.ListenAndServe(":80", nil)
}
