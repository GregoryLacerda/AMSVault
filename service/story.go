package service

import (
	"github.com.br/GregoryLacerda/AMSVault/data"
	"github.com.br/GregoryLacerda/AMSVault/data/model"
	"github.com.br/GregoryLacerda/AMSVault/entity"
	"github.com.br/GregoryLacerda/AMSVault/integration"
)

type StoryService struct {
	data         *data.Data
	Integrations *integration.Integrations
}

func newStoryService(data *data.Data, Integrations *integration.Integrations) *StoryService {
	return &StoryService{
		data:         data,
		Integrations: Integrations,
	}
}

func (s *StoryService) CreateStory(story entity.Story) error {
	if err := story.Validate(); err != nil {
		return err
	}

	modelStory := model.ToModelStory(story)

	if err := s.data.Mysql.StoryDB.Insert(modelStory); err != nil {
		return err
	}

	return nil
}

func (s *StoryService) GetStoriesByName(name string) (storys []entity.Story, err error) {

	stories, err := s.Integrations.MALIntegration.GetStoriesByName(name)
	if err != nil {
		return nil, err
	}

	return stories, nil
}

func (s *StoryService) FindByID(id int64) (entity.Story, error) {
	story, err := s.Integrations.MALIntegration.GetStoryByID(id)
	if err != nil {
		return entity.Story{}, err
	}

	return story, nil
}

func (s *StoryService) Update(story entity.Story) error {

	err := s.data.Mysql.StoryDB.Update(story)
	if err != nil {
		return err
	}

	return nil
}

func (s *StoryService) Delete(id int64) error {
	err := s.data.Mysql.StoryDB.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
