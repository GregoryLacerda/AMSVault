package model

import (
	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type Story struct {
	ID           int64  `db:"id" bson:"id"`
	Name         string `db:"name" bson:"name"`
	MALID        int64  `db:"mal_id" bson:"mal_id"`
	Source       string `db:"source" bson:"source"`
	Description  string `db:"description" bson:"description"`
	TotalSeason  int64  `db:"total_season,omitempty" bson:"total_season, omitempty"`
	TotalEpisode int64  `db:"total_episode,omitempty" bson:"total_episode, omitempty"`
	TotalVolume  int64  `db:"total_volume,omitempty" bson:"total_volume, omitempty"`
	TotalChapter int64  `db:"total_chapter,omitempty" bson:"total_chapter, omitempty"`
	Status       string `db:"status" bson:"status"`
	MediumImage  string `db:"medium_image" bson:"medium_image"`
	LargeImage   string `db:"large_image" bson:"large_image"`
}

func (s Story) ToEntity() (retVal entity.Story) {

	return entity.Story{
		ID:           s.ID,
		Name:         s.Name,
		MALID:        s.MALID,
		Source:       s.Source,
		Description:  s.Description,
		TotalSeason:  s.TotalSeason,
		TotalEpisode: s.TotalEpisode,
		TotalChapter: s.TotalChapter,
		TotalVolume:  s.TotalVolume,
		Status:       s.Status,
		MainPicture: entity.MainPicture{
			Medium: s.MediumImage,
			Large:  s.LargeImage,
		},
	}
}

func ToModelStory(s entity.Story) Story {

	return Story{
		ID:           s.ID,
		Name:         s.Name,
		MALID:        s.MALID,
		Source:       s.Source,
		Description:  s.Description,
		TotalSeason:  s.TotalSeason,
		TotalEpisode: s.TotalEpisode,
		TotalVolume:  s.TotalVolume,
		TotalChapter: s.TotalChapter,
		Status:       s.Status,
		MediumImage:  s.MainPicture.Medium,
		LargeImage:   s.MainPicture.Large,
	}

}
