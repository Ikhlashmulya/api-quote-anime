package test

import (
	"net/http"
	"net/http/httptest"
	"quote-anime-api/app"
	"quote-anime-api/controller"
	"quote-anime-api/middleware"
	"quote-anime-api/repository"
	"quote-anime-api/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupRouter() http.Handler {
	db := app.NewDB()

	quoteAnimeRepository := repository.NewQuoteAnimeRepository()
	quoteAnimeService := service.NewQuoteAnimeService(quoteAnimeRepository, db)
	quoteAnimeController := controller.NewQuoteAnimeController(quoteAnimeService)

	router := app.NewRouter(quoteAnimeController)

	return middleware.NewAuthMiddleware(router)
}

func TestFindRandom(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/quote/random", nil)
	request.Header.Set("X-API-Key", "Rahasia")
	recorder := httptest.NewRecorder()

	router := setupRouter()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestFindByNameSuccess(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/quote?query=naruto", nil)
	request.Header.Set("X-API-Key", "Rahasia")
	recorder := httptest.NewRecorder()

	router := setupRouter()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestFindByNameNotFound(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/quote?query=tidakada", nil)
	request.Header.Set("X-API-Key", "Rahasia")
	recorder := httptest.NewRecorder()

	router := setupRouter()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestFindByNameBadRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/quote", nil)
	request.Header.Set("X-API-Key", "Rahasia")
	recorder := httptest.NewRecorder()

	router := setupRouter()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}

func TestUnauthorized(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/quote/random", nil)
	recorder := httptest.NewRecorder()

	router := setupRouter()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)
}

