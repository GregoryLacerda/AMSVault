package controller

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel/request"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel/response"
	"github.com.br/GregoryLacerda/AMSVault/entity"
	"github.com.br/GregoryLacerda/AMSVault/service"
)

type StoryController struct {
	cfg          *config.Config
	StoryService *service.StoryService
}

func newStoryController(cfg *config.Config, service *service.Service) *StoryController {
	return &StoryController{
		cfg:          cfg,
		StoryService: service.StoryService,
	}
}

func (c *StoryController) CreateStory(storyRequest *request.StoryRequestViewModel) error {

	story, err := entity.NewStory(*storyRequest)
	if err != nil {
		return err
	}

	return c.StoryService.CreateStory(story)
}

func (c *StoryController) FindByID(id string) (response.StoryResponseViewModel, error) {
	story, err := c.StoryService.FindByID(id)
	if err != nil {
		return response.StoryResponseViewModel{}, err
	}

	return response.ParseStoryToResponseViewModel(story), nil

}

func (c *StoryController) FindAllByUser(user string) ([]response.StoryResponseViewModel, error) {
	stories, err := c.StoryService.FindAllByUser(user)
	if err != nil {
		return nil, err
	}

	var storiesResponse []response.StoryResponseViewModel
	for _, story := range stories {
		storiesResponse = append(storiesResponse, response.ParseStoryToResponseViewModel(story))
	}

	return storiesResponse, nil
}
