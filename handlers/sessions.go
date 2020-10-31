package handlers

import (
	"log"
	"net/http"
	"strconv"
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

type KeySession struct{}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// GenericSuccessMessage is a generic error message returned by a server
type GenericSuccessMessage struct {
	Message string `json:"message"`
}

// getSessionID returns the Session ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getSessionID(r *http.Request) int {
	// parse the Session id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["session_id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}