package middleware

import (
	"net/http"
	"quote-anime-api/helper"
	"quote-anime-api/model/web"
)

type AuthMiddleware struct {
	http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-Key") == "Rahasia" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		})
	}
}
