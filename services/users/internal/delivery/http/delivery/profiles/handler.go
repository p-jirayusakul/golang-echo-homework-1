package profiles

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	pkg_middleware "github.com/p-jirayusakul/golang-echo-homework-1/pkg/middleware"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/utils"

	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/delivery/http/request"
)

func (h *ProfileHttpHandler) createProfiles(c echo.Context) error {
	var r request.CreateProfilesRequest

	err := c.Bind(&r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	arg := entities.Profiles{
		FirstName: &r.FirstName,
		LastName:  &r.LastName,
		Email:     r.Email,
		Phone:     &r.Phone,
	}

	var userId uuid.UUID
	userId.Scan(r.UserID)
	arg.UserID = userId

	err = h.ProfilesUsecase.CreateProfiles(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (h *ProfileHttpHandler) findProfiles(c echo.Context) error {
	var r request.FindProfilesByUserId

	err := c.Bind(&r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := h.ProfilesUsecase.Read(r.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return utils.RespondWithJSON(c, http.StatusOK, result)
}

func (h *ProfileHttpHandler) updateProfile(c echo.Context) error {
	var r request.UpdateProfilesRequest

	err := c.Bind(&r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	arg := entities.UpdateProfilesDTO{
		FirstName: r.FirstName,
		LastName:  r.LastName,
	}

	uid := pkg_middleware.DecodeToken(c)

	var userId uuid.UUID
	userId.Scan(uid)
	arg.UserID = userId

	err = h.ProfilesUsecase.Update(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusOK, payload)
}

func (h *ProfileHttpHandler) deleteProfiles(c echo.Context) error {

	arg := entities.DeleteProfilesDTO{}

	uid := pkg_middleware.DecodeToken(c)

	var userId uuid.UUID
	userId.Scan(uid)
	arg.UserID = userId

	err := h.ProfilesUsecase.Delete(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusOK, payload)
}
