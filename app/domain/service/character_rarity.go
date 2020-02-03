package service

import (
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
	"github.com/taka011002/go_sample_api_server/app/domain/repository"
)

type characterRarityServiceImpl struct {
	characterRarityRepository repository.CharacterRarityRepository
}

type CharacterRarityService interface {
	Create(cr *entity.CharacterRarity) error
	Update(cr *entity.CharacterRarity) error
	Delete(cr *entity.CharacterRarity) error
	CreateOrUpdate(cr *entity.CharacterRarity) error
}

func NewCharacterRarityService(r repository.CharacterRarityRepository) CharacterRarityService {
	return &characterRarityServiceImpl{characterRarityRepository: r}
}

func (cs characterRarityServiceImpl) Create(c *entity.CharacterRarity) error {
	if err := cs.characterRarityRepository.Create(c.Name, c.Rarity); err != nil {
		return err
	}
	return nil
}

func (cs characterRarityServiceImpl) Update(c *entity.CharacterRarity) error {
	if err := cs.characterRarityRepository.Update(c.Id, c.Name, c.Rarity); err != nil {
		return err
	}
	return nil
}

func (cs characterRarityServiceImpl) Delete(c *entity.CharacterRarity) error {
	err := cs.characterRarityRepository.Delete(c.Id)
	if err != nil {
		return err
	}
	return nil
}

func (cs characterRarityServiceImpl) CreateOrUpdate(c *entity.CharacterRarity) error {
	cc, err := cs.characterRarityRepository.GetByName(c.Name)

	if cc == nil {
		err = cs.Create(c)
	} else {
		c.Id = cc.Id
		err = cs.Update(c)
	}

	if err != nil {
		return err
	}

	return nil
}