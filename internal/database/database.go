package database

import (
	"errors"

	"github.com/BasbustDama/note-service/internal/entity"
)

var ErrorNotFound = errors.New("rows not found")

type NoteModel struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

func (model NoteModel) ToEntity() entity.Note {
	return entity.Note{
		ID:          model.ID,
		Title:       model.Title,
		Description: model.Description,
	}
}

func ModelListToEntity(modelList []NoteModel) []entity.Note {
	entityList := make([]entity.Note, 0, len(modelList))
	for _, model := range modelList {
		entityList = append(entityList, model.ToEntity())
	}

	return entityList
}
