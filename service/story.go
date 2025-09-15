package service

import (
	"strings"

	"github.com.br/GregoryLacerda/AMSVault/data"
	"github.com.br/GregoryLacerda/AMSVault/data/model"
	"github.com.br/GregoryLacerda/AMSVault/entity"
	"github.com.br/GregoryLacerda/AMSVault/integration"
	"github.com.br/GregoryLacerda/AMSVault/pkg/errors"
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
		return errors.NewValidationError(err.Error())
	}

	modelStory := model.ToModelStory(story)

	_, err := s.data.Mysql.StoryDB.Insert(modelStory)
	if err != nil {
		return err
	}

	return nil
}

func (s *StoryService) GetStoriesByName(name string) (storys []entity.Story, err error) {

	dbStories, err := s.data.Mysql.StoryDB.FindAllByName(name)
	if err != nil {
		return nil, err
	}

	if len(dbStories) >= 10 {
		return storys, nil
	}

	malStories, err := s.Integrations.MALIntegration.GetStoriesByName(name)
	if err != nil {
		return nil, errors.NewExternalServiceError("MAL", err)
	}

	if len(dbStories) == 0 {
		return malStories, nil
	}

	for _, malStory := range malStories {
		for _, story := range dbStories {
			if malStory.Name != story.Name {
				storys = append(storys, malStory)
			}
		}
	}

	return storys, nil
}

func (s *StoryService) FindByID(id int64) (entity.Story, error) {
	story, err := s.data.Mysql.StoryDB.FindByID(id)
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return entity.Story{}, err
	}

	if story.ID == 0 {
		story, err = s.Integrations.MALIntegration.GetStoryByID(id)
		if err != nil {
			return entity.Story{}, errors.NewExternalServiceError("MAL", err)
		}

		savedStory, err := s.data.Mysql.StoryDB.Insert(model.ToModelStory(story))
		if err != nil {
			return entity.Story{}, err
		}
		return savedStory, nil
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
