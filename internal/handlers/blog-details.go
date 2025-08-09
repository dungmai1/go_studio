package handlers

import (
	"go_studio/internal/templates"
	"net/http"
)

type BlogDetailsHandLer struct{}

func NewBlogDetailsHandler() *BlogDetailsHandLer {
	return &BlogDetailsHandLer{}
}

func (h *BlogDetailsHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Blog_Details()
	err := templates.Layout(c, "blog-details-page","blog_details").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}