package repository

import (
	"context"

	"github.com/BasbustDama/note-service/internal/entity"
	customerRepository "github.com/BasbustDama/note-service/internal/repository/internal/customer"
	noteRepository "github.com/BasbustDama/note-service/internal/repository/internal/note"
	"github.com/uptrace/bun"
)

type (
	RepositoryManager interface {
		Begin() (RepositoryManager, error)
		Commit() error
		Rollback() error

		GetNoteRepository() NoteRepository
		GetCustomerRepository() CustomerRepository
	}

	CustomerRepository interface {
		Insert(username string, passwordHash string) (entity.Customer, error)
		SelectByID(id int) (entity.Customer, error)
		SelectByCreds(username string, passwordHash string) (entity.Customer, error)
	}

	NoteRepository interface {
		Insert(note *entity.Note) error
		SelectOne(id int) (entity.Note, error)
		SelectMany(offset, limit int) ([]entity.Note, error)
		SelectCount() (int, error)
		Update(id int, title *string, description *string) error
		Delete(ctx context.Context, id int) error
	}
)

type repositoryManager struct {
	DB *bun.DB
	Tx *bun.Tx
}

func NewManager(db *bun.DB, tx *bun.Tx) RepositoryManager {
	return &repositoryManager{
		DB: db,
		Tx: tx,
	}
}

func (manager *repositoryManager) Begin() (RepositoryManager, error) {
	tx, err := manager.DB.Begin()
	if err != nil {
		return nil, err
	}

	return NewManager(manager.DB, &tx), nil
}

func (manager *repositoryManager) clearTx() {
	manager.Tx = nil
}

func (manager *repositoryManager) Commit() error {
	defer manager.clearTx()
	return manager.Tx.Commit()
}

func (manager *repositoryManager) Rollback() error {
	defer manager.clearTx()
	return manager.Tx.Rollback()
}

func (manager *repositoryManager) getConnection() bun.IDB {
	if manager.Tx != nil {
		return manager.Tx
	}

	return manager.DB
}

func (manager *repositoryManager) GetCustomerRepository() CustomerRepository {
	return customerRepository.NewRepository(manager.getConnection())
}

func (manager *repositoryManager) GetNoteRepository() NoteRepository {
	return noteRepository.NewRepository(manager.getConnection())
}
