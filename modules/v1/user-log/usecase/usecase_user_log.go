package usecase

import (
	"fmt"
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

func (us *userLogActivityUseCaseImpl) FindAllUserActivityLog() ([]dao.ListUserActivityLog, error) {
	findAllUserActivityLog, errorHandlerRepo := us.UserLogActivityRepository.FindAll()
	if errorHandlerRepo != nil {
		util.LoggerOutput("Error when get find all user log", "Error", errorHandlerRepo.Error())
		return nil, errorHandlerRepo
	}

	return findAllUserActivityLog, nil
}

func (us *userLogActivityUseCaseImpl) CountAllUserActivityLog() (int64, error) {
	countAllUserActivityLog, errorHandlerRepo := us.UserLogActivityRepository.Count()
	if errorHandlerRepo != nil {
		util.LoggerOutput("Error when count all user log", "Error", errorHandlerRepo.Error())
		return 0, errorHandlerRepo
	}

	return countAllUserActivityLog, nil
}

func (us *userLogActivityUseCaseImpl) FindByUserActivityByID(id string) (dao.DetailUserActivityLog, error) {
	findUserActivityByID, errorHandlerRepo := us.UserLogActivityRepository.FindByID(id)
	if errorHandlerRepo != nil {
		util.LoggerOutput("Error when find user activity by id", "Error", errorHandlerRepo.Error())
		return dao.DetailUserActivityLog{}, errorHandlerRepo
	}

	return findUserActivityByID, nil
}

func (us *userLogActivityUseCaseImpl) UpdateUserActivityLog(id string, payload *dao.UpdateUserActivityLog) (*dao.UpdateUserActivityLog, error) {
	_, errorHandlerRepoUserActivityByID := us.UserLogActivityRepository.FindByID(id)
	if errorHandlerRepoUserActivityByID != nil {
		util.LoggerOutput("Error when find by id user activity", "Error", errorHandlerRepoUserActivityByID.Error())
		errorNotFound := fmt.Errorf("User activity not found with id : %s", id)
		return nil, errorNotFound
	}

	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()
	updateUserActivity, errorHandlerRepoUpdateUserActivity := us.UserLogActivityRepository.Update(id, payload)
	if errorHandlerRepoUpdateUserActivity != nil {
		util.LoggerOutput("Error when update user activity", "Error", errorHandlerRepoUpdateUserActivity.Error())
		return nil, errorHandlerRepoUpdateUserActivity
	}

	return updateUserActivity, nil
}
