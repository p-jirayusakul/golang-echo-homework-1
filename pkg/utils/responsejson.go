package utils

import (
	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func RespondWithError(c echo.Context, code int, message string) error {
	return c.JSON(code, ErrorResponse{Message: message, Status: "error"})
}

func RespondWithJSON(c echo.Context, code int, payload interface{}) error {
	return c.JSON(code, SuccessResponse{Message: "success", Status: "success", Data: payload})
}
