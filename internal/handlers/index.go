package handlers

import (
	"go_studio/internal/templates"
	"net/http"
)

type IndexHandLer struct{}

func NewIndexHandler() *IndexHandLer {
	return &IndexHandLer{}
}

func (h *IndexHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Index()
	err := templates.Layout(c, "index-page","home").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}