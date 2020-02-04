package service

import (
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
	"github.com/taka011002/go_sample_api_server/app/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
}

type UserService interface {
	Create(user *entity.User) error
	SignIn(username string, password string) error
	Update(user *entity.User) error
	Delete(user *entity.User) error
	GetByUsername(userID string) (*entity.User, error)
}

func NewUserService(u repository.UserRepository) UserService {
	return &userServiceImpl{userRepository: u}
}

func (uu userServiceImpl) Create(user *entity.User) error {
	//TODO 色々バリデーションをする

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	err = uu.userRepository.Create(user.Username, string(hashedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (uu userServiceImpl) SignIn(username string, password string) error {
	user, err := uu.userRepository.GetByUsername(username)
	if err != nil {
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}

func (uu userServiceImpl) Update(user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	err = uu.userRepository.Update(user.Id, user.Username, string(hashedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (uu userServiceImpl) GetByUsername(username string) (*entity.User, error) {
	user, err := uu.userRepository.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userServiceImpl) Delete(user *entity.User) error {
	err := uu.userRepository.Delete(user.Id)
	if err != nil {
		return err
	}
	return nil
}