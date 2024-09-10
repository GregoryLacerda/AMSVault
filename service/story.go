package service

import (
	"github.com.br/GregoryLacerda/AMSVault/data"
	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type StoryService struct {
	data *data.Data
}

func newStoryService(data *data.Data) *StoryService {
	return &StoryService{
		data: data,
	}
}

func (s *StoryService) CreateStory(story *entity.Story) error {
	if err := story.Validate(); err != nil {
		return err
	}

	err := s.data.Mongo.Insert("story", story)
	if err != nil {
		return err
	}

	return nil
}

func (s *StoryService) FindAllByUser(user string) ([]entity.Story, error) {
	stories, err := s.data.Mongo.FindAllByField("story", "user", user)
	if err != nil {
		return nil, err
	}

	return stories, nil
}

func (s *StoryService) FindByID(id string) (entity.Story, error) {
	story, err := s.data.Mongo.FindByID("story", id)
	if err != nil {
		return story, err
	}

	return story, nil
}

func (s *StoryService) DeleteStory(id string) error {
	err := s.data.Mongo.Delete("story", id)
	if err != nil {
		return err
	}

	return nil
}
