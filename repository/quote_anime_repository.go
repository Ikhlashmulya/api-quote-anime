package repository

import (
	"context"
	"database/sql"
	"quote-anime-api/model/domain"
)

type QuoteAnimeRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, quoteAnimeId int) (domain.QuoteAnime, error)
	FindByName(ctx context.Context, tx *sql.Tx, characterName string) ([]domain.QuoteAnime, error)
}
