package customerUsecase

import (
	"errors"

	"github.com/BasbustDama/note-service/internal/database"
	"github.com/BasbustDama/note-service/internal/entity"
	"github.com/sagikazarmark/slog-shim"
)

type (
	CustomerUsecase interface {
		Create(username string, password string) (entity.Customer, error)
		GetByID(id int) (entity.Customer, error)
		GetByCreds(username string, passwordHash string) (entity.Customer, error)
	}

	CustomerDatabase interface {
		Insert(username string, passwordHash string) (entity.Customer, error)
		SelectByID(id int) (entity.Customer, error)
		SelectByCreds(username string, passwordHash string) (entity.Customer, error)
	}

	PasswordGenerator interface {
		GeneratePassword(password string) (string, error)
	}
)

type customerUsecase struct {
	Database  CustomerDatabase
	Generator PasswordGenerator
}

func New(database CustomerDatabase) CustomerUsecase {
	return &customerUsecase{Database: database}
}

func (usecase *customerUsecase) Create(username string, password string) (entity.Customer, error) {
	passwordHash, err := usecase.Generator.GeneratePassword(password)
	if err != nil {
		slog.Error(err.Error())
		return entity.Customer{}, entity.ErrorInternal
	}

	customer, err := usecase.Database.Insert(username, passwordHash)
	if err != nil {
		slog.Error(err.Error())
		return entity.Customer{}, entity.ErrorInternal
	}

	return customer, nil
}

func (usecase *customerUsecase) GetByCreds(username string, passwordHash string) (entity.Customer, error) {
	customer, err := usecase.Database.SelectByCreds(username, passwordHash)
	if err != nil {
		if errors.Is(err, database.ErrorNotFound) {
			return entity.Customer{}, entity.ErrorNotFound
		}

		slog.Error(err.Error(), slog.String("username", username), slog.String("passwordHash", passwordHash))
		return entity.Customer{}, entity.ErrorInternal
	}

	return customer, nil
}

func (usecase *customerUsecase) GetByID(id int) (entity.Customer, error) {
	customer, err := usecase.Database.SelectByID(id)
	if err != nil {
		if errors.Is(err, database.ErrorNotFound) {
			return entity.Customer{}, entity.ErrorNotFound
		}

		slog.Error(err.Error(), slog.Int("customer_id", id))
		return entity.Customer{}, entity.ErrorInternal
	}

	return customer, nil
}
