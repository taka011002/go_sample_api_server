package service

import (
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
	"github.com/taka011002/go_sample_api_server/app/domain/repository"
	"github.com/taka011002/go_sample_api_server/app/infra"
	"github.com/taka011002/go_sample_api_server/app/infra/persistence"
	"math/rand"
	"sort"
	"time"
)

type gachaServiceImpl struct {
	userCharacterRepository repository.UserCharacterRepository
}

type GachaService interface {
	Draw(user *entity.User, times int) ([]*entity.Character, error)
}

func NewGachaService(u repository.UserCharacterRepository) GachaService {
	return &gachaServiceImpl{userCharacterRepository: u}
}

func (gs gachaServiceImpl) Draw(user *entity.User, times int) ([]*entity.Character, error) {
	crper := persistence.NewCharacterRarityPersistence(infra.DB)
	characterRarities, err := crper.GetAll()

	if err != nil {
		return nil, err
	}

	per := persistence.NewUserCharacterPersistence(infra.DB)
	uc := NewUserCharacterService(per)

	cper := persistence.NewCharacterPersistence(infra.DB)

	var res []*entity.Character

	sort.Sort(sort.Reverse(characterRarities))

	// 抽選結果チェックの基準となる境界値を生成
	boundaries := make([]int, characterRarities.Len()+1)
	for i := 1; i < characterRarities.Len()+1; i++ {
		boundaries[i] = boundaries[i-1] + (*characterRarities)[i-1].Rarity
	}
	boundaries = boundaries[1:len(boundaries)]

	// times回抽選を行う
	rand.Seed(time.Now().UnixNano())
	n := times
	for i := 0; i < n; i++ {
		draw := rand.Intn(boundaries[len(boundaries)-1]) + 1
		for i, boundary := range boundaries {
			if draw <= boundary {
				c, err := cper.GetRand(i+1)

				if err != nil {
					return nil, err
				}

				res = append(res, c)

				break
			}
		}
	}

	err = uc.Creates(user, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
