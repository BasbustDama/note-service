package customerHandler

import "github.com/BasbustDama/note-service/internal/entity"

type GetByIDCustomer interface {
	GetByID(customerID int) (entity.Customer, error)
}

func NewGetByIDCustomer() {}
