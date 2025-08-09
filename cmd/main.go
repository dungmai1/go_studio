package main

import (
	"fmt"
	"go_studio/internal/handlers"
	"net/http"
)

func main() {
	// Static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Main pages
	http.Handle("/", handlers.NewIndexHandler())
	http.Handle("/blog", handlers.NewBlogHandler())
	http.Handle("/about", handlers.NewAboutHandler())
	http.Handle("/pricing", handlers.NewPricingHandLer())
	http.Handle("/service", handlers.NewServiceHandLer())
	http.Handle("/contact", handlers.NewContactHandLer())
	http.Handle("/portfolio", handlers.NewPortfolioHandLer())

	// SEO endpoints
	http.Handle("/sitemap.xml", handlers.NewSitemapHandler())
	http.Handle("/robots.txt", handlers.NewRobotsHandler())

	fmt.Println("Listening on :8080")
	fmt.Println("SEO endpoints:")
	fmt.Println("- Sitemap: http://localhost:8080/sitemap.xml")
	fmt.Println("- Robots: http://localhost:8080/robots.txt")
	http.ListenAndServe(":8080", nil)
}
