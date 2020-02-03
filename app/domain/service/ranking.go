package service

import (
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
	"github.com/taka011002/go_sample_api_server/app/domain/repository"
)

type rankingServiceImpl struct {
	rankingRepository repository.RankingRepository
}

type RankingService interface {
	CharacterPower(user *entity.User) (*[]entity.Ranking, error)
}

func NewRankingService(r repository.RankingRepository) RankingService {
	return &rankingServiceImpl{rankingRepository: r}
}

func (rs rankingServiceImpl) CharacterPower(user *entity.User) (*[]entity.Ranking, error) {
	return rs.rankingRepository.CharacterPower()
}
