package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/sarmanfatemi/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apiKey"`
}

func dbModelToUser(input database.User) User {
	return User{
		ID:        input.ID,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		Name:      input.Name,
		ApiKey:    input.ApiKey,
	}
}

type Feed struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserId    uuid.UUID `json:"userId"`
}

func dbModelToFeed(input database.Feed) Feed {
	return Feed{
		Id:        input.ID,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		Name:      input.Name,
		Url:       input.Url,
		UserId:    input.UserID,
	}
}

func dbModelsToFeeds(inputs []database.Feed) []Feed {
	feeds := []Feed{}
	for _, input := range inputs {
		feeds = append(feeds, dbModelToFeed(input))
	}

	return feeds
}
