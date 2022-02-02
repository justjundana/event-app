package comment

import (
	"database/sql"
)

type CommentRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}
