package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

// func checkID(rw http.ResponseWriter, r *http.Request) bool {
// 	id, err := uuid.FromString(mux.Vars(r)["id"])
// 	if err != nil {
// 		http.Error(rw, err.Error(), 400)
// 		return false
// 	}

// 	if _, ok := mp[id]; ok {
// 		return true
// 	}
// 	fmt.Fprintln(rw, "Not present in map")
// 	return false
// }

// Home ...
func Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Welcome to the home route")
}

// Get ...
func Get(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "You have reached to get route")
}

// Post ...
func Post(rw http.ResponseWriter, r *http.Request) {
	u := decodeJSON(rw, r)
	id := uuid.Must(uuid.NewV4())

	// Storing the user in the map with id as the key
	mp[id] = u
	writeResponse(rw, u)
	fmt.Fprintln(rw, "Sccussfully created the user", id)
}

// Put ...
func Put(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	u := decodeJSON(rw, r)
	id, _ := uuid.FromString(vars["id"])
	if _, ok := mp[id]; ok {
		mp[id] = u
		fmt.Fprintln(rw, "Updated User!")
		writeResponse(rw, u)
		return
	}
	fmt.Fprintln(rw, "ID not present in map")
}

// Delete ...
func Delete(rw http.ResponseWriter, r *http.Request) {
	id, _ := uuid.FromString(mux.Vars(r)["id"])
	if _, ok := mp[id]; !ok {
		fmt.Fprintln(rw, "ID not present in map")
		return
	}
	delete(mp, id)
	fmt.Fprintln(rw, "Successfully deleted!")
}
