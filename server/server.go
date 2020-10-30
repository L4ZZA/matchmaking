package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github/com/L4ZZA/matchmaking/handlers"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "server ", log.LstdFlags)

	// create the handlers
	ph := handlers.NewSessions(l)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", ph.Greetings)
	getRouter.HandleFunc("/sessions", ph.GetSessions)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/join", ph.AddSession)
	postRouter.Use(ph.MiddlewareValidateSession)

	// sm.Handle("/", ph)

	// create a new server
	s := http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	l.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
