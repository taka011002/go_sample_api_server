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

func (up userPersistence) Create(username string, firstName string, lastName string, email string, password string, phone string, userStatus int) error {
	stmt, err := up.DB.Prepare("INSERT INTO users(id, username, first_name, last_name, email, password, phone, user_status) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(0, username, firstName, lastName, email, password, phone, userStatus)
	return err
}

func (up userPersistence) Update(id int, username string, firstName string, lastName string, email string, password string, phone string, userStatus int) error {
	stmt, err := up.DB.Prepare("UPDATE users SET username = ?, first_name = ?, last_name = ?, email = ?, password = ?, phone = ?, user_status = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username, firstName, lastName, email, password, phone, userStatus, id)
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
	err := row.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Phone, &user.UserStatus)
	if err != nil {
		return nil, err
	}
	return &user, nil
}