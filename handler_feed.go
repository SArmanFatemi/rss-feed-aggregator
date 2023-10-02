package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sarmanfatemi/rssagg/internal/common"
	"github.com/sarmanfatemi/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(responseWriter http.ResponseWriter, request *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		common.RespondWithError(responseWriter, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(request.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		common.RespondWithError(responseWriter, 500, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}
	common.RespondWithJson(responseWriter, 201, dbModelToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(responseWriter http.ResponseWriter, request *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(request.Context())
	if err != nil {
		common.RespondWithError(responseWriter, 500, fmt.Sprintf("Couldn't fetch feeds: %v", err))
		return
	}
	common.RespondWithJson(responseWriter, 200, dbModelsToFeeds(feeds))
}
