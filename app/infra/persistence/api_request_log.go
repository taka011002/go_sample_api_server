package persistence

import (
	"database/sql"
	"github.com/taka011002/go_sample_api_server/app/domain/repository"
)

type apiRequestLogPersistence struct {
	DB *sql.DB
}

func NewApiRequestLogPersistence(DB *sql.DB) repository.ApiRequestLogRepository {
	return &apiRequestLogPersistence{DB: DB}
}

func (ap apiRequestLogPersistence) Create(userId int, method string, path string, params string) error {
	stmt, err := ap.DB.Prepare("INSERT INTO api_request_logs(id, user_id, method, path, params) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(0, userId, method, path, params)
	return err
}
