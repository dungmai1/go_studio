package handlers

import (
	"go_studio/internal/helpers"
	"go_studio/internal/templates"
	"net/http"
)

type ServiceHandLer struct{}

func NewServiceHandLer() *ServiceHandLer {
	return &ServiceHandLer{}
}

func (h *ServiceHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	seoData := helpers.NewServicePage()
	c := templates.Service()
	err := templates.Layout(c, "services-page", "service", seoData).Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
