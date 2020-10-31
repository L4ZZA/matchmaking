package handlers

import (
	"fmt"
	"context"
	"net/http"

	"matchmaking.test/data"
)

func (s Sessions) MiddlewareValidatePlayer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		player := data.Player{}

		err := player.FromJSON(r.Body)
		if err != nil {
			s.l.Println("[ERROR] deserializing Player", err)
			http.Error(rw, "Error reading Player", http.StatusBadRequest)
			return
		}

		// validate the Player
		err = player.Validate()
		if err != nil {
			s.l.Println("[ERROR] validating Player", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating Player: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// add the Player to the context
		ctx := context.WithValue(r.Context(), KeyPlayer{}, player)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}

func (s Sessions) MiddlewareValidateSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		sess := data.Session{}

		err := sess.FromJSON(r.Body)
		if err != nil {
			s.l.Println("[ERROR] deserializing Session", err)
			http.Error(rw, "Error reading Session", http.StatusBadRequest)
			return
		}

		// validate the Session
		err = sess.Validate()
		if err != nil {
			s.l.Println("[ERROR] validating Session", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating Session: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// add the Session to the context
		ctx := context.WithValue(r.Context(), KeySession{}, sess)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}