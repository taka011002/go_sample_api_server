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

func (up userPersistence) Create(username string, password string) error {
	stmt, err := up.DB.Prepare("INSERT INTO users(id, username, password) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(0, username, password)
	return err
}

func (up userPersistence) Update(id int, username string, password string) error {
	stmt, err := up.DB.Prepare("UPDATE users SET username = ?, password = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username, password, id)
	return err
}

func (up userPersistence) Delete(id int) error {
	stmt, err := up.DB.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}

func (up userPersistence) GetByUsername(username string) (*entity.User, error) {
	row := up.DB.QueryRow("SELECT * FROM users WHERE username = ?", username)
	//row型をgolangで利用できる形にキャストする。
	return convertToUser(row)
}

func convertToUser(row *sql.Row) (*entity.User, error) {
	user := entity.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
