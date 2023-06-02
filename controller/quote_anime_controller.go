package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type QuoteAnimeController interface {
	FindRandom(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByName(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
