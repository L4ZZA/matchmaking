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

	p.l.Println("[DEBUG:DELETE] removing player id:", sessId)

	err := data.RemovePlayer(sessId, playerId)
	if err == data.ErrPlayerNotFound {
		message := fmt.Sprintf("Could not delete player id: %d from session id: %d", playerId, sessId)
		m := fmt.Sprintf("[ERROR:DELETE] %s", message)
		p.l.Println(m)

		m = fmt.Sprintf("%s - %s", message, err.Error())
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: m}, rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR:DELETE] deleting player -", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
	data.ToJSON(&GenericSuccessMessage{Message: err.Error()}, rw)
	p.l.Println("[DEBUG:DELETE] removing player id:", sessId)
}