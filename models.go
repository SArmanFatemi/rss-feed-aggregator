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

func dbModelsToFeedFollows(inputs []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, input := range inputs {
		feedFollows = append(feedFollows, dbModelToFeedFollow(input))
	}

	return feedFollows
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"publishedAt"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feedId"`
}

func dbModelToPost(input database.Post) Post {
	var description *string
	if input.Description.Valid {
		description = &input.Description.String
	}

	return Post{
		ID:          input.ID,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
		Title:       input.Title,
		Description: description,
		PublishedAt: input.PublishAt,
		Url:         input.Url,
		FeedID:      input.FeedID,
	}
}

func dbModelsToPosts(inputs []database.Post) []Post {
	posts := []Post{}
	for _, input := range inputs {
		posts = append(posts, dbModelToPost(input))
	}

	return posts
}
