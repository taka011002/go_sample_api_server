package persistence

import (
	"database/sql"
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
	"github.com/taka011002/go_sample_api_server/app/domain/repository"
)

type rankingPersistence struct {
	DB *sql.DB
}

func NewRankingPersistence(DB *sql.DB) repository.RankingRepository {
	return &rankingPersistence{DB: DB}
}

func (rp rankingPersistence) CharacterPower() (*[]entity.Ranking, error) {
	q := `SELECT uc.user_id, SUM(ch.power) AS power_sum
	FROM user_character AS uc
	INNER JOIN characters AS ch ON uc.character_id = ch.id
	GROUP BY uc.user_id
	ORDER BY power_sum DESC;`

	rows, err := rp.DB.Query(q)
	if err != nil {
		return nil, err
	}

	var res []entity.Ranking
	i := 1
	for rows.Next() {
		r := entity.Ranking{}
		err := rows.Scan(&r.User.Id,&r.Score)

		if err != nil {
			return nil, err
		}

		r.Rank = i
		res = append(res, r)
		i++
	}

	return &res, nil
}