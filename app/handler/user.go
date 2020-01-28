package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/taka011002/go_sample_api_server/app/domain/service"
	"net/http"
)

// Userに対するHandlerのインターフェース
type UserHandler interface {
	GetUser(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
	CreateUser(http.ResponseWriter, *http.Request)
}

type userHandlerImpl struct {
	userService service.UserService
}

func NewUserHandler(us service.UserService) UserHandler {
	return &userHandlerImpl{
		userService: us,
	}
}

func (uh userHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POST")
}

func (uh userHandlerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, err := uh.userService.GetByUsername(username)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, user)
}

func (uh userHandlerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POST")
}