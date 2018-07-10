package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// Declare a router
	router := mux.NewRouter()

	// 
	router.HandleFunc("/hello", handler)
	// Listen to the port
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"hi the world is beautiful")
}
