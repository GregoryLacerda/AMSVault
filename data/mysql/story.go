package mysql

import (
	"database/sql"
	"errors"
	"time"

	"github.com.br/GregoryLacerda/AMSVault/data/model"
	"github.com.br/GregoryLacerda/AMSVault/entity"
	"github.com.br/GregoryLacerda/AMSVault/utils"
)

type StoryDB struct {
	DB *sql.DB
}

func newStoryDB(db *sql.DB) *StoryDB {
	return &StoryDB{
		DB: db,
	}
}

func (s *StoryDB) Insert(story model.Story) error {

	result, err := s.SelectByName(story.Name)
	if err != nil {
		return err
	}
	if result.ID != 0 {
		return errors.New("story already exists")
	}

	query := `INSERT INTO storys (user, name, source, description, season, episode, volume, chapter, status, medium_image, large_image, created_at, updated_at, deleted_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	if _, err := s.DB.Exec(query,
		story.User,
		story.Name,
		story.Source,
		story.Description,
		story.Season,
		story.Episode,
		story.Volume,
		story.Chapter,
		story.Status,
		story.MediumImage,
		story.LargeImage,
		story.CreatedAt,
		story.UpdatedAt,
		story.DeletedAt,
	); err != nil {
		return err
	}

	return nil

}

func (s *StoryDB) SelectByID(ID int64) (entity.Story, error) {
	var story entity.Story
	query := "SELECT * FROM storys WHERE id = ?"
	err := s.DB.QueryRow(query, ID).Scan(
		&story.ID,
		&story.User,
		&story.Name,
		&story.Source,
		&story.Description,
		&story.Season,
		&story.Episode,
		&story.Volume,
		&story.Chapter,
		&story.Status,
		&story.MainPicture.Medium,
		&story.MainPicture.Large,
		&story.CreatedAt,
		&story.UpdatedAt,
		&story.DeletedAt,
	)
	if err != nil {
		return story, err
	}
	return story, nil
}

func (s *StoryDB) SelectByName(name string) (entity.Story, error) {
	var story entity.Story
	query := "SELECT * FROM storys WHERE name = ?"
	err := s.DB.QueryRow(query, name).Scan(
		&story.ID,
		&story.User,
		&story.Name,
		&story.Source,
		&story.Description,
		&story.Season,
		&story.Episode,
		&story.Volume,
		&story.Chapter,
		&story.Status,
		&story.MainPicture.Medium,
		&story.MainPicture.Large,
		&story.CreatedAt,
		&story.UpdatedAt,
		&story.DeletedAt,
	)
	if err != nil {
		return entity.Story{}, err
	}

	return story, nil
}

func (s *StoryDB) Update(story entity.Story) error {

	query := `
	UPDATE storys
	SET user = ?, name = ?, source = ?, description = ?, season = ?, episode = ?, volume = ?, chapter = ?, status = ?, medium_image = ?, large_image = ?, created_at = ?, updated_at = ?, deleted_at = ?
	WHERE id = ?
	`

	updatedat := time.Now().In(utils.GetDefaultLocation()).Format(time.DateTime)
	storyModel := model.ConvertStoryToModel(story)

	if _, err := s.DB.Exec(query,
		storyModel.User,
		storyModel.Name,
		storyModel.Source,
		storyModel.Description,
		storyModel.Season,
		storyModel.Episode,
		storyModel.Volume,
		storyModel.Chapter,
		storyModel.Status,
		storyModel.MediumImage,
		storyModel.LargeImage,
		storyModel.CreatedAt,
		updatedat,
		storyModel.DeletedAt,
		storyModel.ID,
	); err != nil {
		return err
	}

	return nil
}

func (s *StoryDB) Delete(ID int64) error {

	query := "DELETE FROM storys WHERE id = ?"
	if _, err := s.DB.Exec(query, ID); err != nil {
		return err
	}

	return nil
}
