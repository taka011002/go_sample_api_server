package service

import (
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
	"github.com/taka011002/go_sample_api_server/app/domain/repository"
)

type apiRequestLogServiceImpl struct {
	apiRequestLogRepository repository.ApiRequestLogRepository
}

type ApiRequestLogService interface {
	Create(log entity.ApiRequestLog) error
}

func NewApiRequestLogService(a repository.ApiRequestLogRepository) ApiRequestLogService {
	return &apiRequestLogServiceImpl{apiRequestLogRepository: a}
}

func (as apiRequestLogServiceImpl) Create(log entity.ApiRequestLog) error {
	err := as.apiRequestLogRepository.Create(log.UserId, log.Method, log.Path, log.Params)
	if err != nil {
		return err
	}
	return nil
}
