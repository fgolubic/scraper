package tools

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {

	dat, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to marshal json response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(status)
	w.Write(dat)

}

func RespondWithError(w http.ResponseWriter, code int, msg string) {

	if code >= 500 {
		log.Println("Responding with 5XX error:", msg)

	}

	type errResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(w, code, errResponse{
		Error: msg,
	})
}
