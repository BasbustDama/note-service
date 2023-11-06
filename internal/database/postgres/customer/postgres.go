package customerPostgres

import (
	"github.com/BasbustDama/note-service/internal/entity"
	"github.com/jmoiron/sqlx"
)

type PostgresDatabase struct {
	Connection *sqlx.DB
}

func New(connection *sqlx.DB) *PostgresDatabase {
	return &PostgresDatabase{Connection: connection}
}

func (database *PostgresDatabase) Insert(customer *entity.Customer) error {
	const query = "INSERT INTO customer (username, password) VALUES ($1, $2) RETURNING id"

	var customerId int
	err := database.Connection.QueryRowx(query, customer.Username, customer.Password).Scan(&customerId)
	if err != nil {
		return err
	}

	customer.ID = customerId
	return nil
}

func (database *PostgresDatabase) SelectByID(id int) (entity.Customer, error) {
	return entity.Customer{}, nil
}
