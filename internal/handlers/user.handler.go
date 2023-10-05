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

func HandlerCreateUser(responseWriter http.ResponseWriter, request *http.Request, apiCfg *models.ApiConfiguration) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		common.RespondWithError(responseWriter, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	user, err := apiCfg.DbQueries.CreateUser(request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		common.RespondWithError(responseWriter, 500, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}
	common.RespondWithJson(responseWriter, 201, models.DbModelToUser(user))
}

func HandlerGetUserByApiKey(responseWriter http.ResponseWriter, request *http.Request, user database.User, apiCfg *models.ApiConfiguration) {
	common.RespondWithJson(responseWriter, 200, models.DbModelToUser(user))
}

func HandlerGetPostsForUser(responseWriter http.ResponseWriter, request *http.Request, user database.User, apiCfg *models.ApiConfiguration) {
	posts, err := apiCfg.DbQueries.GetPostsForUserId(request.Context(), database.GetPostsForUserIdParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		common.RespondWithError(responseWriter, 500, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	common.RespondWithJson(responseWriter, 200, models.DbModelsToPosts(posts))
}
