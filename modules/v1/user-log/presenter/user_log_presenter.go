package presenter

import (
	"net/http"
	"svc-logger-go/modules/v1/user-log/dao"
	"svc-logger-go/modules/v1/user-log/usecase"
	"svc-logger-go/util"

	"github.com/labstack/echo"
)

type userLogActivityHandlerImpl struct {
	UserLogActivityUseCase usecase.UseCase
}

func NewUserLogActivityHandler(e *echo.Echo, userLogActivityUseCase usecase.UseCase) {
	injectionHandler := &userLogActivityHandlerImpl{UserLogActivityUseCase: userLogActivityUseCase}
	groupPath := e.Group("api/v1/")
	groupPath.POST("user-log", injectionHandler.CreateNewUserLogActivityHandler)
	groupPath.GET("user-logs", injectionHandler.GetAllUserActivityLog)
	groupPath.GET("user-log/:id", injectionHandler.GetSingleUserActivityHandler)
}

func (uh *userLogActivityHandlerImpl) CreateNewUserLogActivityHandler(ctx echo.Context) error {
	userLogActivityPayload := new(dao.CreateActivityLogUser)

	errorHandlerBind := ctx.Bind(userLogActivityPayload)
	if errorHandlerBind != nil {
		util.LoggerOutput("Error when bind json", "Error", errorHandlerBind.Error())
		return util.ErrorResponseBadRequest(ctx, "Error when bind json. More info view on logs")
	}

	saveCreateNewUserLogActivity, errorHandlerUseCase := uh.UserLogActivityUseCase.CreateNewUserActivityLog(userLogActivityPayload)
	if errorHandlerUseCase != nil {
		util.LoggerOutput("Error when saving new user log", "Error", errorHandlerUseCase.Error())
		return util.ErrorResponseBadRequest(ctx, "Error when saving new user log. More info view on logs")
	}

	return util.CustomResponseMessage(ctx, http.StatusCreated, "Sukses create new user log", saveCreateNewUserLogActivity)
}

func (uh *userLogActivityHandlerImpl) GetAllUserActivityLog(ctx echo.Context) error {
	findAllUserActivityLog, errorHandlerUseCase := uh.UserLogActivityUseCase.FindAllUserActivityLog()
	if errorHandlerUseCase != nil {
		util.LoggerOutput("Error when find all user", "Error", errorHandlerUseCase.Error())
		return util.ErrorResponseBadRequest(ctx, "Error when find all user activity")
	}

	countAllUserActivityLog, errorHandlerCountUseCase := uh.UserLogActivityUseCase.CountAllUserActivityLog()
	if errorHandlerCountUseCase != nil {
		util.LoggerOutput("Error when count all user", "Error", errorHandlerCountUseCase.Error())
		return util.ErrorResponseBadRequest(ctx, "Error when count all user. More info view logs")
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"count": len(findAllUserActivityLog),
		"total": countAllUserActivityLog,
		"data":  findAllUserActivityLog,
	})
}

func (uh *userLogActivityHandlerImpl) GetSingleUserActivityHandler(ctx echo.Context) error {
	id := ctx.Param("id")

	findUserLogActivityById, errorHandlerUseCaseFindById := uh.UserLogActivityUseCase.FindByUserActivityByID(id)
	if errorHandlerUseCaseFindById != nil {
		util.LoggerOutput("Error when get find user activity by id", "Error", errorHandlerUseCaseFindById.Error())
		return util.ErrorResponseBadRequest(ctx, errorHandlerUseCaseFindById.Error())
	}

	return ctx.JSON(http.StatusOK, echo.Map{"data": findUserLogActivityById})
}
