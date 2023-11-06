package noteRepository

import (
	"context"

	"github.com/BasbustDama/note-service/internal/entity"
	"github.com/BasbustDama/note-service/internal/repository/internal/note/internal/model"
	"github.com/uptrace/bun"
)

type noteRepository struct {
	DB bun.IDB
}

func NewRepository(db bun.IDB) *noteRepository {
	return &noteRepository{DB: db}
}

func (repository *noteRepository) Delete(ctx context.Context, id int) error {
	model := model.Note{ID: id}

	_, err := repository.DB.NewDelete().Model(&model).WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Insert implements NoteRepository.
func (repository *noteRepository) Insert(note *entity.Note) error {
	panic("unimplemented")
}

// SelectCount implements NoteRepository.
func (repository *noteRepository) SelectCount() (int, error) {
	panic("unimplemented")
}

// SelectMany implements NoteRepository.
func (repository *noteRepository) SelectMany(offset int, limit int) ([]entity.Note, error) {
	panic("unimplemented")
}

// SelectOne implements NoteRepository.
func (repository *noteRepository) SelectOne(id int) (entity.Note, error) {
	panic("unimplemented")
}

// Update implements NoteRepository.
func (repository *noteRepository) Update(id int, title *string, description *string) error {
	panic("unimplemented")
}
