package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/common"
	pkg_middleware "github.com/p-jirayusakul/golang-echo-homework-1/pkg/middleware"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/utils"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/delivery/http/request"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/delivery/http/response"
)

func (h *AccountsHttpHandler) createRegister(c echo.Context) error {
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

	id, err := h.AccountsUsecase.Create(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	payload := response.RegisterResponse{
		ID: id,
	}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (h *AccountsHttpHandler) login(c echo.Context) error {
	var r request.LoginRequest

	err := c.Bind(&r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := h.AccountsUsecase.Read(r.Email)
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
		Secret:    h.Cfg.JWT_SECRET,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload response.LoginResponse
	payload.Token = token

	return utils.RespondWithJSON(c, http.StatusOK, payload)
}

func (h *AccountsHttpHandler) requestResetPassword(c echo.Context) error {
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

	id, err := h.ResetPasswordUsecase.Create(arg)
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

func (h *AccountsHttpHandler) resetPassword(c echo.Context) error {
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

	err = h.AccountsUsecase.UpdatePassword(arg)
	if err != nil {
		if errors.Is(err, common.ErrDataNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}
