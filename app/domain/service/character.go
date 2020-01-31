package service

import (
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
	"github.com/taka011002/go_sample_api_server/app/domain/repository"
)

type characterServiceImpl struct {
	characterRepository repository.CharacterRepository
}

type CharacterService interface {
	Create(user *entity.Character) error
	Update(user *entity.Character) error
	Delete(user *entity.Character) error
	GetByName(name string) (*entity.Character, error)
}

func NewCharacterService(r repository.CharacterRepository) CharacterService {
	return &characterServiceImpl{characterRepository: r}
}

func (cs characterServiceImpl) Create(c *entity.Character) error {
	if err := cs.characterRepository.Create(c.Name); err != nil {
		return err
	}
	return nil
}

func (cs characterServiceImpl) Update(c *entity.Character) error {
	if err := cs.characterRepository.Update(c.Id, c.Name); err != nil {
		return err
	}
	return nil
}

func (cs characterServiceImpl) GetByName(n string) (*entity.Character, error) {
	c, err := cs.characterRepository.GetByName(n)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cs characterServiceImpl) Delete(c *entity.Character) error {
	err := cs.characterRepository.Delete(c.Id)
	if err != nil {
		return err
	}
	return nil
}
