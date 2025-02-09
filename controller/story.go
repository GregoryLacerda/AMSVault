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
	TokenService *service.TokenService
}

func newStoryController(cfg *config.Config, service *service.Service) *StoryController {
	return &StoryController{
		cfg:          cfg,
		StoryService: service.StoryService,
		TokenService: service.TokenService,
	}
}

func (c *StoryController) CreateStory(storyRequest request.StoryRequestViewModel) error {

	// userID := c.TokenService.GetUserIdFromToken(token)

	// userIDParsed, _ := strconv.ParseInt(userID, 10, 64)

	// if storyRequest.UserID == 0 {
	// 	storyRequest.UserID = userIDParsed
	// }

	story, err := entity.NewStory(storyRequest)
	if err != nil {
		return err
	}

	return c.StoryService.CreateStory(story)
}

func (s *StoryController) FindByName(name string) ([]response.StoryResponseViewModel, error) {
	stories, err := s.StoryService.GetStoriesByName(name)
	if err != nil {
		return nil, err
	}

	storiesViewModel := []response.StoryResponseViewModel{}
	for _, story := range stories {
		storiesViewModel = append(storiesViewModel, response.ParseStoryToResponseViewModel(story))
	}

	return storiesViewModel, nil
}

func (c *StoryController) FindByID(id int64) (response.StoryResponseViewModel, error) {
	story, err := c.StoryService.FindByID(id)
	if err != nil {
		return response.StoryResponseViewModel{}, err
	}

	return response.ParseStoryToResponseViewModel(story), nil

}

// func (c *StoryController) FindAllByUser(userID int64) ([]response.StoryResponseViewModel, error) {
// 	stories, err := c.StoryService.FindAllByUser(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var storiesResponse []response.StoryResponseViewModel
// 	for _, story := range stories {
// 		storiesResponse = append(storiesResponse, response.ParseStoryToResponseViewModel(story))
// 	}
// 	if len(storiesResponse) == 0 {
// 		return nil, errors.New("no stories found")
// 	}

// 	return storiesResponse, nil
// }

func (c *StoryController) Update(storyRequest request.StoryRequestViewModel) error {
	story, err := entity.NewStory(storyRequest)
	if err != nil {
		return err
	}

	err = c.StoryService.Update(story)
	if err != nil {
		return err
	}

	return nil
}

func (c *StoryController) Delete(id int64) error {
	return c.StoryService.Delete(id)
}
