package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Representing a repo for storing movies
type MovieRepository interface {
	Insert(movie *Movie) error
	Get(id int64) (*Movie, error)
	Update(movie *Movie) error
	Delete(id int64) error
	GetAll(title string, genres []string, filters Filters) ([]*Movie, error)
}

// Struct with all basic repos
type Models struct {
	Movies MovieRepository
}

// Fabric function to create a new Models instance
func NewModels(db *sql.DB) *Models {
	return &Models{
		Movies: &MovieModel{DB: db},
	}
}
