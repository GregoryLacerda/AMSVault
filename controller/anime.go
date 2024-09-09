package controller

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel"
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

func (c *StoryController) CreateStory(storyViewModel *viewmodel.StoryRequestViewModel) error {

	story, err := entity.NewStory(storyViewModel.Name, storyViewModel.Season, storyViewModel.Episode, storyViewModel.Status, storyViewModel.User)
	if err != nil {
		return err
	}

	return c.StoryService.CreateStory(story)
}

func (c *StoryController) FindByID(id string) (viewmodel.StoryResponseViewModel, error) {
	story, err := c.StoryService.FindByID(id)
	if err != nil {
		return viewmodel.StoryResponseViewModel{}, err
	}

	return viewmodel.ParseStoryToResponseViewModel(story), nil

}

func (c *StoryController) FindAllByUser(user string) ([]viewmodel.StoryResponseViewModel, error) {
	stories, err := c.StoryService.FindAllByUser(user)
	if err != nil {
		return nil, err
	}

	var storiesResponse []viewmodel.StoryResponseViewModel
	for _, story := range stories {
		storiesResponse = append(storiesResponse, viewmodel.ParseStoryToResponseViewModel(story))
	}

	return storiesResponse, nil
}
