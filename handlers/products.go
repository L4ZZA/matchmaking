package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github/com/L4ZZA/matchmaking/data"

	"github.com/gorilla/mux"
)

// Sessions is a http.Handler
type Sessions struct {
	l *log.Logger
}

// NewSessions creates a Sessions handler with the given logger
func NewSessions(l *log.Logger) *Sessions {
	return &Sessions{l}
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

func (p Sessions) UpdateSessions(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		p.l.Println("UpdateSessions - can't convert id", id, err)
		http.Error(rw, "Unable to cast id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Session", id)
	prod := r.Context().Value(KeySession{}).(data.Session)

	err = data.UpdateSession(id, &prod)
	if err == data.ErrSessionNotFound {
		p.l.Println("UpdateSessions - ERROR2 ", err)
		http.Error(rw, "Session not found", http.StatusNotFound)
		return
	}

	if err != nil {
		p.l.Println("UpdateSessions - ERROR3")
		http.Error(rw, "Session not found", http.StatusInternalServerError)
		return
	}
	p.l.Println("UpdateSessions - COMPLETED")
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
