package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Person is a struct decsribing its properties
type Person struct {
	Nama       string `json:"nama"`
	Birthday   string `json:"birthday"`
	Occupation string `json:"occupation"`
}

var persons []Person

func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Convert the "persons" variable to json
	personListBytes, err := json.Marshal(persons)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Write JSON list of persons to response
	w.Write(personListBytes)
}

func createPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML form data
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the person from the form info
	person := Person{}
	person.Nama = r.Form.Get("nama")
	person.Birthday = r.Form.Get("birthday")
	person.Occupation = r.Form.Get("occupation")

	// Append our existing list of persons with a new entry
	persons = append(persons, person)

	//Redirect to the original HTML page
	http.Redirect(w, r, "/", http.StatusFound)
}
