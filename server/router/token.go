package router

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel"
	"github.com/labstack/echo"
)

func RegisterTokenRouter(r *echo.Group, cfg *config.Config, ctrl *controller.Controller) {

	const (
		token = "token"
	)
	router := NewTokenRouters(cfg, ctrl)

	r.POST(token, router.CreateToken)
}

type TokenRouter struct {
	Ctrl *controller.Controller
	cfg  *config.Config
}

func NewTokenRouters(cfg *config.Config, ctrl *controller.Controller) *TokenRouter {
	return &TokenRouter{
		Ctrl: ctrl,
		cfg:  cfg,
	}
}

func (a *TokenRouter) CreateToken(c echo.Context) error {

	request := new(viewmodel.TokenRequestViewModel)
	c.Bind(&request)

	if request.Email == "" || request.Password == "" {
		return c.JSON(400, "invalid email or password")
	}

	tokenResponse, err := a.Ctrl.Token.CreateToken(request.Email, request.Password)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, tokenResponse)

}
