package main

import (
	"net/http"

	"github.com/unrolled/render"
)

var sampleString = string("This is a my sample!")

func main() {
	r := render.New()
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcome, visit page now."))
	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
		r.Data(w, http.StatusOK, []byte("Some binary data here."))
	})

	mux.HandleFunc("/text", func(w http.ResponseWriter, req *http.Request) {
		r.Text(w, http.StatusOK, "Plain text here")

	})

	mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
		// Assumes you have a template in ./templates called "example.tmpl".
		// $ mkdir -p templates && echo "<h1>Hello HTML world.</h1>" > templates/example.tmpl

		r.HTML(w, http.StatusOK, "example", sampleString)

	})
	http.ListenAndServe(":3000", mux)

}
