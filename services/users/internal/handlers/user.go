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

	app.POST("/profiles", handler.createProfiles)
	app.POST("/register", handler.createRegister)
	app.POST("/address", handler.createAddress)
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
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Phone:     r.Phone,
	}

	err = h.profilesUsecase.Create(arg)
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
