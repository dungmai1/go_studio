package handlers

import (
	"go_studio/internal/templates"
	"net/http"
)

type ServiceHandLer struct{}

func NewServiceHandLer() *ServiceHandLer {
	return &ServiceHandLer{}
}

func (h *ServiceHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Service()
	err := templates.Layout(c, "services-page","service").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}