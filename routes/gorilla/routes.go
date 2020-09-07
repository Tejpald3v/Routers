package gorilla

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HandleRequest ...
func HandleRequest(r *mux.Router) {

	getRouter := r.Methods("Get").Subrouter()
	getRouter.HandleFunc("/", Home)
	getRouter.HandleFunc("/get", Get)

	putRouter := r.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/put/{id}", Put)
	putRouter.Use(checkID)

	postRouter := r.Methods("POST").Subrouter()
	postRouter.HandleFunc("/post", Post)

	deleteRouter := r.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/delete/{id}", Delete)
	putRouter.Use(checkID)

}
