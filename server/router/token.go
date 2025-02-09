package router

import (
	"net/http"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel"
	"github.com/labstack/echo"
)

func registerTokenRouter(r *echo.Group, cfg *config.Config, ctrl *controller.Controller) {

	const (
		token = "login"
	)
	router := newTokenRouters(cfg, ctrl)

	r.POST(token, router.CreateToken)
}

type TokenRouter struct {
	Ctrl *controller.Controller
	cfg  *config.Config
}

func newTokenRouters(cfg *config.Config, ctrl *controller.Controller) *TokenRouter {
	return &TokenRouter{
		Ctrl: ctrl,
		cfg:  cfg,
	}
}

func (a *TokenRouter) CreateToken(c echo.Context) error {

	request := new(viewmodel.TokenRequestViewModel)
	c.Bind(&request)

	if request.Email == "" || request.Password == "" {
		return c.JSON(http.StatusBadRequest, "invalid email or password")
	}

	tokenResponse, err := a.Ctrl.TokenController.CreateToken(request.Email, request.Password)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if tokenResponse.Token == "" {
		return c.JSON(http.StatusBadRequest, "can't create token")
	}

	return c.JSON(http.StatusOK, tokenResponse)

}
