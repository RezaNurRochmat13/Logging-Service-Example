package util

import (
	"net/http"

	"github.com/labstack/echo"
)

func ErrorResponseBadRequest(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusBadRequest, echo.Map{
		"message": message,
	})
}
