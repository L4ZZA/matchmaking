package handlers

import "log"

// Players is a http.Handler
type Players struct {
	l *log.Logger
}

// NewSessions creates a Players handler with the given logger
func NewPlayers(l *log.Logger) *Players {
	return &Players{l}
}

type KeyPlayer struct{}

