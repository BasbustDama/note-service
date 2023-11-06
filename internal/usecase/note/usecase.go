package noteUsecase

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/BasbustDama/note-service/internal/database"
	"github.com/BasbustDama/note-service/internal/entity"
	"github.com/BasbustDama/note-service/internal/repository"
)

type (
	NoteUsecase interface {
		Create(title string, description string) (entity.Note, error)
		Get(id int) (entity.Note, error)
		GetList(offset, limit int) ([]entity.Note, int, error)
		Update(id int, title *string, description *string) error
		Delete(id int) error
	}
)

type noteUsecase struct {
	Database repository.RepositoryManager

	defaultTimeout time.Duration
}

func New(database repository.RepositoryManager, timeout time.Duration) NoteUsecase {
	return &noteUsecase{Database: database, defaultTimeout: timeout}
}

func (usecase *noteUsecase) Create(title string, description string) (entity.Note, error) {
	note := entity.Note{
		Title:       title,
		Description: description,
	}

	if err := usecase.Database.GetNoteRepository().Insert(&note); err != nil {
		slog.Error(err.Error())
		return entity.Note{}, entity.ErrorInternal
	}

	return note, nil
}

func (usecase *noteUsecase) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), usecase.defaultTimeout)
	defer cancel()

	if err := usecase.Database.GetNoteRepository().Delete(ctx, id); err != nil {
		return entity.ErrorInternal
	}

	return nil
}

func (usecase *noteUsecase) Get(id int) (entity.Note, error) {
	note, err := usecase.Database.GetNoteRepository().SelectOne(id)
	if err != nil {
		if errors.Is(err, database.ErrorNotFound) {
			return entity.Note{}, entity.ErrorNotFound
		}

		slog.Error(err.Error(), slog.Int("note_id", id))
		return entity.Note{}, entity.ErrorInternal
	}

	return note, nil
}

func (usecase *noteUsecase) GetList(offset int, limit int) ([]entity.Note, int, error) {
	noteRepository := usecase.Database.GetNoteRepository()
	note, err := noteRepository.SelectMany(offset, limit)
	if err != nil {
		slog.Error(err.Error())
		return nil, 0, entity.ErrorInternal
	}

	count, err := noteRepository.SelectCount()
	if err != nil {
		slog.Error(err.Error())
		return nil, 0, entity.ErrorInternal
	}

	return note, count, nil
}

func (usecase *noteUsecase) Update(id int, title *string, description *string) error {
	noteRepository := usecase.Database.GetNoteRepository()
	if _, err := noteRepository.SelectOne(id); err != nil {
		if errors.Is(err, database.ErrorNotFound) {
			return entity.ErrorNotFound
		}

		slog.Error(err.Error(), slog.Int("note_id", id))
		return entity.ErrorInternal
	}

	err := noteRepository.Update(id, title, description)
	if err != nil {
		slog.Error(err.Error())
		return entity.ErrorInternal
	}

	return nil
}
