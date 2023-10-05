package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sarmanfatemi/rssagg/internal/common"
	"github.com/sarmanfatemi/rssagg/internal/database"
	"github.com/sarmanfatemi/rssagg/internal/models"
)

func HandlerCreateFeed(responseWriter http.ResponseWriter, request *http.Request, user database.User, apiCfg *models.ApiConfiguration) {
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
	feed, err := apiCfg.DbQueries.CreateFeed(request.Context(), database.CreateFeedParams{
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
	common.RespondWithJson(responseWriter, 201, models.DbModelToFeed(feed))
}

func HandlerGetFeeds(responseWriter http.ResponseWriter, request *http.Request, apiCfg *models.ApiConfiguration) {
	feeds, err := apiCfg.DbQueries.GetFeeds(request.Context())
	if err != nil {
		common.RespondWithError(responseWriter, 500, fmt.Sprintf("Couldn't fetch feeds: %v", err))
		return
	}
	common.RespondWithJson(responseWriter, 200, models.DbModelsToFeeds(feeds))
}
