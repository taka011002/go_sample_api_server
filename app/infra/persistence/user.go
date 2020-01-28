package persistence

import (
	"database/sql"
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
	"github.com/taka011002/go_sample_api_server/app/domain/repository"
)

type userPersistence struct {
	DB *sql.DB
}

func NewUserPersistence(DB *sql.DB) repository.UserRepository {
	return &userPersistence{DB: DB}
}

//ユーザ登録
func (up userPersistence) Create(userID, name, email string) error {
	stmt, err := up.DB.Prepare("INSERT INTO user(user_id, name, email) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, name, email)
	return err
}

//userIDによってユーザ情報を取得する
func (up userPersistence) GetByUsername(username string) (*entity.User, error) {
	row := up.DB.QueryRow("SELECT * FROM users WHERE username = ?", username)
	//row型をgolangで利用できる形にキャストする。
	return convertToUser(row)
}

//row型をuser型に紐づける
func convertToUser(row *sql.Row) (*entity.User, error) {
	user := entity.User{}
	err := row.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Phone, &user.UserStatus)
	if err != nil {
		return nil, err
	}
	return &user, nil
}