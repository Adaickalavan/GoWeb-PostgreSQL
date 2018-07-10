package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create router
	router := newRouter()
	// Listen to the port
	log.Fatal(http.ListenAndServe(":8080", router))
}

func newRouter() *mux.Router {
	// Declare a router
	router := mux.NewRouter()
	// Handler for specified path
	router.HandleFunc("/hello", handler).Methods("POST")
	// Declare static file directory
	// staticFileDirectory := http.Dir("/assests/")
	//
	// staticFileHandler := http.StripPrefix("/assets/", http.FileServer())

	data := todo{
		page: "hi",
	}

	return router
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi, the world is beautiful")
}


