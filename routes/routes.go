package routes

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// HandleRequest ...
func HandleRequest() {

	r := mux.NewRouter()

	// create a new server
	s := http.Server{
		Addr:         "127.0.0.1:8080",  // configure the bind address
		Handler:      r,                 // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		log.Println("Starting server on port 8080")

		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

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

	http.ListenAndServe(":8080", r)
}
