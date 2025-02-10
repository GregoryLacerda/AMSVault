package model

import (
	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type Story struct {
	ID          int64  `db:"id" bson:"story_id"`
	Name        string `db:"name" bson:"name"`
	Source      string `db:"source" bson:"source"`
	Description string `db:"description" bson:"description"`
	Season      int64  `db:"season,omitempty" bson:"season, omitempty"`
	Episode     int64  `db:"episode,omitempty" bson:"episode, omitempty"`
	Volume      int64  `db:"volume,omitempty" bson:"volume, omitempty"`
	Chapter     int64  `db:"chapter,omitempty" bson:"chapter, omitempty"`
	Status      string `db:"status" bson:"status"`
	MediumImage string `db:"medium_image" bson:"medium_image"`
	LargeImage  string `db:"large_image" bson:"large_image"`
}

func (s Story) ToEntity() (retVal entity.Story) {

	return entity.Story{
		ID:          s.ID,
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
	}
}

func ToModelStory(s entity.Story) Story {

	return Story{
		ID:          s.ID,
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
	}

}
