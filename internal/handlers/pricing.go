package handlers

import (
	"go_studio/internal/templates"
	"net/http"
)

type PricingHandLer struct{}

func NewPricingHandLer() *PricingHandLer {
	return &PricingHandLer{}
}

func (h *PricingHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Princing()
	err := templates.Layout(c, "pricing-page","pricing").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}