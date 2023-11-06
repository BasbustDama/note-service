package main

import (
	"time"

	"github.com/BasbustDama/note-service/config"
	customerPostgres "github.com/BasbustDama/note-service/internal/database/postgres/customer"
	notePostgres "github.com/BasbustDama/note-service/internal/database/postgres/note"
	handler "github.com/BasbustDama/note-service/internal/handler"
	customerUsecase "github.com/BasbustDama/note-service/internal/usecase/customer"
	noteUsecase "github.com/BasbustDama/note-service/internal/usecase/note"
	"github.com/BasbustDama/note-service/pkg/server"
	"github.com/BasbustDama/note-service/pkg/sqlx"

	_ "github.com/lib/pq"
)

const (
	defaultUsecaseTimeout = 5 * time.Second
)

func main() {
	config := config.MustGetConfig("./")

	conn := sqlx.MustOpen(config.DatabaseDsn)
	defer conn.Close()

	notePostgresDatabase := notePostgres.New(conn)
	customerPostgresDatabase := customerPostgres.New(conn)

	noteUsecase := noteUsecase.New(notePostgresDatabase, defaultUsecaseTimeout)
	customerUsecase := customerUsecase.New(customerPostgresDatabase)

	handler := handler.New(noteUsecase, customerUsecase)

	server := server.New(handler, config.HttpPort)
	server.Run()
}
