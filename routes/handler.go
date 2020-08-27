package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"

	uuid "github.com/satori/go.uuid"
)

type m struct {
	message string
}

var mp = make(map[uuid.UUID]User)

// For json package to access the properties it needs to be capital which then can be exported and used by the json package

// User is ...
type User struct {
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Percentage float64 `json:"percentage"`
	Time       string  `json:"time"`
}

func Home(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	rw.Header().Set("Content-Type", "application/json")

	json.NewEncoder(rw).Encode(m{
		message: "Welcome to the home page route",
	})
	// fmt.Fprintln(rw, m)
}

func Get(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	rw.Header().Set("Content-Type", "application/json")

	json.NewEncoder(rw).Encode(m{
		message: "You have reached to get route",
	})
	// fmt.Fprintln(rw, "")
}

func Post(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	fmt.Println(u)
	if err != nil {
		log.Fatal(err)
		// Bad request
		http.Error(rw, err.Error(), 400)
		return
	}
	rw.Header().Set("Content-Type", "application/json")

	json.NewEncoder(rw).Encode(u)
	id := uuid.Must(uuid.NewV4())
	fmt.Fprintln(rw, "Sccussfull created the user", id)

	mp[id] = u
	// fmt.Fprintln(rw, "")
}

func Put(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		return
	}

	id, err := uuid.FromString(path.Base(r.URL.Path))
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	var u User
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		// log.Fatal(err)
		// Bad request
		http.Error(rw, err.Error(), 400)
		return
	}

	if _, ok := mp[id]; ok {
		mp[id] = u
	} else {
		fmt.Fprintln(rw, "Not present in map")
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(u)
	fmt.Fprintln(rw, "Sccussfully updated the user", id)
}

func Delete(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		return
	}

	id, err := uuid.FromString(path.Base(r.URL.Path))
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	if _, ok := mp[id]; ok {
		delete(mp, id)
		fmt.Fprintln(rw, "Successfully deleted!")
	} else {
		fmt.Fprintln(rw, "Not present in map")
	}
}
