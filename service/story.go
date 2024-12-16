package service

import (
	"github.com.br/GregoryLacerda/AMSVault/constants"
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

func (s *StoryService) CreateStory(story *entity.Story) error {
	if err := story.Validate(); err != nil {
		return err
	}

	err := s.data.Mongo.Insert(constants.STORY_COLLECTION, story)
	if err != nil {
		return err
	}

	modelStory := model.ConvertStoryToModel(*story)

	if err = s.data.Mysql.StoryDB.Insert(modelStory); err != nil {
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

func (s *StoryService) FindAllByUser(user string) ([]entity.Story, error) {
	stories, err := s.data.Mongo.FindAllByField(constants.STORY_COLLECTION, "user", user)
	if err != nil {
		return nil, err
	}

	return stories, nil
}

func (s *StoryService) FindByID(id string) (entity.Story, error) {
	story, err := s.data.Mongo.FindOne(constants.STORY_COLLECTION, id)
	if err != nil {
		return story, err
	}

	return story, nil
}

func (s *StoryService) Update(story *entity.Story) (entity.Story, error) {

	storyUpdated, err := s.data.Mongo.UpdateOne(constants.STORY_COLLECTION, story)
	if err != nil {
		return entity.Story{}, err
	}
	return storyUpdated, nil
}

func (s *StoryService) Delete(id string) error {
	err := s.data.Mongo.DeleteOne(constants.STORY_COLLECTION, id)
	if err != nil {
		return err
	}

	return nil
}
