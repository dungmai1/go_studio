package handlers

import (
	"go_studio/internal/helpers"
	"go_studio/internal/templates"
	"net/http"
)

type PricingHandLer struct{}

func NewPricingHandLer() *PricingHandLer {
	return &PricingHandLer{}
}

func (h *PricingHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Tạo SEO data cho pricing page
	seoData := helpers.NewPricingPage()
	c := templates.Princing()
	err := templates.Layout(c, "pricing-page", "pricing", seoData).Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
