module hello

go 1.15

replace example.com/handlers => ../handlers

require example.com/handlers v0.0.0-00010101000000-000000000000 // indirect
