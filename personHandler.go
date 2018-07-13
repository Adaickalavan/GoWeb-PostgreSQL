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

func getPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve people from postgresql database using our `store` interface variable defined in `store.go` file
	personList, err := store.GetPerson()

	// Convert the `personList` variable to json
	personListBytes, err := json.Marshal(personList)

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

	// Write new person details into postgresql database using our `store` interface variable's `CreatePerson` method
	err = store.CreatePerson(&person)
	if err != nil {
		fmt.Println(err)
	}

	//Redirect to the original HTML page
	http.Redirect(w, r, "/", http.StatusFound)
}
