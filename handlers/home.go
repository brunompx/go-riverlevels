package handlers

import (
	"net/http"

	"github.com/brunompx/go-riverlevels/views"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}
