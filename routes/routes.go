package routes

import (
	"net/http"
)

func HandleRequest() {
	// fmt.Println("Server")
	http.HandleFunc("/", Home)
	http.HandleFunc("/get", Get)
	http.HandleFunc("/post", Post)
	http.HandleFunc("/put/", Put)
	http.HandleFunc("/delete/", Delete)

	http.ListenAndServe(":8080", nil)
}
