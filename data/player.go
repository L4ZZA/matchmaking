package data

import (
	"fmt"
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

// ErrPlayerNotFound is an error raised when a Player can not be found in the database
var ErrPlayerNotFound = fmt.Errorf("Player not found")

// Player defines the structure for an API Player
type Player struct {
	ID          int     `json:"id"`
	Name        string `json:"name" validate:"required"`
	SessionID   int    `json:"session_id"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
	DeletedOn   string `json:"-"`
}

func (p *Player) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

func (p *Player) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Player) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// Players is a collection of Player
type Players map[int]*Player