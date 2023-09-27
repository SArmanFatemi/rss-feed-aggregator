package main

import "net/http"

func handlerReadiness(responseWriter http.ResponseWriter, r *http.Request) {
	respondWithJson(responseWriter, 200, struct{}{})
}
