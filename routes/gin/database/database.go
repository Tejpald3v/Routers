package database

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type Getter interface {
	GetAll() map[uuid.UUID]User
}

type AddOrUpdate interface {
	Check(id uuid.UUID) error
	Add(id uuid.UUID, user User)
}

type Delete interface {
	Check(id uuid.UUID) error
	Delete(id uuid.UUID)
}

// User is ...
type User struct {
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Percentage float64 `json:"percentage"`
	Time       string  `json:"time"`
}

type Repo struct {
	Users map[uuid.UUID]User
}

func New() *Repo {
	return &Repo{
		Users: make(map[uuid.UUID]User),
	}
}

func (r *Repo) Add(id uuid.UUID, user User) {
	r.Users[id] = user
}

func (r *Repo) GetAll() map[uuid.UUID]User {
	return r.Users
}

func (r *Repo) Check(id uuid.UUID) error {
	if _, ok := r.Users[id]; !ok {
		return errors.New("ID not present in map")
	}
	return nil
}

func (r *Repo) Delete(id uuid.UUID) {
	delete(r.Users, id)
}
