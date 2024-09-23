package handlers

import (
	"net/http"

	"github.com/brunompx/go-riverlevels/templates/pages"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(r.Context(), w)
}
