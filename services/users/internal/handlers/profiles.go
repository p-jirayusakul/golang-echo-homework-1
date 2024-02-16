package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/utils"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/entities"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/usecases"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/handlers/request"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/usecases/profiles"
)

type ProfilesHandler struct {
	profilesUsecase usecases.ProfilesUsecase
}

func NewUserHttpHandler(
	app *echo.Echo,
	profilesRepo *repositories.ProfilesRepository,
) {
	handler := &ProfilesHandler{
		profilesUsecase: profiles.NewProfilesInteractor(
			profilesRepo,
		),
	}

	app.POST("/profiles", handler.createProfiles)
}

func (h *ProfilesHandler) createProfiles(c echo.Context) error {
	var r request.CreateProfilesReqiest

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
