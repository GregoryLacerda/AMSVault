package mysql

import (
	"database/sql"
	"errors"

	"github.com.br/GregoryLacerda/AMSVault/data/model"
	"github.com.br/GregoryLacerda/AMSVault/entity"
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

	query := `INSERT INTO storys (name, source, description, total_season, total_episode, total_volume, total_chapter, status, medium_image, large_image) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	if _, err := s.DB.Exec(query,
		story.Name,
		story.Source,
		story.Description,
		story.TotalSeason,
		story.TotalEpisode,
		story.TotalVolume,
		story.TotalChapter,
		story.Status,
		story.MediumImage,
		story.LargeImage,
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
		&story.Name,
		&story.Source,
		&story.Description,
		&story.TotalSeason,
		&story.TotalEpisode,
		&story.TotalVolume,
		&story.TotalChapter,
		&story.Status,
		&story.MainPicture.Medium,
		&story.MainPicture.Large,
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
		&story.Name,
		&story.Source,
		&story.Description,
		&story.TotalSeason,
		&story.TotalEpisode,
		&story.TotalVolume,
		&story.TotalChapter,
		&story.Status,
		&story.MainPicture.Medium,
		&story.MainPicture.Large,
	)
	if err != nil {
		return entity.Story{}, err
	}

	return story, nil
}

func (s *StoryDB) Update(story entity.Story) error {

	query := `
	UPDATE storys
	SET name = ?, source = ?, total_description = ?, total_season = ?, total_episode = ?, total_volume = ?, chapter = ?, status = ?, medium_image = ?, large_image = ?, 
	WHERE id = ?
	`

	storyModel := model.ToModelStory(story)

	if _, err := s.DB.Exec(query,
		storyModel.Name,
		storyModel.Source,
		storyModel.Description,
		storyModel.TotalSeason,
		storyModel.TotalEpisode,
		storyModel.TotalVolume,
		storyModel.TotalChapter,
		storyModel.Status,
		storyModel.MediumImage,
		storyModel.LargeImage,
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
