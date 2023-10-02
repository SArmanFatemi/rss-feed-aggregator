package main

import (
	"net/http"

	"github.com/sarmanfatemi/rssagg/internal/common"
)

func handlerReadiness(responseWriter http.ResponseWriter, r *http.Request) {
	common.RespondWithJson(responseWriter, 200, struct{}{})
}
