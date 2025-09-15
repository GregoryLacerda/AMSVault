package router

import (
	"net/http"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel"
	"github.com.br/GregoryLacerda/AMSVault/pkg/errors"
	"github.com/labstack/echo/v4"
)

func registerLoginRouter(r *echo.Group, cfg *config.Config, ctrl *controller.Controller) {

	const (
		token = "login"
	)
	router := newLoginRouters(cfg, ctrl)

	r.POST(token, router.CreateLogin)
}

type LoginRouter struct {
	Ctrl *controller.Controller
	cfg  *config.Config
}

func newLoginRouters(cfg *config.Config, ctrl *controller.Controller) *LoginRouter {
	return &LoginRouter{
		Ctrl: ctrl,
		cfg:  cfg,
	}
}

func (a *LoginRouter) CreateLogin(c echo.Context) error {

	request := new(viewmodel.LoginRequestViewModel)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewValidationError(err.Error()))
	}

	if request.Email == "" || request.Password == "" {
		return c.JSON(http.StatusBadRequest, errors.NewValidationError("invalid email or password"))
	}

	tokenResponse, err := a.Ctrl.TokenController.CreateToken(request.Email, request.Password)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	if tokenResponse.Token == "" {
		return c.JSON(http.StatusBadRequest, errors.New("invalid email or password"))
	}

	return c.JSON(http.StatusOK, tokenResponse)

}
