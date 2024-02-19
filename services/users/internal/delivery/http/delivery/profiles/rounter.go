package profiles

import (
	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/configs"
	pkg_middleware "github.com/p-jirayusakul/golang-echo-homework-1/pkg/middleware"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/usecases"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/config"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories/factories"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/usecases/profiles"
)

type ProfileHttpHandler struct {
	Cfg             *config.UserConfig
	ProfilesUsecase usecases.ProfilesUsecase
}

func NewProfilesHttpHandler(
	app *echo.Echo,
	config *config.UserConfig,
	dbFactory *factories.DBFactory,

) {
	handler := &ProfileHttpHandler{
		Cfg: config,
		ProfilesUsecase: profiles.NewProfilesInteractor(
			config,
			dbFactory,
		),
	}

	g := app.Group("/users")
	g.Use(pkg_middleware.ConfigJWT(configs.Config.JWT_SECRET))

	g.POST("/profiles", handler.createProfiles)
	g.GET("/profiles/:user_id", handler.findProfiles)
	g.PATCH("/profiles", handler.updateProfile)
	g.DELETE("/profiles", handler.deleteProfiles)
}
