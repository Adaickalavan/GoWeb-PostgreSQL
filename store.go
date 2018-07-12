package main

// The sql go library is needed to interact with the database
import (
	"database/sql"
)

// Store will have two methods, to add a new person, and to get all existing people
type Store interface {
	CreatePerson(person *Person) error
	GetPerson() ([]*Person, error)
}

// `dbStore` struct implements the `Store` interface
// Variable db takes the pointer to the SQL database connection object
type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreatePerson(person *Person) error {
	// 'Person' is a struct which has "nama", "birthday", and "occupation" attributes
	// Note: `peopleInfo` is the name of the table within our `peopleDatabase`
	_, err := store.db.Query(
		"INSERT INTO peopleInfo(nama,birthday,occupation) VALUES ($1,$2,$3)",
		person.Nama, person.Birthday, person.Occupation)
	return err
}

func (store *dbStore) GetPeople() ([]*Person, error) {
	// Query the database for all persons, and return the result to the `rows` object
	// Note: `peopleInfo` is the name of the table within our `peopleDatabase``
	rows, err := store.db.Query("SELECT nama, birthday, occupation from peopleInfo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create the data structure that is returned from the function.
	// By default, this will be an empty array of persons
	personList := []*Person{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a person,
		person := &Person{}
		// Populate the `Name`, `Birthday`, and `Occupation` attributes of the person
		if err := rows.Scan(&person.Nama, &person.Birthday, &person.Occupation); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for the next row
		personList = append(personList, person)
	}
	return personList, nil
}

// Define a `store` variable which is a package level variable
var store Store

/*
We will need to call the InitStore method to initialize the store. This will
typically be done at the beginning of our application (in this case, when the server starts up)
This can also be used to set up the store as a mock, which we will be observing
later on
*/
func InitStore(s Store) {
	store = s
}
