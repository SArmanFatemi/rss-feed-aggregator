package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sarmanfatemi/rssagg/internal/database"
)

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

func DbModelToPost(input database.Post) Post {
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

func DbModelsToPosts(inputs []database.Post) []Post {
	posts := []Post{}
	for _, input := range inputs {
		posts = append(posts, DbModelToPost(input))
	}

	return posts
}
