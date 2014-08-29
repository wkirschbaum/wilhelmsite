package main

import "net/http"

func homePageHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "public/index.html")
}

func main() {
	http.HandleFunc("/", homePageHandler)
	http.ListenAndServe(":8000", nil)
}
