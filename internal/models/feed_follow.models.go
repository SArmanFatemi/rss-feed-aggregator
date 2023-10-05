package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sarmanfatemi/rssagg/internal/database"
)

type FeedFollow struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserId    uuid.UUID `json:"userId"`
	FeedId    uuid.UUID `json:"FeedId"`
}

func DbModelToFeedFollow(input database.FeedFollow) FeedFollow {
	return FeedFollow{
		Id:        input.ID,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		UserId:    input.UserID,
		FeedId:    input.FeedID,
	}
}

func DbModelsToFeedFollows(inputs []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, input := range inputs {
		feedFollows = append(feedFollows, DbModelToFeedFollow(input))
	}

	return feedFollows
}
