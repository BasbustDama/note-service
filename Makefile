migrate-up: 
	migrate -path ./migrations -database postgres://postgres:postgres@localhost:5432/notes?sslmode=disable up

migrate-down:
	migrate -path ./migrations -database postgres://postgres:postgres@localhost:5432/notes?sslmode=disable down