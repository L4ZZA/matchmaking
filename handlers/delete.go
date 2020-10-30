package handlers

import (
	"fmt"
	"net/http"

	"github/com/L4ZZA/matchmaking/data"
)

// RemovePlayer handles DELETE requests and removes items from the database
func (p *Sessions) RemovePlayer(rw http.ResponseWriter, r *http.Request) {
	sessId := getSessionID(r)
	playerId := getPlayerID(r)

	p.l.Println("[DEBUG] removing player id:", sessId)

	err := data.RemovePlayer(sessId, playerId)
	if err == data.ErrPlayerNotFound {
		p.l.Println(fmt.Sprintf("[ERROR] deleting player id: %d does not exist", playerId))

		rw.WriteHeader(http.StatusNotFound)
		// data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting player -", err)

		rw.WriteHeader(http.StatusInternalServerError)
		// data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}