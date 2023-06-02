package service

import (
	"context"
	"quote-anime-api/model/web"
)

type QuoteAnimeService interface {
	FindByName(ctx context.Context, characterName string) []web.QuoteAnimeResponse
	FindRandom(ctx context.Context) web.QuoteAnimeResponse
}