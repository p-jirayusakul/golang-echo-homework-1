package address

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	pkg_middleware "github.com/p-jirayusakul/golang-echo-homework-1/pkg/middleware"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/utils"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/delivery/http/request"
)

func (h *AddressHttpHandler) createAddress(c echo.Context) error {
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

	err = h.AddressUsecase.Create(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusCreated, payload)
}

func (h *AddressHttpHandler) findAddress(c echo.Context) error {
	var r request.FindAddressRequest

	err := c.Bind(&r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := h.AddressUsecase.Read(r.AddressId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return utils.RespondWithJSON(c, http.StatusOK, result)
}

func (h *AddressHttpHandler) updateAddress(c echo.Context) error {
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

	err = h.AddressUsecase.Update(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusOK, payload)
}

func (h *AddressHttpHandler) deleteAddress(c echo.Context) error {

	var r request.DeleteAddressRequest

	err := c.Bind(&r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.AddressUsecase.Delete(r.AddressId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var payload interface{}
	return utils.RespondWithJSON(c, http.StatusOK, payload)
}
