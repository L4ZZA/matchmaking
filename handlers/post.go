package handlers

import (
	"net/http"

	"matchmaking.test/data"
)

func (s *Sessions) AddSession(rw http.ResponseWriter, r *http.Request) {

	sess := r.Context().Value(KeySession{}).(data.Session)
	data.AddSession(&sess)
	s.l.Println("[DEBUG:POST] AddSession - COMPLETED")
	rw.WriteHeader(http.StatusOK)
	data.ToJSON(&GenericSuccessMessage{Message: "Session Added"}, rw)
}


func (s *Sessions) AddPlayer(rw http.ResponseWriter, r *http.Request) {

	player := r.Context().Value(KeyPlayer{}).(data.Player)

	s.l.Println("[DEBUG:POST] Adding player..")
	err := data.AddPlayer(&player)
	if(err != nil){

		s.l.Println("[DEBUG:POST] No Idle session available. Creating new session..")
		data.CreateSession()

		s.l.Println("[DEBUG:POST] Retrying to add the player..")
		err := data.AddPlayer(&player)
		if(err != nil){

			s.l.Println("[ERROR:POST] Player could not be created again -", err)
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}
	}
	rw.WriteHeader(http.StatusOK)
	data.ToJSON(&GenericSuccessMessage{Message: "Player Added."}, rw)
	s.l.Println("[DEBUG:POST] Player Added.")
}