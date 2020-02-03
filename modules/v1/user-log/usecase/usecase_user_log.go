package usecase

import (
	"svc-logger-go/modules/v1/user-log/dao"
	"svc-logger-go/modules/v1/user-log/repository"
	"svc-logger-go/util"
	"time"
)

type userLogActivityUseCaseImpl struct {
	UserLogActivityRepository repository.Repository
}

func NewUserLogActivityUseCase(userLogActivityRepository repository.Repository) UseCase {
	return &userLogActivityUseCaseImpl{UserLogActivityRepository: userLogActivityRepository}
}

func (us *userLogActivityUseCaseImpl) CreateNewUserActivityLog(payload *dao.CreateActivityLogUser) (*dao.CreateActivityLogUser, error) {
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	createNewUserLog, errorHandlerRepo := us.UserLogActivityRepository.Save(payload)
	if errorHandlerRepo != nil {
		util.LoggerOutput("Error when saving user log to repo", "Error", errorHandlerRepo.Error())
		return nil, errorHandlerRepo
	}

	return createNewUserLog, nil
}
