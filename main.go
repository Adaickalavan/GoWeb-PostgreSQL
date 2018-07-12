package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	// Declare a router
	r := mux.NewRouter()
	// Declare static file directory
	staticFileDirectory := http.Dir("./static/")
	// Create static file server
	staticFileServer := http.FileServer(staticFileDirectory)
	// Create file handler
	staticFileHandler := http.StripPrefix("/", staticFileServer)
	// Add handler to router
	r.Handle("/", staticFileHandler).Methods("GET")
	// Add handler for get and post people
	r.HandleFunc("/person", getPersonHandler).Methods("GET")
	r.HandleFunc("/person", createPersonHandler).Methods("POST")

	return r
}

func main() {
	// Create router
	r := newRouter()
	// Listen to the port
	http.ListenAndServe(":8080", r)
}
