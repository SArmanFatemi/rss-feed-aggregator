package handlers

import (
	"fmt"
	"net/http"

	"github.com/sarmanfatemi/rssagg/internal/auth"
	"github.com/sarmanfatemi/rssagg/internal/common"
	"github.com/sarmanfatemi/rssagg/internal/database"
	"github.com/sarmanfatemi/rssagg/internal/models"
)

type handlerInput func(response http.ResponseWriter, request *http.Request, apiConfiguration *models.ApiConfiguration)
type guardedHandlerInput func(response http.ResponseWriter, request *http.Request, user database.User, apiConfiguration *models.ApiConfiguration)

func Use(handler handlerInput, apiConfiguration *models.ApiConfiguration) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		handler(responseWriter, request, apiConfiguration)
	}
}

func UseWithGuard(handler guardedHandlerInput, apiConfiguration *models.ApiConfiguration) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		apiKey, err := auth.GetApiKey(request.Header)
		if err != nil {
			common.RespondWithError(responseWriter, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiConfiguration.DbQueries.GetUserByApiKey(request.Context(), apiKey)
		if err != nil {
			common.RespondWithError(responseWriter, 404, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(responseWriter, request, user, apiConfiguration)
	}
}
