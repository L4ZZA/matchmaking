package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-playground/validator"
)

// Session defines the structure for an API Session
type Session struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Lobby 		Players `json:"lobby"`
	IsWaiting 	bool 	`json:"is_waiting"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (s *Session) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(s)
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Sessions) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (s *Session) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}

// Sessions is a collection of Session
type Sessions []*Session

const EnoughPlayers int = 4

func New() Session {
	s:= Session{}
	s.ID = getNextSessionID()
	s.Name = fmt.Sprintf("Session%d", s.ID)
	s.Lobby = Players{}
	s.IsWaiting = true
	s.CreatedOn = time.Now().UTC().String()
	s.UpdatedOn = time.Now().UTC().String()
	return s
}

// GetSessions returns a list of Sessions
func GetSessions() Sessions {
	return SessionList
}

func AddSession(s *Session) {
	s.ID = getNextSessionID()
	s.Name = fmt.Sprintf("Session%d", s.ID)
	s.Lobby = nil
	SessionList = append(SessionList, s)
}

func CreateSession() *Session {
	s:= New()
	SessionList = append(SessionList, &s)
	return &s
}

func AddPlayer(p *Player) error {
	ss, err := findAvailableSession()
	if(err != nil){
		return err
	}
	p.ID = getNextPlayerID()
	p.SessionID = ss.ID
	ss.Lobby = append(ss.Lobby, p)

	if(len(ss.Lobby) >= EnoughPlayers){
		ss.IsWaiting = false
	}
	return nil
}


// RemovePlayer deletes a Player from the database
func RemovePlayer(sessionId int, playerId int) error {

	s, si, _ := findSession(sessionId)
	if si == -1 {
		return ErrSessionNotFound
	}

	lobbySize := len(s.Lobby)
	if(lobbySize > 0) {
		_, pi, _ := findPlayer(playerId, s)
		if pi == -1 {
			return ErrPlayerNotFound
		}

		if(lobbySize == 1){
			s.Lobby = Players{}
		} else {
			s.Lobby = append(s.Lobby[:pi], s.Lobby[pi+1])
		}
	} else {
		return ErrPlayerNotFound
	}


	return nil
}

var ErrSessionNotFound = fmt.Errorf("Session not found")

func findSession(id int) (*Session, int, error) {
	if(len(SessionList) > 0){
		for i, s := range SessionList {
			if s.ID == id {
				return s, i, nil
			}
		}
	}

	return nil, -1, ErrSessionNotFound
}

func findPlayer(id int, s *Session) (*Player, int, error) {
	for i, p := range s.Lobby {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrPlayerNotFound
}

func findAvailableSession() (*Session, error) {
	for _, s := range SessionList {
		if s.IsWaiting == true {
			return s, nil
		}
	}

	return nil, ErrSessionNotFound
}

func getNextSessionID() int {
	return len(SessionList)
}

func getNextPlayerID() int {
	s, err:= findAvailableSession()
	if(err != nil){
		return 0
	}
	return len(s.Lobby)
}

// SessionList is a hard coded list of Sessions for this
// example data source
var SessionList = Sessions{}
