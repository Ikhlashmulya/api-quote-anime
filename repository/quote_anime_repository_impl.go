package repository

import (
	"context"
	"database/sql"
	"errors"
	"quote-anime-api/helper"
	"quote-anime-api/model/domain"
)

type QuoteAnimeRepositoryImpl struct {
}

func NewQuoteAnimeRepository() QuoteAnimeRepository {
	return &QuoteAnimeRepositoryImpl{}
}

func (repository *QuoteAnimeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, quoteAnimeId int) (domain.QuoteAnime, error) {
	SQL := "select id, anime, character_name, quote from quotes_anime where id = ?"

	rows, err := tx.QueryContext(ctx, SQL, quoteAnimeId)
	helper.PanicIfError(err)
	defer rows.Close()

	quoteAnime := domain.QuoteAnime{}
	if rows.Next() {
		err := rows.Scan(&quoteAnime.Id, &quoteAnime.Anime, &quoteAnime.CharacterName, &quoteAnime.Quote)
		helper.PanicIfError(err)
		return quoteAnime, nil
	} else {
		return quoteAnime, errors.New("data not found")
	}
}

func (repository *QuoteAnimeRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, characterName string) ([]domain.QuoteAnime, error) {
	SQL := "select id, anime, character_name, quote from quotes_anime where character_name like ? limit 10"
	characterName = "%" + characterName + "%"

	rows, err := tx.QueryContext(ctx, SQL, characterName)
	helper.PanicIfError(err)
	defer rows.Close()

	quotes := []domain.QuoteAnime{}

	if rows.Next() {
		quote := domain.QuoteAnime{}
		err := rows.Scan(&quote.Id, &quote.Anime, &quote.CharacterName, &quote.Quote)
		helper.PanicIfError(err)
		quotes = append(quotes, quote)
		for rows.Next() {
			quote := domain.QuoteAnime{}
			err := rows.Scan(&quote.Id, &quote.Anime, &quote.CharacterName, &quote.Quote)
			helper.PanicIfError(err)
			quotes = append(quotes, quote)
		}
	} else {
		return quotes, errors.New("quotes not found")
	}

	return quotes, nil
}
