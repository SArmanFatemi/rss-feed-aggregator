package handlers

import (
	"net/http"

	"github.com/sarmanfatemi/rssagg/internal/common"
)

func HandlerReadiness(responseWriter http.ResponseWriter, r *http.Request) {
	common.RespondWithJson(responseWriter, 200, struct{}{})
}
