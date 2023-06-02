package exception

import (
	"net/http"
	"quote-anime-api/helper"
	"quote-anime-api/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(ValidationError)
	if ok {
		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error,
		})
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		helper.WriteToResponseBody(writer, web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		})
		return true
	} else {
		return false
	}

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err any) {
	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
	})
}
