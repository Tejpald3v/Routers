package main

import (
	"Routers/routes"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server")
	http.HandleFunc("/", routes.Home)
	http.HandleFunc("/get", routes.Get)
	http.HandleFunc("/post", routes.Post)
	http.HandleFunc("/put/", routes.Put)
	http.HandleFunc("/delete/", routes.Delete)

	http.ListenAndServe(":8080", nil)
}
