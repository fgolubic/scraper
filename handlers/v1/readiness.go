package handlers

import (
	"net/http"
	"github.com/fgolubic/scraper/tools"
)

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	tools.RespondWithJSON(w, 200, struct{}{})
}