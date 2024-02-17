package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/common"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/utils"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/usecases"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/handlers/request"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/handlers/response"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/repositories"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/usecases/accounts"
)

type AuthHandler struct {
	accountsUsecase usecases.AccountsUsecase
}

func NewAuthHttpHandler(
	app *echo.Echo,
	accountsRepo *repositories.AccountsRepository,
) {
	handler := &AuthHandler{
		accountsUsecase: accounts.NewAccountsInteractor(
			accountsRepo,
		),
	}

	g := app.Group("/auth")
	g.POST("/register", handler.createRegister)
	g.POST("/login", handler.login)
}

func (h *AuthHandler) createRegister(c echo.Context) error {
	var r request.RegisterRequest

	err := c.Bind(&r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	arg := entities.Accounts{
		Email:    r.Email,
		Password: r.Password,
	}

	id, err := h.accountsUsecase.Create(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload response.RegisterResponse
	payload.UserID = id
	return utils.RespondWithJSON(c, http.StatusOK, payload)
}

func (h *AuthHandler) login(c echo.Context) error {
	var r request.LoginRequest

	err := c.Bind(&r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := h.accountsUsecase.Read(r.Email)
	if err != nil {
		if errors.Is(err, common.ErrDataNotFound) {
			return echo.NewHTTPError(http.StatusUnauthorized, common.ErrLoginFail.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = utils.CheckPassword(r.Password, result.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, common.ErrLoginFail.Error())
	}

	var payload response.LoginResponse
	return utils.RespondWithJSON(c, http.StatusOK, payload)
}
