package handlers

import (
	"go_studio/internal/helpers"
	"go_studio/internal/templates"
	"net/http"
)

type IndexHandLer struct{}

func NewIndexHandler() *IndexHandLer {
	return &IndexHandLer{}
}

func (h *IndexHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	seoData := helpers.NewHomePage()
	c := templates.Index()
	err := templates.Layout(c, "index-page", "home", seoData).Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
