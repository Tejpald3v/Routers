package gorilla

import (
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

func checkID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		_, err := uuid.FromString(mux.Vars(r)["id"])
		if err != nil {
			sendResponse(rw, "Bad Request", 400, err.Error(), nil)
			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
