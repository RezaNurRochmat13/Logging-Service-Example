package usecase

import "svc-logger-go/modules/v1/user-log/dao"

type UseCase interface {
	CreateNewUserActivityLog(payload *dao.CreateActivityLogUser) (*dao.CreateActivityLogUser, error)
	FindAllUserActivityLog() ([]dao.ListUserActivityLog, error)
	CountAllUserActivityLog() (int64, error)
}
