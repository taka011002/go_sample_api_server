package handler

import (
	"net/http"

	"github.com/taka011002/go_sample_api_server/app/domain/service"
)

type RankingHandler interface {
	CharacterPower(http.ResponseWriter, *http.Request)
}

type rankingHandlerImpl struct {
	rankingService service.RankingService
}

func NewRankingHandler(rs service.RankingService) RankingHandler {
	return &rankingHandlerImpl{
		rankingService: rs,
	}
}

func (rh rankingHandlerImpl) CharacterPower(w http.ResponseWriter, r *http.Request) {
	user, err := GetLoginUser(r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ranking, err := rh.rankingService.CharacterPower(user)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, ranking)
}