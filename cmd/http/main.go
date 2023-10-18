package main

import (
	"github.com/BasbustDama/note-service/internal/database/postgres"
	"github.com/BasbustDama/note-service/internal/handler"
	"github.com/BasbustDama/note-service/internal/usecase"
	"github.com/BasbustDama/note-service/pkg/server"
	"github.com/BasbustDama/note-service/pkg/sqlx"

	_ "github.com/lib/pq"
)

func main() {
	conn := sqlx.MustOpen("postgres://postgres:postgres@192.168.3.28:5432/note_database?sslmode=disable")
	defer conn.Close()

	postgresDatabase := postgres.New(conn)

	noteUsecase := usecase.New(postgresDatabase)

	handler := handler.New(noteUsecase)

	server := server.New(handler, ":8080")
	server.Run()
}
