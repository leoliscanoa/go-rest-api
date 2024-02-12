package helper

import (
	"encoding/json"
	"net/http"
	"rest-api/dto"
	"time"
)

func JsonResponseSuccess(writer http.ResponseWriter, data any, code int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	response := dto.ResponseDto[any]{
		Message: "Success",
		Time:    time.Now().UTC().String(),
		Data:    data,
	}
	json.NewEncoder(writer).Encode(response)
}

func JsonResponseError(writer http.ResponseWriter, error string, code int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	response := dto.ResponseDto[string]{
		Message: "Error",
		Time:    time.Now().UTC().String(),
		Error:   error,
	}
	json.NewEncoder(writer).Encode(response)
}
