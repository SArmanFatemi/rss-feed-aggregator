package models

import "github.com/sarmanfatemi/rssagg/internal/database"

type ApiConfiguration struct {
	DbQueries *database.Queries
}
