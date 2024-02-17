package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/common"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/configs"
	pkg_middleware "github.com/p-jirayusakul/golang-echo-homework-1/pkg/middleware"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/utils"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/usecases"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/handlers/request"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/handlers/response"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/repositories"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/usecases/accounts"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/usecases/reset_password"
)

type AuthHandler struct {
	accountsUsecase      usecases.AccountsUsecase
	resetPasswordUsecase usecases.ResetPasswordUsecase
}

func NewAuthHttpHandler(
	app *echo.Echo,
	accountsRepo *repositories.AccountsRepository,
	resetPasswordRepo *repositories.ResetPasswordRepository,
) {
	handler := &AuthHandler{
		accountsUsecase: accounts.NewAccountsInteractor(
			accountsRepo,
			resetPasswordRepo,
		),
		resetPasswordUsecase: reset_password.NewResetPasswordInteractor(
			resetPasswordRepo,
			accountsRepo,
		),
	}

	g := app.Group("/auth")
	g.POST("/register", handler.createRegister)
	g.POST("/login", handler.login)
	g.POST("/request-reset-password", handler.requestResetPassword)
	g.POST("/reset-password", handler.resetPassword)
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

	payload := response.RegisterResponse{
		ID: id,
	}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
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

	token, err := pkg_middleware.CreateToken(pkg_middleware.CreateTokenDTO{
		UserID:    result.UserID.String(),
		ExpiresAt: time.Now().Add(time.Hour * 2),
		Secret:    configs.Config.JWT_SECRET,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload response.LoginResponse
	payload.Token = token

	return utils.RespondWithJSON(c, http.StatusOK, payload)
}

func (h *AuthHandler) requestResetPassword(c echo.Context) error {
	var r request.RequestResetPasswordRequest

	err := c.Bind(&r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	arg := entities.ResetPassword{
		Email: r.Email,
	}

	id, err := h.resetPasswordUsecase.Create(arg)
	if err != nil {
		if errors.Is(err, common.ErrDataNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	payload := response.ResetPasswordResponse{
		RequestID: id,
	}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (h *AuthHandler) resetPassword(c echo.Context) error {
	var r request.ResetPasswordRequest

	err := c.Bind(&r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	arg := entities.UpdatePasswordAccountDTO{
		RequestID: r.RequestID,
		Password:  r.Password,
	}

	err = h.accountsUsecase.UpdatePassword(arg)
	if err != nil {
		if errors.Is(err, common.ErrDataNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}
