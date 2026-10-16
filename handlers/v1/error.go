package handlers

import (
	"net/http"
	"github.com/fgolubic/scraper/tools"
)

func HandleError(w http.ResponseWriter, r *http.Request) {
	tools.RespondWithError(w, 400, "Something went wrong.")
}