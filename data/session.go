package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-playground/validator"
)

var ErrSessionNotFound = fmt.Errorf("Session not found")

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
	id := getNextPlayerID()
	p.SessionID = ss.ID
	ss.Lobby[id] = p

	if(len(ss.Lobby) >= EnoughPlayers){
		// start game session
		// TODO: defer session activation until lobby is full (100 players) or countdown expires
		ss.IsWaiting = false
	}
	return nil
}


// RemovePlayer deletes a Player from the storage
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
		delete(s.Lobby, pi)
		RemovedIDs = append(RemovedIDs, pi)
	} else {
		return ErrPlayerNotFound
	}

	lobbySize = len(s.Lobby)
	if(lobbySize < EnoughPlayers){
		// end game session
		s.IsWaiting = true
	}

	return nil
}

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
	// how to check for map elements https://stackoverflow.com/a/2050629/6120464
	if p, ok := s.Lobby[id]; ok {
		return p, id, nil
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

// getNextPlayerID returns the first available id starting from zero.
// if the id has been flagged as deleted it gets recycled.
func getNextPlayerID() int {
	_, err:= findAvailableSession()
	if(err == nil){
		idsLeft := len(RemovedIDs)
		if idsLeft > 0 {
			// fetch the now available id
			id := RemovedIDs[idsLeft-1]
			// remove it from list
			RemovedIDs = RemovedIDs[:idsLeft-1]
			// return it
			return id
		}
	}

	LastUsedID++
	return LastUsedID
}

// SessionList is a hard coded list of Sessions for this
// example data source
var SessionList = Sessions{}
// contains a list of all the ids of the removed players
var RemovedIDs = []int{}
// caching last used id assigned to give a fallback for the id recycling algorithm
var LastUsedID int = -1
