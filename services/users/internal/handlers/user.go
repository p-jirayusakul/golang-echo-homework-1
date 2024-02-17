package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/configs"
	pkg_middleware "github.com/p-jirayusakul/golang-echo-homework-1/pkg/middleware"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/utils"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/usecases"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/handlers/request"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/usecases/address"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/usecases/profiles"
)

type UserHandler struct {
	profilesUsecase usecases.ProfilesUsecase
	addressUsecase  usecases.AddressUsecase
}

func NewUserHttpHandler(
	app *echo.Echo,
	profilesRepo *repositories.ProfilesRepository,
	addresssRepo *repositories.AddressRepository,
) {
	handler := &UserHandler{
		profilesUsecase: profiles.NewProfilesInteractor(
			profilesRepo,
		),
		addressUsecase: address.NewAddressInteractor(
			addresssRepo,
		),
	}

	g := app.Group("/users")
	g.Use(pkg_middleware.ConfigJWT(configs.Config.JWT_SECRET))

	g.POST("/profiles", handler.createProfiles)
	g.GET("/profiles/:user_id", handler.findProfiles)
	g.PATCH("/profiles", handler.updateProfile)
	g.DELETE("/profiles", handler.deleteProfiles)

	g.POST("/address", handler.createAddress)
	g.GET("/address/:address_id", handler.findAddress)
	g.PATCH("/address", handler.updateAddress)
	g.DELETE("/address/:address_id", handler.deleteAddress)
}

func (h *UserHandler) createProfiles(c echo.Context) error {
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

	err = h.profilesUsecase.Create(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (h *UserHandler) findProfiles(c echo.Context) error {
	var r request.FindProfilesByUserId

	err := c.Bind(&r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := h.profilesUsecase.Read(r.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return utils.RespondWithJSON(c, http.StatusOK, result)
}

func (h *UserHandler) updateProfile(c echo.Context) error {
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

	err = h.profilesUsecase.Update(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusOK, payload)
}

func (h *UserHandler) deleteProfiles(c echo.Context) error {

	arg := entities.DeleteProfilesDTO{}

	uid := pkg_middleware.DecodeToken(c)

	var userId uuid.UUID
	userId.Scan(uid)
	arg.UserID = userId

	err := h.profilesUsecase.Delete(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusOK, payload)
}

func (h *UserHandler) createAddress(c echo.Context) error {
	var r request.CreateAddressRequest

	err := c.Bind(&r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	r.UserID = pkg_middleware.DecodeToken(c)

	var userId uuid.UUID
	userId.Scan(r.UserID)

	arg := entities.Address{
		UserID:   userId,
		AddrType: r.AddrType,
		AddrNo:   r.AddrNo,
		Street:   r.Street,
		City:     r.City,
		State:    r.State,
	}

	err = h.addressUsecase.Create(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (h *UserHandler) findAddress(c echo.Context) error {
	var r request.FindAddressRequest

	err := c.Bind(&r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := h.addressUsecase.Read(r.AddressId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return utils.RespondWithJSON(c, http.StatusOK, result)
}

func (h *UserHandler) updateAddress(c echo.Context) error {
	var r request.UpdateAddressRequest

	err := c.Bind(&r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var addressId uuid.UUID
	addressId.Scan(r.AddressId)

	arg := entities.Address{
		AddressId: addressId,
		AddrType:  r.AddrType,
		AddrNo:    r.AddrNo,
		Street:    r.Street,
		City:      r.City,
		State:     r.State,
	}

	uid := pkg_middleware.DecodeToken(c)

	var userId uuid.UUID
	userId.Scan(uid)
	arg.UserID = userId

	err = h.addressUsecase.Update(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusOK, payload)
}

func (h *UserHandler) deleteAddress(c echo.Context) error {

	var r request.DeleteAddressRequest

	err := c.Bind(&r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.addressUsecase.Delete(r.AddressId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusOK, payload)
}
