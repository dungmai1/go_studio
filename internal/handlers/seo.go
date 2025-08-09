package handlers

import (
	"fmt"
	"net/http"
	"time"
)

type SitemapHandler struct{}

func NewSitemapHandler() *SitemapHandler {
	return &SitemapHandler{}
}

func (h *SitemapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")

	sitemap := generateSitemap()
	w.Write([]byte(sitemap))
}

func generateSitemap() string {
	now := time.Now().Format("2006-01-02")

	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	<url>
		<loc>https://lemondigital.store/</loc>
		<lastmod>%s</lastmod>
		<changefreq>daily</changefreq>
		<priority>1.0</priority>
	</url>
	<url>
		<loc>https://lemondigital.store/about</loc>
		<lastmod>%s</lastmod>
		<changefreq>monthly</changefreq>
		<priority>0.8</priority>
	</url>
	<url>
		<loc>https://lemondigital.store/service</loc>
		<lastmod>%s</lastmod>
		<changefreq>weekly</changefreq>
		<priority>0.9</priority>
	</url>
	<url>
		<loc>https://lemondigital.store/portfolio</loc>
		<lastmod>%s</lastmod>
		<changefreq>weekly</changefreq>
		<priority>0.8</priority>
	</url>
	<url>
		<loc>https://lemondigital.store/blog</loc>
		<lastmod>%s</lastmod>
		<changefreq>daily</changefreq>
		<priority>0.9</priority>
	</url>
	<url>
		<loc>https://lemondigital.store/contact</loc>
		<lastmod>%s</lastmod>
		<changefreq>monthly</changefreq>
		<priority>0.7</priority>
	</url>
	<url>
		<loc>https://lemondigital.store/pricing</loc>
		<lastmod>%s</lastmod>
		<changefreq>weekly</changefreq>
		<priority>0.8</priority>
	</url>
</urlset>`, now, now, now, now, now, now, now)
}

type RobotsHandler struct{}

func NewRobotsHandler() *RobotsHandler {
	return &RobotsHandler{}
}

func (h *RobotsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	robots := `User-agent: *
Allow: /

# Sitemaps
Sitemap: https://lemondigital.store/sitemap.xml

# Disallow admin paths (if any)
Disallow: /admin/
Disallow: /private/
Disallow: /*.pdf$

# Allow important paths
Allow: /static/
Allow: /assets/`

	w.Write([]byte(robots))
}
