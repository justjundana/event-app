package user

import (
	_models "github.com/justjundana/event-planner/models"
)

type UserInterface interface {
	Login(email string, password string) (_models.User, error)
}
