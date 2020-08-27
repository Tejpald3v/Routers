package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"

	uuid "github.com/satori/go.uuid"
)

var mp = make(map[uuid.UUID]User)

// For json package to access the properties it needs to be capital which then can be exported and used by the json package

// User is ...
type User struct {
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Percentage float64 `json:"percentage"`
	Time       string  `json:"time"`
}

func decodeJSON(rw http.ResponseWriter, r *http.Request) User {
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		log.Fatal(err)
	}
	return u
}

func writeResponse(rw http.ResponseWriter, s User) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(s)
}

func checkID(rw http.ResponseWriter, r *http.Request) bool {
	id, err := uuid.FromString(path.Base(r.URL.Path))
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return false
	}

	if _, ok := mp[id]; ok {
		return true
	} else {
		fmt.Fprintln(rw, "Not present in map")
	}
	return false
}

func Home(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	fmt.Fprintln(rw, "Welcome to the home page route")
}

func Get(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	fmt.Fprintln(rw, "You have reached to get route")
}

func Post(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	u := decodeJSON(rw, r)
	id := uuid.Must(uuid.NewV4())

	// Storing the user in the map with id as the key
	mp[id] = u
	writeResponse(rw, u)
	fmt.Fprintln(rw, "Sccussfully created the user", id)
}

func Put(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		return
	}

	if checkID(rw, r) {
		u := decodeJSON(rw, r)
		id, _ := uuid.FromString(path.Base(r.URL.Path))
		mp[id] = u
		fmt.Fprintln(rw, "Updated User!")
		writeResponse(rw, u)
	}
}

func Delete(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		return
	}

	if checkID(rw, r) {
		id, _ := uuid.FromString(path.Base(r.URL.Path))
		delete(mp, id)
		fmt.Fprintln(rw, "Successfully deleted!")
	}
}
