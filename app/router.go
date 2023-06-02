package app

import (
	"quote-anime-api/controller"
	"quote-anime-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(quoteAnimeController controller.QuoteAnimeController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/quote/random", quoteAnimeController.FindRandom)
	router.GET("/api/quote", quoteAnimeController.FindByName)

	router.PanicHandler = exception.ErrorHandler

	return router
}
