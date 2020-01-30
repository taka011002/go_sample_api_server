package repository

type ApiRequestLogRepository interface {
	Create(userId int, method string, path string, params string) error
}