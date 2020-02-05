package usecase

import (
	"fmt"
	"svc-logger-go/modules/v1/log-request/dao"
	"svc-logger-go/modules/v1/log-request/repository"
	"svc-logger-go/util"
	"time"
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

func (lu *logRequestUseCaseImpl) CountAllLogRequest() (int64, error) {
	countAllLogRequestRepo, errorHandlerQuery := lu.LogRequestRepository.Count()
	if errorHandlerQuery != nil {
		util.LoggerOutput("Error when count all log request", "Error", errorHandlerQuery.Error())
		return 0, errorHandlerQuery
	}

	return countAllLogRequestRepo, nil
}

func (lu *logRequestUseCaseImpl) SaveNewLogRequest(payload *dao.CreateLogRequest) (*dao.CreateLogRequest, error) {
	// Set modified date
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	createNewLogRequestRepo, errorHandlerRepo := lu.LogRequestRepository.Save(payload)
	if errorHandlerRepo != nil {
		util.LoggerOutput("Error when saving log", "Error", errorHandlerRepo.Error())
	}

	return createNewLogRequestRepo, nil
}

func (lu *logRequestUseCaseImpl) FindByLogRequestId(id string) (dao.DetailLogRequest, error) {
	findLogRequestById, errorHandlerRepo := lu.LogRequestRepository.FindById(id)
	if errorHandlerRepo != nil {
		util.LoggerOutput("Error when find log by id", "Error", errorHandlerRepo.Error())
		return dao.DetailLogRequest{}, errorHandlerRepo
	}

	return findLogRequestById, nil
}

func (lu *logRequestUseCaseImpl) UpdateLogRequest(id string, payload *dao.UpdateLogRequest) (*dao.UpdateLogRequest, error) {
	// Find log request by id first
	_, errorHandlerRepo := lu.LogRequestRepository.FindById(id)
	if errorHandlerRepo != nil {
		util.LoggerOutput("Error when find log by id", "Error", errorHandlerRepo.Error())
		return nil, errorHandlerRepo
	}

	// Update log request when id found
	updateLogRequestRepo, errorHandlerRepos := lu.LogRequestRepository.Update(id, payload)
	if errorHandlerRepos != nil {
		util.LoggerOutput("Error when update log request", "Error", errorHandlerRepos.Error())
		return nil, errorHandlerRepos
	}

	return updateLogRequestRepo, nil
}

func (lu *logRequestUseCaseImpl) DeleteLogRequest(id string) (dao.DetailLogRequest, error) {
	// Find log request by id first
	findByLogRequestId, errorHandlerRepo := lu.LogRequestRepository.FindById(id)
	if errorHandlerRepo != nil {
		util.LoggerOutput("Error when find log by id", "Error", errorHandlerRepo.Error())
		errorFindNotFound := fmt.Errorf("ID Log Request Not found")
		return dao.DetailLogRequest{}, errorFindNotFound
	}

	// Delete log request when id found
	errorHandlerDeleteLogRequest := lu.LogRequestRepository.Delete(id)
	if errorHandlerDeleteLogRequest != nil {
		util.LoggerOutput("Error when delete log request", "Error", errorHandlerDeleteLogRequest.Error())
		return dao.DetailLogRequest{}, errorHandlerDeleteLogRequest
	}

	return findByLogRequestId, nil
}

func (lu *logRequestUseCaseImpl) FindLogRequestByRequestName(name string) (dao.DetailLogRequest, error) {
	findByRequestNameRepo, errorHandlerRepo := lu.LogRequestRepository.FindByName(name)
	if errorHandlerRepo != nil {
		util.LoggerOutput("Error when find log request name", "Error", errorHandlerRepo.Error())
		errorFindNotFound := fmt.Errorf("Name Log Request Not found")
		return dao.DetailLogRequest{}, errorFindNotFound
	}

	return findByRequestNameRepo, nil
}
