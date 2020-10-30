package handlers

import (
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// Players is a http.Handler
type Players struct {
	l *log.Logger
}

// NewSessions creates a Players handler with the given logger
func NewPlayers(l *log.Logger) *Players {
	return &Players{l}
}

type KeyPlayer struct{}


// getPlayerID returns the Player ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getPlayerID(r *http.Request) int {
	// parse the Player id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["player_id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}