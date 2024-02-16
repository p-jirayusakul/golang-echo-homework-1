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
	var body request.CreateProfilesReqiest

	err := c.Bind(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	arg := entities.Profiles{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Phone:     body.Phone,
	}

	err = h.profilesUsecase.Create(arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, utils.ErrorResponse{
		Status:  true,
		Message: "success",
	})
}
