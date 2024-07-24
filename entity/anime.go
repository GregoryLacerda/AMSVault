package entity

import (
	"errors"
	"time"

	"github.com.br/GregoryLacerda/AMSVault/constants"
	"github.com.br/GregoryLacerda/AMSVault/pkg/entity"
)

type Anime struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Season    int64     `json:"season"`
	Episode   int64     `json:"episode"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewAnime(name string, season, episode int64, status string) (*Anime, error) {
	anime := &Anime{
		ID:        entity.NewID(),
		Name:      name,
		Season:    season,
		Episode:   episode,
		Status:    status,
		CreatedAt: time.Now(),
	}

	err := anime.Validate()
	if err != nil {
		return nil, err
	}

	return anime, nil
}

func (a *Anime) Validate() error {

	if a.ID.String() == "" {
		return errors.New(constants.ERROR_ID_REQUIRED)
	}
	if a.Name == "" {
		return errors.New(constants.ERROR_NAME_REQUIRED)
	}
	if a.Season < 0 {
		return errors.New(constants.ERROR_SEASON_INVALID)
	}
	if a.Episode < 0 {
		return errors.New(constants.ERROR_EPISODE_INVALID)
	}
	if a.Status == "" {
		return errors.New(constants.ERROR_STATUS_REQUIRED)
	}

	return nil
}
