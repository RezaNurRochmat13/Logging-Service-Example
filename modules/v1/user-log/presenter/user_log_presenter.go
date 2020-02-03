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
