package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github/com/L4ZZA/matchmaking/data"
)

// Sessions is a http.Handler
type Sessions struct {
	l *log.Logger
}

// NewSessions creates a Sessions handler with the given logger
func NewSessions(l *log.Logger) *Sessions {
	return &Sessions{l}
}

func (s *Sessions) Greetings(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("[DEBUG] greet new player")
	rw.Write([]byte("Thanks for joining THE GAME. Use the handle /join to enter the first available lobby."))
}

// GetSessions returns the Sessions from the data store
func (p *Sessions) GetSessions(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Sessions")

	// fetch the Sessions from the datastore
	lp := data.GetSessions()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		p.l.Println("GetSessions - Failed to parse to JSON", lp, err)
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	p.l.Println("GetSessions - COMPLETED")
}

func (p *Sessions) AddSession(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Session")

	prod := r.Context().Value(KeySession{}).(data.Session)
	data.AddSession(&prod)
	p.l.Println("AddSession - COMPLETED")
}

type KeySession struct{}

func (p Sessions) MiddlewareValidateSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Session{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing Session", err)
			http.Error(rw, "Error reading Session", http.StatusBadRequest)
			return
		}

		// validate the Session
		err = prod.Validate()
		if err != nil {
			p.l.Println("[MiddlewareValidateSession] validating Session", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating Session: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// add the Session to the context
		ctx := context.WithValue(r.Context(), KeySession{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
