package repository

import "svc-logger-go/modules/v1/log-request/dao"

type Repository interface {
	FindAll() ([]dao.ListLogRequest, error)
	FindById(id string) (dao.DetailLogRequest, error)
	Count() (int64, error)
	Save(payload *dao.CreateLogRequest) (*dao.CreateLogRequest, error)
	Update(id string, payload *dao.UpdateLogRequest) (*dao.UpdateLogRequest, error)
	Delete(id string) error
	FindByName(name string) (dao.DetailLogRequest, error)
}
