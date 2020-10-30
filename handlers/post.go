package handlers

import (
	"net/http"

	"github/com/L4ZZA/matchmaking/data"
)

func (s *Sessions) AddSession(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("[DEBUG] Handle POST Session")

	sess := r.Context().Value(KeySession{}).(data.Session)
	data.AddSession(&sess)
	s.l.Println("[DEBUG] AddSession - COMPLETED")
}


func (s *Sessions) AddPlayer(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("[DEBUG] Handle POST Player")

	player := r.Context().Value(KeyPlayer{}).(data.Player)
	err := data.AddPlayer(&player)
	if(err != nil){
		s.l.Println("[DEBUG] No Idle session available -", err)
		s.l.Println("[DEBUG] Creating new session")

		data.CreateSession()
		err := data.AddPlayer(&player)
		if(err != nil){
			s.l.Println("[ERROR] Player could not be created again -", err)
			return
		}
	}
	s.l.Println("[DEBUG] Player added succesfully")
}