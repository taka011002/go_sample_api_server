package repository

import "github.com/taka011002/go_sample_api_server/app/domain/entity"

type UserCharacterRepository interface {
	Create(userId int, characterId int) error
	Creates(userId int, characterId []int) error
	WhereByUserId(userId int) (*entity.UserCharacters, error)
}