package config

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type messageFormat struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// InternalErrResp -
func InternalErrResp(c echo.Context, err error) error {
	log.Println(err)
	return c.JSON(http.StatusInternalServerError, messageFormat{
		Status:     "failed",
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Server Error",
	})
}

// BadRequestResp -
func BadRequestResp(c echo.Context, msg string) error {
	return c.JSON(http.StatusBadRequest, messageFormat{
		Status:     "failed",
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	})
}

// Data -
func Data(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, messageFormat{
		Status:     "success",
		StatusCode: http.StatusOK,
		Data:       data,
	})
}

// OKResponse -
func OKResponse(c echo.Context) error {
	return c.JSON(http.StatusOK, messageFormat{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "OK",
	})
}

// NotFoundResponse -
func NotFoundResponse(c echo.Context) error {
	return c.JSON(http.StatusNotFound, messageFormat{
		Status:     "failed",
		StatusCode: http.StatusNotFound,
		Message:    "Not Found",
	})
}

// NotAuthorizeResp -
func NotAuthorizeResp(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, messageFormat{
		Status:     "failed",
		StatusCode: http.StatusUnauthorized,
		Message:    "Unauthorized",
	})
}
