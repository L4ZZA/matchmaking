module server

go 1.15

replace matchmaking.test/handlers => ../handlers

replace matchmaking.test/data => ../data

require (
	github.com/gorilla/mux v1.8.0
	matchmaking.test/data v0.0.0-00010101000000-000000000000
	matchmaking.test/handlers v0.0.0-00010101000000-000000000000
)
