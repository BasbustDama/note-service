package main

import (
	"github.com/BasbustDama/note-service/config"
	"github.com/BasbustDama/note-service/internal/database/postgres"
	"github.com/BasbustDama/note-service/internal/handler"
	"github.com/BasbustDama/note-service/internal/usecase"
	"github.com/BasbustDama/note-service/pkg/server"
	"github.com/BasbustDama/note-service/pkg/sqlx"

	_ "github.com/lib/pq"
)

func main() {
	config := config.MustGetConfig("./")

	conn := sqlx.MustOpen(config.DatabaseDsn)
	defer conn.Close()

	postgresDatabase := postgres.New(conn)

	noteUsecase := usecase.New(postgresDatabase)

	handler := handler.New(noteUsecase)

	server := server.New(handler, config.HttpPort)
	server.Run()
}
