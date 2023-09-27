package main

import "net/http"

func handlerError(responseWriter http.ResponseWriter, r *http.Request) {
	respondWithError(responseWriter, 400, "something went wrong")
}
