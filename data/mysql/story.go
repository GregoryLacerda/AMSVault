package mysql

import (
	"database/sql"

	"github.com.br/GregoryLacerda/AMSVault/data/model"
	"github.com.br/GregoryLacerda/AMSVault/entity"
	"github.com.br/GregoryLacerda/AMSVault/pkg/errors"
)

type StoryDB struct {
	DB *sql.DB
}

func newStoryDB(db *sql.DB) *StoryDB {
	return &StoryDB{
		DB: db,
	}
}

func (s *StoryDB) Insert(story model.Story) (entity.Story, error) {
	result, err := s.FindByName(story.Name)
	if err != nil {
		if err != sql.ErrNoRows {
			return entity.Story{}, errors.NewDatabaseError("FindByName", err)
		}
	}
	if result.ID != 0 {
		return entity.Story{}, errors.NewInternalError("story already exists", nil)
	}

	query := `INSERT INTO stories (name, mal_id, source, description, season, episode, volume, chapter, status, medium_image, large_image) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := s.DB.Exec(query,
		story.Name,
		story.MALID,
		story.Source,
		story.Description,
		story.TotalSeason,
		story.TotalEpisode,
		story.TotalVolume,
		story.TotalChapter,
		story.Status,
		story.MediumImage,
		story.LargeImage,
	)
	if err != nil {
		return entity.Story{}, errors.NewDatabaseError("Insert", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return entity.Story{}, errors.NewDatabaseError("Insert", err)
	}

	return s.FindByID(id)
}

func (s *StoryDB) FindByID(ID int64) (entity.Story, error) {
	var story entity.Story
	query := "SELECT * FROM stories WHERE id = ?"
	err := s.DB.QueryRow(query, ID).Scan(
		&story.ID,
		&story.Name,
		&story.MALID,
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
		return story, errors.NewDatabaseError("FindByID", err)
	}
	return story, nil
}

func (s *StoryDB) FindByName(name string) (entity.Story, error) {
	var story entity.Story
	query := "SELECT * FROM stories WHERE name LIKE ?"
	err := s.DB.QueryRow(query, "%"+name+"%").Scan(
		&story.ID,
		&story.Name,
		&story.MALID,
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
		return entity.Story{}, errors.NewDatabaseError("FindByName", err)
	}

	return story, nil
}

func (s *StoryDB) FindAllByName(name string) ([]entity.Story, error) {
	var stories []entity.Story
	query := "SELECT * FROM stories WHERE name LIKE ?"
	rows, err := s.DB.Query(query, "%"+name+"%")
	if err != nil {
		return nil, errors.NewDatabaseError("FindAllByName", err)
	}
	defer rows.Close()

	for rows.Next() {
		var story entity.Story
		err := rows.Scan(
			&story.ID,
			&story.Name,
			&story.MALID,
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
			return nil, errors.NewDatabaseError("FindAllByName", err)
		}
		stories = append(stories, story)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.NewDatabaseError("FindAllByName", err)
	}

	return stories, nil
}

func (s *StoryDB) Update(story entity.Story) error {
	query := `
	UPDATE stories
	SET name = ?, mal_id = ?, source = ?, description = ?, total_season = ?, total_episode = ?, total_volume = ?, total_chapter = ?, status = ?, medium_image = ?, large_image = ?
	WHERE id = ?
	`

	storyModel := model.ToModelStory(story)

	if _, err := s.DB.Exec(query,
		storyModel.Name,
		storyModel.MALID,
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
		return errors.NewDatabaseError("Update", err)
	}

	return nil
}

func (s *StoryDB) Delete(ID int64) error {

	query := "DELETE FROM stories WHERE id = ?"
	if _, err := s.DB.Exec(query, ID); err != nil {
		return errors.NewDatabaseError("Delete", err)
	}

	return nil
}
