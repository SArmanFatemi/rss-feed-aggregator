package main

import (
	"net/http"

	"github.com/sarmanfatemi/rssagg/internal/common"
)

func handlerError(responseWriter http.ResponseWriter, r *http.Request) {
	common.RespondWithError(responseWriter, 400, "something went wrong")
}
