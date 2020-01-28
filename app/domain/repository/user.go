package repository

import (
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
)

type UserRepository interface {
	Create(userID, name, email string) error
	GetByUsername(username string) (*entity.User, error)
}