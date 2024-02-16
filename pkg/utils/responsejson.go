package utils

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func RespondWithError(w http.ResponseWriter, code int, errorMessage string) {

	jsonResponse := ErrorResponse{
		Message: errorMessage,
		Status:  false,
	}
	response, _ := json.Marshal(jsonResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(SuccessResponse{
		Message: "Success",
		Data:    payload,
		Status:  true,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
