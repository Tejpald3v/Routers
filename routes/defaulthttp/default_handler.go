package defaulthttp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"

	uuid "github.com/satori/go.uuid"
)

var mp = make(map[uuid.UUID]User)

// Response data format for HTTP
type Response struct {
	Status  string      `json:"status" bson:"status"`                       // Status code (error|fail|success)
	Code    int         `json:"code"  bson:"code"`                          // HTTP status code
	Message string      `json:"message,omitempty" bson:"message,omitempty"` // Error or status message
	Data    interface{} `json:"data,omitempty" bson:"data,omitempty"`       // Data payload
}

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
	}
	fmt.Fprintln(rw, "Not present in map")
	return false
}

type res struct {
	message string `json:"message"`
}

func incorrect(rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(res{"Incorrect http verb used"})
}

// Home ...
func Home(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		incorrect(rw)
		return
	}
	fmt.Fprintln(rw, "Welcome to the home page route")
}

// Get ...
func Get(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		incorrect(rw)
		return
	}
	fmt.Fprintln(rw, "You have reached to get route")
}

// Post ...
func Post(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		incorrect(rw)
		return
	}
	u := decodeJSON(rw, r)
	id := uuid.NewV4()

	// Storing the user in the map with id as the key
	mp[id] = u
	writeResponse(rw, u)
	fmt.Fprintln(rw, "Sccussfully created the user", id)
}

// Put ...
func Put(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		incorrect(rw)
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

// Delete ...
func Delete(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		incorrect(rw)
		return
	}

	if checkID(rw, r) {
		id, _ := uuid.FromString(path.Base(r.URL.Path))
		delete(mp, id)
		fmt.Fprintln(rw, "Successfully deleted!")
	}
}
