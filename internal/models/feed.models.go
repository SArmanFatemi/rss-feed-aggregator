package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sarmanfatemi/rssagg/internal/database"
)

type Feed struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserId    uuid.UUID `json:"userId"`
}

func DbModelToFeed(input database.Feed) Feed {
	return Feed{
		Id:        input.ID,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		Name:      input.Name,
		Url:       input.Url,
		UserId:    input.UserID,
	}
}

func DbModelsToFeeds(inputs []database.Feed) []Feed {
	feeds := []Feed{}
	for _, input := range inputs {
		feeds = append(feeds, DbModelToFeed(input))
	}

	return feeds
}
