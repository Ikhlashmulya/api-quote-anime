package helper

import (
	"quote-anime-api/model/domain"
	"quote-anime-api/model/web"
)

func ToQuoteAnimeResponse(quoteAnime domain.QuoteAnime) web.QuoteAnimeResponse {
	return web.QuoteAnimeResponse{
		Anime: quoteAnime.Anime,
		CharacterName: quoteAnime.CharacterName,
		Quote: quoteAnime.Quote,
	}
}