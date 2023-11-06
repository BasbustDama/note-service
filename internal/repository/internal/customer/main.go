package customerRepository

import (
	"github.com/BasbustDama/note-service/internal/entity"
	"github.com/uptrace/bun"
)

type CustomerRepository interface {
	Insert(username string, passwordHash string) (entity.Customer, error)
	SelectByID(id int) (entity.Customer, error)
	SelectByCreds(username string, passwordHash string) (entity.Customer, error)
}

type customerRepository struct {
	DB bun.IDB
}

func NewRepository(db bun.IDB) CustomerRepository {
	return &customerRepository{DB: db}
}

// Insert implements CustomerRepository.
func (*customerRepository) Insert(username string, passwordHash string) (entity.Customer, error) {
	panic("unimplemented")
}

// SelectByCreds implements CustomerRepository.
func (*customerRepository) SelectByCreds(username string, passwordHash string) (entity.Customer, error) {
	panic("unimplemented")
}

// SelectByID implements CustomerRepository.
func (*customerRepository) SelectByID(id int) (entity.Customer, error) {
	panic("unimplemented")
}
