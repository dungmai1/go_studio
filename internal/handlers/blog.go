package handlers

import (
	"go_studio/internal/helpers"
	"go_studio/internal/templates"
	"net/http"
)

type BlogHandLer struct{}

func NewBlogHandler() *BlogHandLer {
	return &BlogHandLer{}
}

func (h *BlogHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	seoData := helpers.NewBlogPage()
	c := templates.Blog()
	err := templates.Layout(c, "blog-page", "blog", seoData).Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
