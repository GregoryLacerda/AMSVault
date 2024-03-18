package controller

import (
	"amsvault/api/config"
	"amsvault/api/controller/stories"
)

type Controller struct {
	Stroys *stories.StoriesController
}

func New(cfg *config.Config) (*Controller, error) {

	controller := new(Controller)

	controller.Stroys = stories.New(cfg)

}
