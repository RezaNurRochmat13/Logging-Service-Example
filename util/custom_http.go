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

func CustomResponseMessage(ctx echo.Context, status int, message string, data ...interface{}) error {
	return ctx.JSON(status, echo.Map{
		"message":            message,
		"created_or_updated": data,
	})
}
