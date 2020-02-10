package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("src/static/")))
	http.Handle("/", r)
	http.ListenAndServe(":8100", nil)
}