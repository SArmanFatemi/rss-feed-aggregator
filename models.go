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

type FeedFollow struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserId    uuid.UUID `json:"userId"`
	FeedId    uuid.UUID `json:"FeedId"`
}

func dbModelToFeedFollow(input database.FeedFollow) FeedFollow {
	return FeedFollow{
		Id:        input.ID,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		UserId:    input.UserID,
		FeedId:    input.FeedID,
	}
}
