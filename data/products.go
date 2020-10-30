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
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
	DeletedOn   string `json:"-"`
}

func (p *Session) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Session) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

// Sessions is a collection of Session
type Sessions []*Session

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

// GetSessions returns a list of Sessions
func GetSessions() Sessions {
	return SessionList
}

func AddSession(p *Session) {
	p.ID = getNextID()
	SessionList = append(SessionList, p)
}

var ErrSessionNotFound = fmt.Errorf("Session not found")

func findSession(id int) (*Session, int, error) {
	for i, p := range SessionList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrSessionNotFound
}

func getNextID() int {
	lp := SessionList[len(SessionList)-1]
	return lp.ID + 1
}

// SessionList is a hard coded list of Sessions for this
// example data source
var SessionList = []*Session{
	&Session{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Session{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
