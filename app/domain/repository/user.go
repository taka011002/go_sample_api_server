package repository

import (
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
)

type UserRepository interface {
	Create(username string, password string) error
	Update(id int, username string, password string) error
	GetByUsername(username string) (*entity.User, error)
	Delete(id int) error
}
