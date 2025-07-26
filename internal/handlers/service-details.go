package handlers

import (
	"go_studio/internal/templates"
	"net/http"
)

type ServiceDetailsHandLer struct{}

func NewServiceDetailsHandLer() *ServiceDetailsHandLer {
	return &ServiceDetailsHandLer{}
}

func (h *ServiceDetailsHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Service_Details()
	err := templates.Layout(c, "service-details-page","service_details").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}