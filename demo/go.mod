module example.com/pineapple-demo

go 1.23.0

toolchain go1.23.3

replace github.com/AnishG-git/HackUTD2024/backend => ../backend

require (
	github.com/AnishG-git/HackUTD2024/backend v0.0.0-00010101000000-000000000000
	github.com/gorilla/handlers v1.5.2
	github.com/gorilla/mux v1.8.1
)

require (
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
)
