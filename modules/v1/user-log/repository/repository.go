package repository

import "svc-logger-go/modules/v1/user-log/dao"

type Repository interface {
	Save(payload *dao.CreateActivityLogUser) (*dao.CreateActivityLogUser, error)
	FindAll() ([]dao.ListUserActivityLog, error)
	Count() (int64, error)
}
