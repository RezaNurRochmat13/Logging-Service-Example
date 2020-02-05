package presenter

import (
	"net/http"
	"svc-logger-go/modules/v1/log-request/dao"
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
	groupPath.POST("log-request", injectionHandler.CreateNewLogRequestHandler)
	groupPath.GET("log-request/:id", injectionHandler.GetSingleLogRequestsHandler)
	groupPath.PUT("log-request/:id", injectionHandler.UpdateLogRequestsHandler)
	groupPath.DELETE("log-request/:id", injectionHandler.DeleteLogRequestHandler)
	groupPath.GET("log-request-search/:name", injectionHandler.GetLogRequestByRequestNameHandler)
}

func (lh *logRequestHandlerImpl) GetAllLogRequestsHandler(ctx echo.Context) error {
	findAllLogRequestUseCase, errorHandlerUseCase := lh.LogRequestUseCase.FindAllLogRequests()
	if errorHandlerUseCase != nil {
		return util.ErrorResponseBadRequest(ctx, errorHandlerUseCase.Error())
	}

	countAllLogRequest, errorHandlerUsecaseCount := lh.LogRequestUseCase.CountAllLogRequest()
	if errorHandlerUsecaseCount != nil {
		util.LoggerOutput("Error when count log request", "Error", errorHandlerUsecaseCount.Error())
		return util.ErrorResponseBadRequest(ctx, errorHandlerUseCase.Error())
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"count": len(findAllLogRequestUseCase),
		"data":  findAllLogRequestUseCase,
		"total": countAllLogRequest,
	})
}

func (lh *logRequestHandlerImpl) CreateNewLogRequestHandler(ctx echo.Context) error {
	logRequestPayload := new(dao.CreateLogRequest)

	errBindJSON := ctx.Bind(logRequestPayload)
	if errBindJSON != nil {
		util.LoggerOutput("Error when bind json", "Error", errBindJSON.Error())
		util.ErrorResponseBadRequest(ctx, "Error bind JSON")
	}

	// Saving into log request
	createLogRequest, errorHandlerUseCase := lh.LogRequestUseCase.SaveNewLogRequest(logRequestPayload)
	if errorHandlerUseCase != nil {
		util.LoggerOutput("Error when saving into usecase", "Error", errorHandlerUseCase.Error())
		util.ErrorResponseBadRequest(ctx, "Error when saving into usecase")
	}

	return util.CustomResponseMessage(ctx, http.StatusCreated, "Created Log Requests", createLogRequest)

}

func (lh *logRequestHandlerImpl) GetSingleLogRequestsHandler(ctx echo.Context) error {
	id := ctx.Param("id")

	if id == "" {
		return util.ErrorResponseBadRequest(ctx, "Missing id is required")
	}

	// Find log request by id
	findLogRequstByIdUseCase, errorHandlerUseCase := lh.LogRequestUseCase.FindByLogRequestId(id)
	if errorHandlerUseCase != nil {
		util.LoggerOutput("Error when get log by id", "Error", errorHandlerUseCase.Error())
		return util.ErrorResponseBadRequest(ctx, "Error when find log by id")
	}

	return ctx.JSON(http.StatusOK, echo.Map{"data": findLogRequstByIdUseCase})
}

func (lh *logRequestHandlerImpl) UpdateLogRequestsHandler(ctx echo.Context) error {
	id := ctx.Param("id")
	updateLogRequest := new(dao.UpdateLogRequest)
	errorHandlerBind := ctx.Bind(updateLogRequest)

	if id == "" {
		return util.ErrorResponseBadRequest(ctx, "Missing id is required")
	}

	if errorHandlerBind != nil {
		return util.ErrorResponseBadRequest(ctx, "Error bind json value")
	}

	// Update log request
	updateLogRequestUseCase, errorHandlerUseCase := lh.LogRequestUseCase.UpdateLogRequest(id, updateLogRequest)
	if errorHandlerUseCase != nil {
		util.LoggerOutput("Error when update log request", "Error", errorHandlerUseCase.Error())
		return util.ErrorResponseBadRequest(ctx, "Error when update log request")
	}

	return util.CustomResponseMessage(ctx, http.StatusOK, "Update Log Request Success", updateLogRequestUseCase)
}

func (lh *logRequestHandlerImpl) DeleteLogRequestHandler(ctx echo.Context) error {
	id := ctx.Param("id")

	if id == "" {
		return util.ErrorResponseBadRequest(ctx, "Missing id is required")
	}

	// Delete log request
	deleteLogRequestUseCase, errorHandlerUseCase := lh.LogRequestUseCase.DeleteLogRequest(id)
	if errorHandlerUseCase != nil {
		util.LoggerOutput("Error when update or find log request ", "Error", errorHandlerUseCase.Error())
		return util.ErrorResponseBadRequest(ctx, errorHandlerUseCase.Error())
	}

	return util.CustomResponseMessage(ctx, http.StatusOK, "Delete Log Request Success", deleteLogRequestUseCase)
}

func (lh *logRequestHandlerImpl) GetLogRequestByRequestNameHandler(ctx echo.Context) error {
	name := ctx.Param("name")

	// Find log request by name
	findLogRequestByName, errorHandlerUseCase := lh.LogRequestUseCase.FindLogRequestByRequestName(name)
	if errorHandlerUseCase != nil {
		util.LoggerOutput("Error when find log by name", "Error", errorHandlerUseCase.Error())
		return util.ErrorResponseBadRequest(ctx, errorHandlerUseCase.Error())
	}

	return ctx.JSON(http.StatusOK, echo.Map{"data": findLogRequestByName})
}
