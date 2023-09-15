package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	type errResponse struct {
		Error string `json:"error"`
	}

	if code > 499 {
		log.Println("Responding with 5xx error: ", message)
	}

	RespondWithJSON(w, code, errResponse{
		Error: message,
	})
}
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to marshal JSON response: ", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
