package router

import (
	"net/http"
	"strconv"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel"
	"github.com.br/GregoryLacerda/AMSVault/entity"
	"github.com.br/GregoryLacerda/AMSVault/pkg/errors"
	"github.com.br/GregoryLacerda/AMSVault/server/middleware"
	"github.com/labstack/echo/v4"
)

func registerUserRouter(r *echo.Group, cfg *config.Config, ctrl *controller.Controller) {

	const (
		user = "user"
	)
	router := NewUserRouters(cfg, ctrl)

	r.POST(user, router.CreateUser)

	r.Use(middleware.JWTMiddleware(cfg))
	r.GET(user, router.FindByEmail)
	r.GET(user+"/:id", router.FindById)
	r.DELETE(user+"/:id", router.Delete)
	r.PUT(user, router.Update)
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
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewValidationError(err.Error()))
	}

	if err := u.Ctrl.UserController.CreateUser(user.Name, user.Email, user.Password); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})
}

func (u *UserRouter) FindByEmail(c echo.Context) error {

	user := new(viewmodel.UserRequestViewModel)
	c.Bind(&user)

	userResponse, err := u.Ctrl.UserController.FindByEmail(user.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, userResponse)
}

func (u *UserRouter) Delete(c echo.Context) error {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewInternalError("DeleteUser", err))
	}

	if err := u.Ctrl.UserController.Delete(id); err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

func (u *UserRouter) Update(c echo.Context) error {

	user := new(viewmodel.UserRequestViewModel)
	c.Bind(&user)

	userToUpdate, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	userToUpdate.ID = user.ID

	if err := u.Ctrl.UserController.Update(userToUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User updated successfully"})
}

func (u *UserRouter) FindById(c echo.Context) error {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewInternalError("FindById", err))
	}

	userResponse, err := u.Ctrl.UserController.FindById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, userResponse)
}
