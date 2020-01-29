package usecase

import "svc-logger-go/modules/v1/log-request/dao"

type UseCase interface {
	FindAllLogRequests() ([]dao.ListLogRequest, error)
	SaveNewLogRequest(payload *dao.CreateLogRequest) (*dao.CreateLogRequest, error)
}
