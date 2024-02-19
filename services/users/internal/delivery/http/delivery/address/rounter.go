package address

import (
	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-1/pkg/configs"
	pkg_middleware "github.com/p-jirayusakul/golang-echo-homework-1/pkg/middleware"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/domain/usecases"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/config"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/repositories/factories"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/users/internal/usecases/address"
)

type AddressHttpHandler struct {
	Cfg            *config.UserConfig
	AddressUsecase usecases.AddressUsecase
}

func NewAddressHttpHandler(
	app *echo.Echo,
	config *config.UserConfig,
	dbFactory *factories.DBFactory,

) {
	handler := &AddressHttpHandler{
		Cfg: config,
		AddressUsecase: address.NewAddressInteractor(
			config,
			dbFactory,
		),
	}

	g := app.Group("/users")
	g.Use(pkg_middleware.ConfigJWT(configs.Config.JWT_SECRET))

	g.POST("/address", handler.createAddress)
	g.GET("/address/:address_id", handler.findAddress)
	g.PATCH("/address", handler.updateAddress)
	g.DELETE("/address/:address_id", handler.deleteAddress)
}
