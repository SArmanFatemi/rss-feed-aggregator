package handlers

import (
	"net/http"

	"github.com/sarmanfatemi/rssagg/internal/common"
)

func HandlerError(responseWriter http.ResponseWriter, r *http.Request) {
	common.RespondWithError(responseWriter, 400, "something went wrong")
}
