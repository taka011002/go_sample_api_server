package handler

import (
	"net/http"

	"github.com/taka011002/go_sample_api_server/app/domain/service"
)

// Userに対するHandlerのインターフェース
type CharacterHandler interface {
	List(http.ResponseWriter, *http.Request)
}

type characterHandlerImpl struct {
	userCharacterService service.UserCharacterService
}

func NewCharacterHandler(ucs service.UserCharacterService) CharacterHandler {
	return &characterHandlerImpl{
		userCharacterService: ucs,
	}
}

func (ch characterHandlerImpl) List(w http.ResponseWriter, r *http.Request) {
	user, err := GetLoginUser(r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ucs, err := ch.userCharacterService.WhereByUser(user)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, ucs)
}