package model

import (
	"github.com/BasbustDama/note-service/internal/entity"
	"github.com/uptrace/bun"
)

type Note struct {
	bun.BaseModel `bun:"table:note,alias:n"`

	ID          int    `bun:"id,pk"`
	CustomerID  int    `bun:"customer_id,notnull"`
	Title       string `bun:"title,notnull"`
	Description string `bun:"description"`
}

func NewNote(note entity.Note) Note {
	return Note{
		ID:          note.ID,
		CustomerID:  note.CustomerID,
		Title:       note.Title,
		Description: note.Description,
	}
}

func (model Note) ToEntity() entity.Note {
	return entity.Note{
		ID:          model.ID,
		CustomerID:  model.CustomerID,
		Title:       model.Title,
		Description: model.Description,
	}
}
