package handlers

import (
	"net/http"

	"matchmaking.test/data"
)

func (s *Sessions) Greetings(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("[DEBUG:GET] Greet new player")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Thanks for joining THE GAME. Use the handle /join to enter the first available lobby."))
}


// GetSessions returns the Sessions from the data store
func (s *Sessions) GetSessions(rw http.ResponseWriter, r *http.Request) {

	s.l.Println("[DEBUG:GET] GetSessions")
	// fetch the Sessions from the datastore
	ss := data.GetSessions()

	// serialize the list to JSON
	err := ss.ToJSON(rw)

	if err != nil {
		s.l.Println("[ERROR:GET] GetSessions - Failed to parse to JSON", ss, err)
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	s.l.Println("[DEBUG:GET] GetSessions - COMPLETED")
}