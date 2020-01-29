package repository

import "svc-logger-go/modules/v1/log-request/dao"

type Repository interface {
	FindAll() ([]dao.ListLogRequest, error)
	Save(payload dao.CreateLogRequest) (dao.CreateLogRequest, error)
}
