package router

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com/labstack/echo"
)

func registerSerieRouter(r *echo.Group, cfg *config.Config, ctrl *controller.Controller) {

	const (
		serie     = "/"
		serieByID = "/:id"
	)

	router := NewSerieRouters(cfg, ctrl)

	r.GET(serie, router.GetSerie)
	r.GET(serieByID, router.GetSerieByID)
	r.POST(serie, router.CreateSerie)
	r.PUT(serieByID, router.UpdateSerie)
	r.DELETE(serieByID, router.DeleteSerie)
}

type SerieRouters struct {
	Ctrl *controller.Controller
	cfg  *config.Config
}

func NewSerieRouters(cfg *config.Config, ctrl *controller.Controller) *SerieRouters {
	return &SerieRouters{
		Ctrl: ctrl,
		cfg:  cfg,
	}
}

func (a *SerieRouters) GetSerie(c echo.Context) error {
	return nil
}

func (a *SerieRouters) GetSerieByID(c echo.Context) error {
	return nil
}

func (a *SerieRouters) CreateSerie(c echo.Context) error {
	return nil
}

func (a *SerieRouters) UpdateSerie(c echo.Context) error {
	return nil
}

func (a *SerieRouters) DeleteSerie(c echo.Context) error {
	return nil
}
