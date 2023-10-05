package models

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

func DbModelToUser(input database.User) User {
	return User{
		ID:        input.ID,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		Name:      input.Name,
		ApiKey:    input.ApiKey,
	}
}
