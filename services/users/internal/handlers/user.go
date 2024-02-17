package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/utils"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/usecases"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/handlers/request"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/usecases/accounts"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/usecases/address"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/usecases/profiles"
)

type UserHandler struct {
	profilesUsecase usecases.ProfilesUsecase
	accountsUsecase usecases.AccountsUsecase
	addressUsecase  usecases.AddressUsecase
}

func NewUserHttpHandler(
	app *echo.Echo,
	profilesRepo *repositories.ProfilesRepository,
	accountsRepo *repositories.AccountsRepository,
	addresssRepo *repositories.AddressRepository,
) {
	handler := &UserHandler{
		profilesUsecase: profiles.NewProfilesInteractor(
			profilesRepo,
		),
		accountsUsecase: accounts.NewAccountsInteractor(
			accountsRepo,
			profilesRepo,
		),
		addressUsecase: address.NewAddressInteractor(
			addresssRepo,
		),
	}

	app.POST("/register", handler.createRegister)

	app.GET("/profiles/:user_id", handler.findProfiles)
	app.PATCH("/profiles", handler.updateProfile)
	app.DELETE("/profiles", handler.deleteProfiles)

	app.POST("/address", handler.createAddress)
	app.GET("/address/:address_id", handler.findAddress)
	app.PATCH("/address", handler.updateAddress)
	app.DELETE("/address/:address_id", handler.deleteAddress)
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

	uid := "3780237c-168d-417b-9a09-39e8cf68831a"

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

	uid := "3780237c-168d-417b-9a09-39e8cf68831a"

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

func (h *UserHandler) createRegister(c echo.Context) error {
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

	err = h.accountsUsecase.Create(arg)
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
	return utils.RespondWithJSON(c, http.StatusOK, payload)
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

	uid := "3780237c-168d-417b-9a09-39e8cf68831a"

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
