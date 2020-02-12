package repository

import "github.com/taka011002/go_sample_api_server/app/domain/entity"

type RankingRepository interface {
	CharacterPower() (*[]entity.Ranking, error)
}
