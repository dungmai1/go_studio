package main

import (
	"fmt"
	"go_studio/internal/handlers"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", handlers.NewIndexHandler())
	http.Handle("/blog", handlers.NewBlogHandler())
	http.Handle("/about", handlers.NewAboutHandler())
	http.Handle("/pricing", handlers.NewPricingHandLer())
	http.Handle("/service", handlers.NewServiceHandLer())
	http.Handle("/contact", handlers.NewContactHandLer())
	http.Handle("/portfolio", handlers.NewPortfolioHandLer())

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
