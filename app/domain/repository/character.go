package repository

import (
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
)

type CharacterRepository interface {
	Create(name string) error
	Update(id int, name string) error
	GetByName(name string) (*entity.Character, error)
	Delete(id int) error
}