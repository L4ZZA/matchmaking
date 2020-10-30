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
	sh := handlers.NewSessions(l)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", sh.Greetings)
	getRouter.HandleFunc("/sessions", sh.GetSessions)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/join", sh.AddPlayer)
	postRouter.Use(sh.MiddlewareValidatePlayer)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/leave/{session_id:[0-9]+}/{player_id:[0-9]+}", sh.RemovePlayer)

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
