package model

import (
	"time"

	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type Story struct {
	ID          int64  `db:"id"`
	UserID      int64  `db:"user"`
	Name        string `db:"name"`
	Source      string `db:"source"`
	Description string `db:"description"`
	Season      int64  `db:"season,omitempty"`
	Episode     int64  `db:"episode,omitempty"`
	Volume      int64  `db:"volume,omitempty"`
	Chapter     int64  `db:"chapter,omitempty"`
	Status      string `db:"status"`
	MediumImage string `db:"medium_image"`
	LargeImage  string `db:"large_image"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"update_at"`
	DeletedAt   string `db:"deleted_at"`
}

func (s Story) ToEntity() (retVal entity.Story, err error) {

	createdAt, err := time.Parse(time.DateTime, s.CreatedAt)
	if err != nil {
		return retVal, err
	}

	updatedAt, err := time.Parse(time.DateTime, s.UpdatedAt)
	if err != nil {
		return retVal, err
	}

	deletedAt, err := time.Parse(time.DateTime, s.DeletedAt)
	if err != nil {
		return retVal, err
	}

	return entity.Story{
		ID:          s.ID,
		UserID:      s.UserID,
		Name:        s.Name,
		Source:      s.Source,
		Description: s.Description,
		Season:      s.Season,
		Episode:     s.Episode,
		Chapter:     s.Chapter,
		Volume:      s.Volume,
		Status:      s.Status,
		MainPicture: entity.MainPicture{
			Medium: s.MediumImage,
			Large:  s.LargeImage,
		},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}, nil

}

func ConvertStoryToModel(s entity.Story) Story {

	return Story{
		ID:          s.ID,
		UserID:      s.UserID,
		Name:        s.Name,
		Source:      s.Source,
		Description: s.Description,
		Season:      s.Season,
		Episode:     s.Episode,
		Volume:      s.Volume,
		Chapter:     s.Chapter,
		Status:      s.Status,
		MediumImage: s.MainPicture.Medium,
		LargeImage:  s.MainPicture.Large,
		CreatedAt:   s.CreatedAt.Format(time.DateTime),
		UpdatedAt:   s.UpdatedAt.Format(time.DateTime),
		DeletedAt:   s.DeletedAt.Format(time.DateTime),
	}

}
