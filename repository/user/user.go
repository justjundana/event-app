package user

import (
	"database/sql"

	_models "github.com/justjundana/event-planner/models"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user _models.User) (_models.User, error) {
	_, err := r.db.Exec("INSERT INTO users(name,email,password,address,occupation) VALUES(?,?,?,?,?)", user.Name, user.Email, user.Password, user.Address, user.Occupation)
	return user, err
}
