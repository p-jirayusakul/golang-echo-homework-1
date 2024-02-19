package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/domain/usecases"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/config"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/repositories/factories"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/usecases/accounts"
	"github.com/p-jirayusakul/golang-echo-homework-1/services/auth/internal/usecases/reset_password"
)

type AccountsHttpHandler struct {
	Cfg                  *config.AuthConfig
	AccountsUsecase      usecases.AccountsUsecase
	ResetPasswordUsecase usecases.ResetPasswordUsecase
}

func NewAuthHttpHandler(
	app *echo.Echo,
	config *config.AuthConfig,
	dbFactory *factories.DBFactory,
	grpcClientFactory *factories.GrpcClientFactory,

) {
	handler := &AccountsHttpHandler{
		Cfg: config,
		AccountsUsecase: accounts.NewAccountsInteractor(
			config,
			dbFactory,
			grpcClientFactory,
		),
		ResetPasswordUsecase: reset_password.NewResetPasswordInteractor(
			config,
			dbFactory,
		),
	}

	g := app.Group("/auth")
	g.POST("/register", handler.createRegister)
	g.POST("/login", handler.login)
	g.POST("/request-reset-password", handler.requestResetPassword)
	g.POST("/reset-password", handler.resetPassword)
}
