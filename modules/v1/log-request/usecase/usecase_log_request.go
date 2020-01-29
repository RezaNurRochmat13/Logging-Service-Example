package usecase

import (
	"svc-logger-go/modules/v1/log-request/dao"
	"svc-logger-go/modules/v1/log-request/repository"
	"svc-logger-go/util"
)

type logRequestUseCaseImpl struct {
	LogRequestRepository repository.Repository
}

func NewLogRequestUseCase(logRequestRepo repository.Repository) UseCase {
	return &logRequestUseCaseImpl{LogRequestRepository: logRequestRepo}
}

func (lu *logRequestUseCaseImpl) FindAllLogRequests() ([]dao.ListLogRequest, error) {
	findAllLogRequestsRepo, errorHandlerRepo := lu.LogRequestRepository.FindAll()
	if errorHandlerRepo != nil {
		util.LoggerOutput("Error when get all log", "Error", errorHandlerRepo.Error())
		return nil, errorHandlerRepo
	}

	return findAllLogRequestsRepo, nil
}

func (lu *logRequestUseCaseImpl) SaveNewLogRequest(payload dao.CreateLogRequest) (dao.CreateLogRequest, error) {
	createNewLogRequestRepo, errorHandlerRepo := lu.LogRequestRepository.Save(payload)
	if errorHandlerRepo != nil {
		util.LoggerOutput("Error when saving log", "Error", errorHandlerRepo.Error())
	}

	return createNewLogRequestRepo, nil
}
