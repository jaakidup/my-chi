package main

import "fmt"

// DB ...
// Interface for what the database should look like
// a single record type hasn't been decided on, thus we use an interface{}
type DB interface {
	create(collection string, entry interface{}) (ID string, err error)
	read(collection string, ID string) (entry interface{}, err error)
	update(collection string, ID string) (success bool, err error)
	delete(collection string, ID string) (success bool, err error)
}

// ObjectStore ...
// This ObjectStore stores JSON
type ObjectStore struct {
	Name string
}

func (*ObjectStore) create(collection string, entry interface{}) (ID string, err error) {
	fmt.Println("Storing entry to collection: ", collection)
	fmt.Println(entry)
	return "1", nil
}

func (*ObjectStore) read(collection string, ID string) (entry interface{}, err error) {

	document := struct {
		ID    string `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Value string `json:"value,omitempty"`
	}{
		ID:    "some-id=1234",
		Name:  "Jaaki",
		Value: "Is amped!!",
	}
	return document, nil
}
func (*ObjectStore) update(collection string, ID string) (success bool, err error) {
	return true, nil
}
func (*ObjectStore) delete(collection string, ID string) (success bool, err error) {
	return true, nil
}
