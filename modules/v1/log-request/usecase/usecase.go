package usecase

import "svc-logger-go/modules/v1/log-request/dao"

type UseCase interface {
	FindAllLogRequests() ([]dao.ListLogRequest, error)
	FindByLogRequestId(id string) (dao.DetailLogRequest, error)
	CountAllLogRequest() (int64, error)
	SaveNewLogRequest(payload *dao.CreateLogRequest) (*dao.CreateLogRequest, error)
	UpdateLogRequest(id string, payload *dao.UpdateLogRequest) (*dao.UpdateLogRequest, error)
	DeleteLogRequest(id string) (dao.DetailLogRequest, error)
}
