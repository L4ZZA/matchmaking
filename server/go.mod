module server

go 1.15

replace github/com/L4ZZA/matchmaking/handlers => ../handlers

replace github/com/L4ZZA/matchmaking/data => ../data

require (
	github.com/gorilla/mux v1.8.0
	github/com/L4ZZA/matchmaking/data v0.0.0-00010101000000-000000000000
	github/com/L4ZZA/matchmaking/handlers v0.0.0-00010101000000-000000000000
	golang.org/x/tools/gopls v0.5.2 // indirect
)
