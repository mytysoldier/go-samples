package db

import "unittests/db/model"

type UserInterface interface {
	GetUserById(id int) model.User
}
