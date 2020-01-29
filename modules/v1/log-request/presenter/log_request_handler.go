package presenter

import (
	"net/http"
	"svc-logger-go/modules/v1/log-request/usecase"
	"svc-logger-go/util"

	"github.com/labstack/echo"
)

type logRequestHandlerImpl struct {
	LogRequestUseCase usecase.UseCase
}

func NewLogRequestHandler(e *echo.Echo, logRequestUseCase usecase.UseCase) {
	injectionHandler := &logRequestHandlerImpl{
		LogRequestUseCase: logRequestUseCase,
	}
	groupPath := e.Group("api/v1/")
	groupPath.GET("log-requests", injectionHandler.GetAllLogRequestsHandler)
}

func (lh *logRequestHandlerImpl) GetAllLogRequestsHandler(ctx echo.Context) error {
	findAllLogRequestUseCase, errorHandlerUseCase := lh.LogRequestUseCase.FindAllLogRequests()
	if errorHandlerUseCase != nil {
		return util.ErrorResponseBadRequest(ctx, errorHandlerUseCase.Error())
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"count": len(findAllLogRequestUseCase),
		"data":  findAllLogRequestUseCase,
	})
}
