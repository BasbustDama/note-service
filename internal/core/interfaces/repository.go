package interfaces

import "github.com/BasbustDama/note-service/internal/core/entity"

type (
	RepositoryManager interface {
		GetCustomerRepository() CustomerRepository
		GetNoteRepository() NoteRepository
	}

	CustomerRepository interface {
		Create(user entity.Customer) (int, error)
		Update(user entity.Customer) error
	}

	NoteRepository interface {
		Create(note entity.Note) (int, error)
		Get(userId, noteId int) (entity.Note, error)
		GetList(userId int) ([]entity.Note, error)
		Update(note entity.Note) error
		Delete(userId, noteId int) error
	}
)
