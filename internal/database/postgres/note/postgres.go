package notePostgres

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	core "github.com/BasbustDama/note-service/internal/database"
	"github.com/BasbustDama/note-service/internal/entity"
	"github.com/jmoiron/sqlx"
)

type PostgresDatabase struct {
	Connection *sqlx.DB
}

func New(connection *sqlx.DB) *PostgresDatabase {
	return &PostgresDatabase{Connection: connection}
}

func (database *PostgresDatabase) Insert(note *entity.Note) error {
	const query = "INSERT INTO note (title, description) VALUES ($1, $2) RETURNING id"

	var noteId int
	err := database.Connection.QueryRowx(query, note.Title, note.Description).Scan(&noteId)
	if err != nil {
		return err
	}

	note.ID = noteId
	return nil
}

func (database *PostgresDatabase) SelectOne(id int) (entity.Note, error) {
	const query = "SELECT id, title, description FROM note WHERE id = $1"

	var model core.NoteModel
	if err := database.Connection.Get(&model, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Note{}, core.ErrorNotFound
		}

		return entity.Note{}, err
	}

	return model.ToEntity(), nil
}

func (database *PostgresDatabase) SelectMany(offset, limit int) ([]entity.Note, error) {
	const query = "SELECT id, title, description FROM note ORDER BY id ASC LIMIT $1 OFFSET $2"

	var modelList []core.NoteModel
	err := database.Connection.Select(&modelList, query, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return core.ModelListToEntity(modelList), nil
}

func (database *PostgresDatabase) SelectCount() (int, error) {
	const query = "SELECT COUNT(id) FROM note"

	var count int
	err := database.Connection.Get(&count, query)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (database *PostgresDatabase) Update(id int, title *string, description *string) error {
	fields := make([]string, 0, 1)
	args := make([]any, 0, 2)

	if title != nil {
		fields = append(fields, "title")
		args = append(args, *title)
	}

	if description != nil {
		fields = append(fields, "description")
		args = append(args, *description)
	}

	args = append(args, id)

	setValues := make([]string, len(fields))
	for index, field := range fields {
		setValues[index] = fmt.Sprintf("%s = $%d", field, index+1)
	}

	query := fmt.Sprintf("UPDATE note SET %s WHERE id = $%d",
		strings.Join(setValues, ", "), len(setValues)+1)

	_, err := database.Connection.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (database *PostgresDatabase) Delete(id int) error {
	const query = "DELETE FROM note WHERE id = $1"

	_, err := database.Connection.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
