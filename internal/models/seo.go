package models

import "time"

// SEOData contains metadata for SEO optimization
type SEOData struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Keywords     string `json:"keywords"`
	CanonicalURL string `json:"canonical_url"`
	OGImage      string `json:"og_image"`
	Schema       string `json:"schema"`
}

// BlogPost represents a blog post with SEO data
type BlogPost struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Content     string    `json:"content"`
	Excerpt     string    `json:"excerpt"`
	FeaturedImg string    `json:"featured_img"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
	Tags        []string  `json:"tags"`
}

// Service represents a service with SEO data
type Service struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Slug        string   `json:"slug"`
	Description string   `json:"description"`
	Icon        string   `json:"icon"`
	Price       string   `json:"price"`
	Features    []string `json:"features"`
}

// PortfolioItem represents a portfolio item
type PortfolioItem struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Slug        string   `json:"slug"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Gallery     []string `json:"gallery"`
	Category    string   `json:"category"`
	Client      string   `json:"client"`
	ProjectURL  string   `json:"project_url"`
}
