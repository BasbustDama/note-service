package interfaces

import "github.com/BasbustDama/note-service/internal/core/entity"

type (
	CustomerUseCase interface {
		SignIn(email, password string) (int, error)
		SignUp(email, password string) (int, error)
	}

	NoteUseCase interface {
		CreateNote(userId int, title string, description string) error
		GetNote(userId int, noteId int) (entity.Note, error)
		GetNotes(userId int) (entity.Note, error)
		Update(userId int, title string, description string) error
		Delete(userId int, noteId int) error
	}
)
