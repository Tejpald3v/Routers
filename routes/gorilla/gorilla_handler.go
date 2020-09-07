package gorilla

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

var mp = make(map[uuid.UUID]User)

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
		sendResponse(rw, "Bad Request", 400, err.Error(), nil)
		// TODO: DO not stop the server.
		log.Fatal()
	}
	return u
}

func writeResponse(rw http.ResponseWriter, s User) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(s)
}

// Home ...
func Home(rw http.ResponseWriter, r *http.Request) {
	sendResponse(rw, "Success", 200, "Welcome to the home route", nil)
	// fmt.Fprintln(rw, "")
}

// Get ...
func Get(rw http.ResponseWriter, r *http.Request) {
	sendResponse(rw, "Success", 200, "You have reached to get route", nil)
}

// Post ...
func Post(rw http.ResponseWriter, r *http.Request) {
	u := decodeJSON(rw, r)
	id := uuid.NewV4()

	// Storing the user in the map with id as the key
	mp[id] = u
	// TODO:	Merge id and the user struct. Create struct that include id and user data or convert the id to string type and send speratly indicating the id.
	sendResponse(rw, "Success", 201, "Sccussfully created the user", id)
}

// Put ...
func Put(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	u := decodeJSON(rw, r)
	id, _ := uuid.FromString(vars["id"])
	if _, ok := mp[id]; ok {
		mp[id] = u
		sendResponse(rw, "Success", 200, "Updated user successfully!", u)
		return
	}
	sendResponse(rw, "Bad Request", 400, "ID not present in map!", nil)
}

// Delete ...
func Delete(rw http.ResponseWriter, r *http.Request) {
	id, _ := uuid.FromString(mux.Vars(r)["id"])
	if _, ok := mp[id]; !ok {
		sendResponse(rw, "Bad Request", 400, "ID not present in map!", nil)
		return
	}
	delete(mp, id)
	sendResponse(rw, "Success", 200, "Successfully deleted!", nil)
}
