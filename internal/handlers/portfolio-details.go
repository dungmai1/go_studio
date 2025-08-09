package handlers

import (
	"go_studio/internal/helpers"
	"go_studio/internal/templates"
	"net/http"
)

type PortfolioDetailsHandLer struct{}

func NewPortfolioDetailsHandLer() *PortfolioDetailsHandLer {
	return &PortfolioDetailsHandLer{}
}

func (h *PortfolioDetailsHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Get portfolio item data from URL slug and create dynamic SEO data
	// For now, use default portfolio SEO data
	seoData := helpers.NewPortfolioPage()
	c := templates.PortFolio_Details()
	err := templates.Layout(c, "portfolio-details-page", "portfolio_details", seoData).Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
