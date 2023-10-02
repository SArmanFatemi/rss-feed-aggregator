package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sarmanfatemi/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(responseWriter http.ResponseWriter, request *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(responseWriter, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	user, err := apiCfg.DB.CreateUser(request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(responseWriter, 500, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}
	respondWithJson(responseWriter, 201, dbModelToUser(user))
}

func (apiCfg *apiConfig) handlerGetUserByApiKey(responseWriter http.ResponseWriter, request *http.Request, user database.User) {
	respondWithJson(responseWriter, 200, dbModelToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(responseWriter http.ResponseWriter, request *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUserId(request.Context(), database.GetPostsForUserIdParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(responseWriter, 500, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	respondWithJson(responseWriter, 200, dbModelsToPosts(posts))
}
