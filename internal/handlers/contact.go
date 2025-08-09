package handlers

import (
	"go_studio/internal/helpers"
	"go_studio/internal/templates"
	"net/http"
)

type ContactHandLer struct{}

func NewContactHandLer() *ContactHandLer {
	return &ContactHandLer{}
}

func (h *ContactHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	seoData := helpers.NewContactPage()
	c := templates.Contact()
	err := templates.Layout(c, "contact-page", "contact", seoData).Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
