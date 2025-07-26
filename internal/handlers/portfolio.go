package handlers

import (
	"go_studio/internal/templates"
	"net/http"
)

type PortfolioHandLer struct{}

func NewPortfolioHandLer() *PortfolioHandLer {
	return &PortfolioHandLer{}
}

func (h *PortfolioHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.PortFolio()
	err := templates.Layout(c, "portfolio-page","portfolio").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}