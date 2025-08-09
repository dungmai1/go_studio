package handlers

import (
	"go_studio/internal/helpers"
	"go_studio/internal/templates"
	"net/http"
)

type ServiceDetailsHandLer struct{}

func NewServiceDetailsHandLer() *ServiceDetailsHandLer {
	return &ServiceDetailsHandLer{}
}

func (h *ServiceDetailsHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Get service data from URL slug and create dynamic SEO data
	// For now, use default service SEO data
	seoData := helpers.NewServicePage()
	c := templates.Service_Details()
	err := templates.Layout(c, "service-details-page", "service_details", seoData).Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
