package database

import (
	"errors"

	"github.com.br/GregoryLacerda/AMSVault/entity"
	"gorm.io/gorm"
)

type Story struct {
	DB *gorm.DB
}

func NewStoryDB(db *gorm.DB) *Story {
	return &Story{
		DB: db,
	}
}

func (s *Story) Create(story *entity.Story) error {

	found, err := s.FindByName(story.Name)
	if err != nil {
		return err
	}
	if found == nil {
		return errors.New("story not found in database")
	}

	return s.DB.Create(story).Error
}

func (s *Story) FindByName(name string) (story *entity.Story, err error) {

	if err = s.DB.Find("name = ?", name).First(&story).Error; err != nil {
		return nil, err
	}

	return story, nil
}
