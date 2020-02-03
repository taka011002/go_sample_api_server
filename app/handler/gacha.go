package handler

import (
	"encoding/json"
	"net/http"

	"github.com/taka011002/go_sample_api_server/app/domain/service"
)

// Userに対するHandlerのインターフェース
type GachaHandler interface {
	Draw(http.ResponseWriter, *http.Request)
}

type gachaHandlerImpl struct {
	gachaService service.GachaService
}

type drawParams struct {
	Times int `json:"times"`
}

func NewGachaHandler(gs service.GachaService) GachaHandler {
	return &gachaHandlerImpl{
		gachaService: gs,
	}
}

func (gs gachaHandlerImpl) Draw(w http.ResponseWriter, r *http.Request) {
	user, err := GetLoginUser(r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	drawParams := drawParams{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&drawParams); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := gs.gachaService.Draw(user, drawParams.Times)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, res)
}