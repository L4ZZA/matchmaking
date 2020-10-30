package handlers

import "log"

// Sessions is a http.Handler
type Sessions struct {
	l *log.Logger
}

// NewSessions creates a Sessions handler with the given logger
func NewSessions(l *log.Logger) *Sessions {
	return &Sessions{l}
}

type KeySession struct{}

