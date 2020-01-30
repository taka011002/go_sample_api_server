package repository

import (
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
)

type UserRepository interface {
	Create(username string, firstName string, lastName string, email string, password string, phone string, userStatus int) error
	Update(id int, username string, firstName string, lastName string, email string, password string, phone string, userStatus int) error
	GetByUsername(username string) (*entity.User, error)
	Delete(id int) error
}