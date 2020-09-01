package main

import "Routers/routes/gin/route"

func main() {
	// Gin router
	route.HandleRequest()

	// Gorilla mux router
	// r := mux.NewRouter()
	// gorilla.HandleRequest(r)

	// For using default http router uncomment the 20th line and comment 25th line
	// Default http router
	// defaulthttp.HandleRequest()

	// create a new server
	// s := http.Server{
	// 	Addr:         "127.0.0.1:8080",  // configure the bind address
	// 	Handler:      r,                 // set the default handler
	// 	ReadTimeout:  5 * time.Second,   // max time to read request from the client
	// 	WriteTimeout: 10 * time.Second,  // max time to write response to the client
	// 	IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	// }

	// start the server
	// s.ListenAndServe()
}
