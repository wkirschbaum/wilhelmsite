package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"

	"bitbucket.org/wkirschbaum/wilhelmsite/app"

	"github.com/gorilla/mux"
)

func getCssHandler() func(http.ResponseWriter, *http.Request) {
	filename := "public/css/kernel.css"
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-cache")
		fileBytes, err := ioutil.ReadFile(filename)
		if err == nil {
			etag := fmt.Sprintf("%x", md5.Sum(fileBytes))
			w.Header().Add("ETag", etag)
		}
		http.ServeFile(w, r, (filename))
	}
	handler = app.Gzip(handler)
	return handler
}

func getPageHandler(filename string) func(http.ResponseWriter, *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-store")
		http.ServeFile(w, r, ("public/" + filename))
	}
	return app.Gzip(handler)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getPageHandler("index.html"))
	router.HandleFunc("/public/kernel.css", getCssHandler())
	http.ListenAndServe(":8000", router)
}
