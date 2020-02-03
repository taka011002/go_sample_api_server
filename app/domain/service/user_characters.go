package service

import (
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
	"github.com/taka011002/go_sample_api_server/app/domain/repository"
)

type userCharacterServiceImpl struct {
	userCharacterRepository repository.UserCharacterRepository
}

type UserCharacterService interface {
	Create(user *entity.User, character *entity.Character) error
	Creates(user *entity.User, characters []*entity.Character) error
	WhereByUser(user *entity.User) (*entity.UserCharacters, error)
}

func NewUserCharacterService(u repository.UserCharacterRepository) UserCharacterService {
	return &userCharacterServiceImpl{userCharacterRepository: u}
}

func (uc userCharacterServiceImpl) Create(user *entity.User, character *entity.Character) error {
	err := uc.userCharacterRepository.Create(user.Id, character.Id)

	if err != nil {
		return err
	}
	return nil
}

func (uc userCharacterServiceImpl) Creates(user *entity.User, characters []*entity.Character) error {
	var c []int
	for i := range characters {
		c = append(c, characters[i].Id)
	}

	err := uc.userCharacterRepository.Creates(user.Id, c)

	if err != nil {
		return err
	}
	return nil
}

func (uc userCharacterServiceImpl) WhereByUser(user *entity.User) (*entity.UserCharacters, error) {
	ucs, err := uc.userCharacterRepository.WhereByUserId(user.Id)

	if err != nil {
		return nil, err
	}
	return ucs, nil
}