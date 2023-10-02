package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/sarmanfatemi/rssagg/internal/common"
	"github.com/sarmanfatemi/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(responseWriter http.ResponseWriter, request *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feedId"`
	}
	decoder := json.NewDecoder(request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		common.RespondWithError(responseWriter, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollow(request.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		common.RespondWithError(responseWriter, 500, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}
	common.RespondWithJson(responseWriter, 201, dbModelToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(responseWriter http.ResponseWriter, request *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(request.Context(), user.ID)
	if err != nil {
		common.RespondWithError(responseWriter, 500, fmt.Sprintf("Couldn't fetch feed follows: %v", err))
		return
	}

	common.RespondWithJson(responseWriter, 200, dbModelsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(responseWriter http.ResponseWriter, request *http.Request, user database.User) {
	feedFollowIDString := chi.URLParam(request, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDString)
	if err != nil {
		common.RespondWithError(responseWriter, 400, fmt.Sprintf("Couldn't parse feed follow id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(request.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		common.RespondWithError(responseWriter, 500, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}

	common.RespondWithJson(responseWriter, 200, struct{}{})
}
