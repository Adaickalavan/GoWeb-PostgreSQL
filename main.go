package main

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	// Import the `pq` package with a preceding underscore since it is imported as a side effect.
	// `pq` package is a GO Postgres driver for the `database/sql` package.
	_ "github.com/lib/pq"
)

func newRouter() *mux.Router {
	// Declare a router
	r := mux.NewRouter()
	// Declare static file directory
	staticFileDirectory := http.Dir("./static/")
	// Create static file server for our static files, i.e., .html, .css, etc
	staticFileServer := http.FileServer(staticFileDirectory)
	// Create file handler.
	// Although the static files are placed inside './static/' folder in our local directory, it is served at the root (i.e., http://localhost:8080/) when browsed in a browser.
	// Hence, we need `http.StripPrefix` function to change the serve path.
	staticFileHandler := http.StripPrefix("/", staticFileServer)
	// Add handler to router
	r.Handle("/", staticFileHandler).Methods("GET")
	// Add handler for get and post people
	r.HandleFunc("/person", getPersonHandler).Methods("GET")
	r.HandleFunc("/person", createPersonHandler).Methods("POST")

	return r
}

func main() {
	// Setup connection to our postgresql database
	connString := `user=postgres 
				   password=1234
				   host=localhost
				   port=5432
				   dbname=peopleDatabase 
				   sslmode=disable`
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	// Check whether we can access the database by pinging it
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Iniatialize a `store` variable of type `Store` interface
	// Place our opened database into a `dbstruct` and implement a `Store` interface
	var store Store
	store = &dbStore{db: db}

	// Create router
	r := newRouter()

	// Listen to the port. Go server's default port is 8080
	http.ListenAndServe(":8080", r)
}
