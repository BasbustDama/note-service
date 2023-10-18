package usecase

import (
	"errors"
	"log"
	"log/slog"

	"github.com/BasbustDama/note-service/internal/database"
	"github.com/BasbustDama/note-service/internal/entity"
)

type (
	NoteUsecase interface {
		Create(title string, description string) (entity.Note, error)
		Get(id int) (entity.Note, error)
		GetList(offset, limit int) ([]entity.Note, int, error)
		Update(title, description *string) error
		Delete(id int) error
	}

	NoteDatabase interface {
		Insert(note *entity.Note) error
		SelectOne(id int) (entity.Note, error)
		SelectMany(offset, limit int) ([]entity.Note, int, error)
		Update(title, description *string) error
		Delete(id int) error
	}
)

type noteUsecase struct {
	Database NoteDatabase
}

func New(database NoteDatabase) NoteUsecase {
	return &noteUsecase{Database: database}
}

func (usecase *noteUsecase) Create(title string, description string) (entity.Note, error) {
	note := entity.Note{
		Title:       title,
		Description: description,
	}

	err := usecase.Database.Insert(&note)
	if err != nil {
		slog.Error(err.Error())
		return entity.Note{}, entity.ErrorInternal
	}

	return note, nil
}

func (usecase *noteUsecase) Delete(id int) error {
	if err := usecase.Database.Delete(id); err != nil {
		return entity.ErrorInternal
	}

	return nil
}

func (usecase *noteUsecase) Get(id int) (entity.Note, error) {
	note, err := usecase.Database.SelectOne(id)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, database.ErrorNotFound) {
			return entity.Note{}, entity.ErrorNotFound
		}

		return entity.Note{}, entity.ErrorInternal
	}

	return note, nil
}

func (usecase *noteUsecase) GetList(offset int, limit int) ([]entity.Note, int, error) {
	panic("unimplemented")
}

func (usecase *noteUsecase) Update(title *string, description *string) error {
	panic("unimplemented")
}
