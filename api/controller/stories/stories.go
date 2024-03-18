package stories

import "amsvault/api/config"

type StoriesController struct {
	Read   *ReadController
	Update *UpdateController
	Delete *DeleteController
	Create *CreateController
}

func NewStoriesController(cfg *config.Config) StoriesController {
	controller := new(StoriesController)

}
