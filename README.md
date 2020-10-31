# Matchmaking REST service

## Instructions
I am assuming that you already [Go](https://golang.org/) installed and wokring
 - Unzip to a subdirectory of you GOPATH environment variable.
 - Open a terminal\shell inside the _api_ folder.
 - Run `go run main.go` to start the server
 - Use culr to send HTTP requests to the server:
    - `curl localhost:9090/route` for GET requests
    - `curl localhost:9090/route -XPOST -d '{"json":"data"}'` for POST requests with a boby
    - `curl localhost:9090/route -XDELETE` for DELETE requests.

## User API

The service can be access through three main entry points:
- `sessions`
- `join`
- `leave`

### Sessions
The `/session` handle will return all the sessions that are currently running or that are awaiting more players.

example: `curl localhost:9090/sessions`

### Join
The `/join` hanlde will allow a player to join a session and wait for the match to start. If no sessions are available a new one will be generated for the player.
This message requires a Player JSON obect with a name defined or the player will not be added.
The id is automatically generated.

*Player format*
```
type Player struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	SessionID   int    `json:"session_id"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
	DeletedOn   string `json:"-"`
}
```

example: `curl localhost:9090/join -XPOST -d '{"name":"Player"}'`

### /leave
The `/leave` hanlde will allow a player to leave a session.

If the session will drop under the minimum allowed players it will be stopped and the remaining players will be joined to all other awaiting players.

If in this process the maximum players limit per session is reached, the the players that can't fit in that session will be added to a new one.

The leave handle will require the id of the player to be passes after the leave handle ( i.e. `/leave/15` ) or the request will be denied.

example: `curl localhost:9090/leave/9 -XDELETE`