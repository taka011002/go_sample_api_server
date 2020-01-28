package service

import (
	"github.com/google/uuid"
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
	"github.com/taka011002/go_sample_api_server/app/domain/repository"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
}

type UserService interface {
	Create(name, email string) error
	GetByUsername(userID string) (*entity.User, error)
}

func NewUserService(u repository.UserRepository) UserService {
	return &userServiceImpl{userRepository: u}
}

func (uu userServiceImpl) GetByUsername(username string) (*entity.User, error) {
	user, err := uu.userRepository.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userServiceImpl) Create(name, email string) error {
	//本来ならemailのバリデーションをする

	//一意でランダムな文字列を生成する
	userID, err := uuid.NewRandom()//返り値はuuid型
	if err != nil {
		return err
	}

	//domainを介してinfrastructureで実装した関数を呼び出す。
	// Persistence（Repository）を呼出
	err = uu.userRepository.Create(userID.String(), name, email)
	if err != nil {
		return err
	}
	return nil
}