package controller

import (
	"net/http"
	"quote-anime-api/helper"
	"quote-anime-api/model/web"
	"quote-anime-api/service"

	"github.com/julienschmidt/httprouter"
)

type QuoteAnimeControllerImpl struct {
	QuoteAnimeService service.QuoteAnimeService
}

func NewQuoteAnimeController(quoteAnimeService service.QuoteAnimeService) QuoteAnimeController {
	return &QuoteAnimeControllerImpl{QuoteAnimeService: quoteAnimeService}
}

func (controller *QuoteAnimeControllerImpl) FindRandom(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	quoteAnimeResponse := controller.QuoteAnimeService.FindRandom(request.Context())

	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   quoteAnimeResponse,
	})
}

func (controller *QuoteAnimeControllerImpl) FindByName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	query := request.URL.Query().Get("query")

	quoteAnimeResponse := controller.QuoteAnimeService.FindByName(request.Context(), query)

	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   quoteAnimeResponse,
	})
}
