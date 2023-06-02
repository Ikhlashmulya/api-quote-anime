package helper

import (
	"encoding/json"
	"net/http"
	"quote-anime-api/model/web"
)

func WriteToResponseBody(writer http.ResponseWriter, response web.WebResponse) {
	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(response)
}
