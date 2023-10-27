package main

import (
	"github.com/BasbustDama/note-service/config"
	"github.com/BasbustDama/note-service/internal/database/postgres"
	noteHandler "github.com/BasbustDama/note-service/internal/handler/note"
	noteUsecase "github.com/BasbustDama/note-service/internal/usecase/note"
	"github.com/BasbustDama/note-service/pkg/server"
	"github.com/BasbustDama/note-service/pkg/sqlx"

	_ "github.com/lib/pq"
)

func main() {
	config := config.MustGetConfig("./")

	conn := sqlx.MustOpen(config.DatabaseDsn)
	defer conn.Close()

	postgresDatabase := postgres.New(conn)

	noteUsecase := noteUsecase.New(postgresDatabase)

	handler := noteHandler.New(noteUsecase)

	server := server.New(handler, config.HttpPort)
	server.Run()
}
