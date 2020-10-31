package handlers

import (
	"fmt"
	"net/http"

	"matchmaking.test/data"
)

// RemovePlayer handles DELETE requests and removes items from the database
func (p *Sessions) RemovePlayer(rw http.ResponseWriter, r *http.Request) {

	playerId := getPlayerID(r)
	player, err := data.GetPlayerFromId(playerId)

	if err != nil {
		message := fmt.Sprintf("Could not find player id: %d in any session", playerId)
		p.l.Println(fmt.Sprintf("[ERROR:DELETE] %s", message))

		m := fmt.Sprintf("%s", message)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: m}, rw)
		return
	}


	p.l.Println("[DEBUG:DELETE] removing player id:", playerId)

	err = data.RemovePlayer(player.SessionID, playerId)
	if err == data.ErrPlayerNotFound {
		message := fmt.Sprintf("Could not delete player id: %d from session id: %d", playerId, player.SessionID)
		p.l.Println(fmt.Sprintf("[ERROR:DELETE] %s", message))

		m := fmt.Sprintf("%s", message)
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
	data.ToJSON(&GenericSuccessMessage{Message: "Player Removed"}, rw)
	p.l.Println("[DEBUG:DELETE] Player Removed.")
}