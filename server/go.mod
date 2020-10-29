module server

go 1.15

replace example.com/handlers => ../handlers
replace example.com/data => ../data

require (
    example.com/handlers v0.0.0-00010101000000-000000000000
    example.com/data v0.0.0-00010101000000-000000000000
)