package handlers

import (
	"go_studio/internal/helpers"
	"go_studio/internal/templates"
	"net/http"
)

type BlogDetailsHandLer struct{}

func NewBlogDetailsHandler() *BlogDetailsHandLer {
	return &BlogDetailsHandLer{}
}

func (h *BlogDetailsHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Get blog post data from URL slug and create dynamic SEO data
	// For now, use default blog SEO data
	seoData := helpers.NewBlogPage()
	c := templates.Blog_Details()
	err := templates.Layout(c, "blog-details-page", "blog_details", seoData).Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
