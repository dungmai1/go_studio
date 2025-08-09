package handlers

import (
	"go_studio/internal/templates"
	"net/http"
)

type PortfolioDetailsHandLer struct{}

func NewPortfolioDetailsHandLer() *PortfolioDetailsHandLer {
	return &PortfolioDetailsHandLer{}
}

func (h *PortfolioDetailsHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.PortFolio_Details()
	err := templates.Layout(c, "portfolio-details-page","portfolio_details").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}