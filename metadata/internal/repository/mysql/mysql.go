package mysql

import (
	"context"
	"database/sql"

	"github.com/cornelia247/nyxfilms/metadata/internal/repository"
	"github.com/cornelia247/nyxfilms/metadata/pkg/model"
	_ "github.com/go-sql-driver/mysql"
)

// Repository defines a MYSQL-based film metadate repository.
type Repository struct {
	db *sql.DB
}

// New creates a new MYSQL-based repository.
func New() (*Repository, error) {
	db, err := sql.Open("mysql", "root:password@/nyx")
	if err != nil {
		return nil, err
	}
	return &Repository{db}, nil
}

// Get retrives film metadata by film id
func (r *Repository) Get(ctx context.Context, id string) (*model.Metadata, error) {
	var title, description, director string
	row := r.db.QueryRowContext(ctx, "SELECT title, description, director FROM films WHERE id = ?", id)
	if err := row.Scan(&title, &description, &director); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &model.Metadata{
		ID:          id,
		Title:       title,
		Description: description,
		Director:    director,
	}, nil

}

// Put adds film metadata for a given film id
func (r *Repository) Put(ctx context.Context, id string, metadata *model.Metadata) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO films (id, title,description, director) VALUES(?,?,?,?)", id, metadata.Title, metadata.Description, metadata.Director)
	return err
}
