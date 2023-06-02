package service

import (
	"context"
	"database/sql"
	"math/rand"
	"quote-anime-api/exception"
	"quote-anime-api/helper"
	"quote-anime-api/model/web"
	"quote-anime-api/repository"
)

type QuoteAnimeServiceImpl struct {
	QuoteAnimeRepository repository.QuoteAnimeRepository
	DB                   *sql.DB
}

func NewQuoteAnimeService(quoteAnimeRepository repository.QuoteAnimeRepository, db *sql.DB) QuoteAnimeService {
	return &QuoteAnimeServiceImpl{QuoteAnimeRepository: quoteAnimeRepository, DB: db}
}

func (service *QuoteAnimeServiceImpl) FindByName(ctx context.Context, characterName string) []web.QuoteAnimeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer tx.Commit()

	if characterName == "" {
		panic(exception.NewValidationError("query parameters needed ('query=....')"))
	}

	quotesAnime, err := service.QuoteAnimeRepository.FindByName(ctx, tx, characterName)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	var quoteAnimeResponses []web.QuoteAnimeResponse
	for _, quote := range quotesAnime {
		quoteAnimeResponses = append(quoteAnimeResponses, helper.ToQuoteAnimeResponse(quote))
	}
	return quoteAnimeResponses
}

func (service *QuoteAnimeServiceImpl) FindRandom(ctx context.Context) web.QuoteAnimeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer tx.Commit()

	id := rand.Intn(8000)

	quoteAnime, err := service.QuoteAnimeRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToQuoteAnimeResponse(quoteAnime)
}
