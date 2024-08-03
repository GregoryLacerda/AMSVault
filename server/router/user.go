package router

import (
	"net/http"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel"
	"github.com/labstack/echo"
)

func RegisterUserRouter(r *echo.Group, cfg *config.Config, ctrl *controller.Controller) {

	const (
		user = "user"
	)
	router := NewUserRouters(cfg, ctrl)

	r.POST(user, router.CreateUser)
	r.GET(user, router.FindByEmail)
	r.DELETE(user+"/:id", router.Delete)
}

type UserRouter struct {
	Ctrl *controller.Controller
	cfg  *config.Config
}

func NewUserRouters(cfg *config.Config, ctrl *controller.Controller) *UserRouter {
	return &UserRouter{
		Ctrl: ctrl,
		cfg:  cfg,
	}
}

func (u *UserRouter) CreateUser(c echo.Context) error {

	user := new(viewmodel.UserRequestViewModel)
	c.Bind(&user)

	if err := u.Ctrl.UserController.CreateUser(user.Name, user.Email, user.Password); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "")
}

func (u *UserRouter) FindByEmail(c echo.Context) error {

	user := new(viewmodel.UserRequestViewModel)
	c.Bind(&user)

	userResponse, err := u.Ctrl.UserController.FindByEmail(user.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, userResponse)
}

func (u *UserRouter) Delete(c echo.Context) error {

	id := c.Param("id")

	if err := u.Ctrl.UserController.Delete(id); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "Deleted")
}
