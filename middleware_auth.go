package main

import (
	"fmt"
	"net/http"

	"github.com/sarmanfatemi/rssagg/internal/auth"
	"github.com/sarmanfatemi/rssagg/internal/common"
	"github.com/sarmanfatemi/rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		apiKey, err := auth.GetApiKey(request.Header)
		if err != nil {
			common.RespondWithError(responseWriter, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(request.Context(), apiKey)
		if err != nil {
			common.RespondWithError(responseWriter, 404, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(responseWriter, request, user)
	}
}
