package main

import (
	"net/http"
	"quote-anime-api/app"
	"quote-anime-api/controller"
	"quote-anime-api/helper"
	"quote-anime-api/repository"
	"quote-anime-api/service"
)

func main() {
	db := app.NewDB()

	quoteAnimeRepository := repository.NewQuoteAnimeRepository()
	quoteAnimeService := service.NewQuoteAnimeService(quoteAnimeRepository, db)
	quoteAnimeController := controller.NewQuoteAnimeController(quoteAnimeService)

	router := app.NewRouter(quoteAnimeController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
