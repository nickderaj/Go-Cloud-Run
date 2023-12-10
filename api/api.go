package api

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ApiError struct {
	Error string `json:"error"`
}

func HandleError(c echo.Context, msg string) error {
	err := c.JSON(http.StatusBadRequest, ApiError{
		Error: msg,
	})
	if err != nil {
		return err
	}
	return errors.New(msg) // pass error forward
}
