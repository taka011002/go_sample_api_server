package handler

import (
	"encoding/json"
	"net/http"

	"github.com/taka011002/go_sample_api_server/app/domain/entity"
	"github.com/taka011002/go_sample_api_server/app/domain/service"
)

// Userに対するHandlerのインターフェース
type UserHandler interface {
	GetUser(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	SignUp(http.ResponseWriter, *http.Request)
	SignIn(http.ResponseWriter, *http.Request)
}

type userHandlerImpl struct {
	userService service.UserService
}

type loginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserHandler(us service.UserService) UserHandler {
	return &userHandlerImpl{
		userService: us,
	}
}

func (uh userHandlerImpl) SignUp(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := uh.userService.Create(&user); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := CreateToken(user.Username)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, map[string]string{"token": token})
}

func (uh userHandlerImpl) SignIn(w http.ResponseWriter, r *http.Request) {
	loginUser := loginUser{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginUser); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := uh.userService.SignIn(loginUser.Username, loginUser.Password); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := CreateToken(loginUser.Username)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (uh userHandlerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := GetLoginUser(r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user.Password = ""
	respondJSON(w, http.StatusOK, user)
}

func (uh userHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	loginUser, err := GetLoginUser(r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user := entity.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user.Id = loginUser.Id
	if err := uh.userService.Update(&user); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusNoContent, nil)
}

func (uh userHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	user, err := GetLoginUser(r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := uh.userService.Delete(user); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusNoContent, nil)
}
