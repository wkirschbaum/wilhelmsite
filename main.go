package main

import (
	"net/http"

	"bitbucket.org/wkirschbaum/wilhelmsite/app"

	"github.com/gorilla/mux"
)

func getCssHandler() func(http.ResponseWriter, *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "max-age=600")
		http.ServeFile(w, r, ("public/site.css"))
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
	router.HandleFunc("/", getFileHandler("index.html"))
	router.HandleFunc("/public/site.css", getCssHandler())
	http.ListenAndServe(":8000", router)
}
