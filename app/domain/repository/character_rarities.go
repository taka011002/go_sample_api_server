package repository

import "github.com/taka011002/go_sample_api_server/app/domain/entity"

type CharacterRarityRepository interface {
	Create(name string, rarity int) error
	Update(id int, name string, rarity int) error
	Delete(id int) error
	GetByName(name string) (*entity.CharacterRarity, error)
	GetAll() (*entity.CharacterRarities, error)
}
