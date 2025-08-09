package helpers

import (
	"fmt"
	"go_studio/internal/models"
	"strings"
)

// NewHomePage creates SEO data for homepage
func NewHomePage() models.SEOData {
	return models.SEOData{
		Title:        "Lemon Digital - Agency hàng đầu tại Việt Nam | Ecommerce Enabler",
		Description:  "Chuyên về thiết kế website, marketing digital, và các giải pháp thương mại điện tử. Đối tác tin cậy cho doanh nghiệp Việt Nam.",
		Keywords:     "website design, digital marketing, ecommerce, vietnam agency, thiết kế web, marketing online",
		CanonicalURL: "https://lemondigital.store/",
		OGImage:      "https://lemondigital.store/static/assets/img/hero-carousel/hero-carousel-1.jpg",
		Schema:       GenerateOrganizationSchema(),
	}
}

// NewAboutPage creates SEO data for about page
func NewAboutPage() models.SEOData {
	return models.SEOData{
		Title:        "Giới thiệu về Lemon Digital - Agency chuyên nghiệp tại Việt Nam",
		Description:  "Tìm hiểu về Lemon Digital - đội ngũ chuyên gia hàng đầu trong lĩnh vực thiết kế web và marketing digital tại Việt Nam.",
		Keywords:     "giới thiệu, về chúng tôi, lemon digital, agency vietnam, đội ngũ chuyên gia",
		CanonicalURL: "https://lemondigital.store/about",
		OGImage:      "https://lemondigital.store/static/assets/img/about.jpg",
		Schema:       GenerateAboutPageSchema(),
	}
}

// NewServicePage creates SEO data for services page
func NewServicePage() models.SEOData {
	return models.SEOData{
		Title:        "Dịch vụ Digital Marketing & Thiết kế Website | Lemon Digital",
		Description:  "Khám phá các dịch vụ chuyên nghiệp: thiết kế website, SEO, social media marketing, và giải pháp ecommerce toàn diện.",
		Keywords:     "dịch vụ, thiết kế website, seo, social media marketing, ecommerce, digital marketing",
		CanonicalURL: "https://lemondigital.store/services",
		OGImage:      "https://lemondigital.store/static/assets/img/services.jpg",
		Schema:       GenerateServicePageSchema(),
	}
}

// NewBlogPage creates SEO data for blog listing page
func NewBlogPage() models.SEOData {
	return models.SEOData{
		Title:        "Blog - Tin tức & Kiến thức Digital Marketing | Lemon Digital",
		Description:  "Cập nhật tin tức mới nhất về digital marketing, thiết kế web, SEO và các xu hướng công nghệ trong ngành.",
		Keywords:     "blog, tin tức, digital marketing, thiết kế web, seo, xu hướng công nghệ",
		CanonicalURL: "https://lemondigital.store/blog",
		OGImage:      "https://lemondigital.store/static/assets/img/blog/blog-1.jpg",
		Schema:       GenerateBlogPageSchema(),
	}
}

// NewBlogDetailPage creates SEO data for specific blog post
func NewBlogDetailPage(post models.BlogPost) models.SEOData {
	return models.SEOData{
		Title:        post.Title + " | Lemon Digital Blog",
		Description:  post.Excerpt,
		Keywords:     strings.Join(post.Tags, ", "),
		CanonicalURL: fmt.Sprintf("https://lemondigital.store/blog/%s", post.Slug),
		OGImage:      post.FeaturedImg,
		Schema:       GenerateArticleSchema(post),
	}
}

// NewPortfolioPage creates SEO data for portfolio page
func NewPortfolioPage() models.SEOData {
	return models.SEOData{
		Title:        "Portfolio - Dự án tiêu biểu | Lemon Digital",
		Description:  "Khám phá các dự án website và marketing digital tiêu biểu mà Lemon Digital đã thực hiện cho khách hàng.",
		Keywords:     "portfolio, dự án, website, marketing digital, khách hàng, case study",
		CanonicalURL: "https://lemondigital.store/portfolio",
		OGImage:      "https://lemondigital.store/static/assets/img/masonry-portfolio/masonry-portfolio-1.jpg",
		Schema:       GeneratePortfolioPageSchema(),
	}
}

// NewContactPage creates SEO data for contact page
func NewContactPage() models.SEOData {
	return models.SEOData{
		Title:        "Liên hệ - Tư vấn miễn phí | Lemon Digital",
		Description:  "Liên hệ với Lemon Digital để được tư vấn miễn phí về thiết kế website và giải pháp marketing digital.",
		Keywords:     "liên hệ, tư vấn, miễn phí, hotline, email, địa chỉ",
		CanonicalURL: "https://lemondigital.store/contact",
		OGImage:      "https://lemondigital.store/static/assets/img/logo.png",
		Schema:       GenerateContactPageSchema(),
	}
}

// NewPricingPage creates SEO data for pricing page
func NewPricingPage() models.SEOData {
	return models.SEOData{
		Title:        "Bảng giá dịch vụ Digital Marketing | Lemon Digital",
		Description:  "Tham khảo bảng giá các dịch vụ thiết kế website, SEO, quảng cáo TikTok, Shopee và các gói tư vấn ecommerce.",
		Keywords:     "bảng giá, giá dịch vụ, thiết kế website, seo, quảng cáo, tiktok, shopee",
		CanonicalURL: "https://lemondigital.store/pricing",
		OGImage:      "https://lemondigital.store/static/assets/img/services.jpg",
		Schema:       GeneratePricingPageSchema(),
	}
}

// Schema generators
func GenerateOrganizationSchema() string {
	return `{
		"@context": "https://schema.org",
		"@type": "Organization",
		"name": "Lemon Digital",
		"description": "Agency hàng đầu tại Việt Nam - Ecommerce Enabler",
		"url": "https://lemondigital.store",
		"logo": "https://lemondigital.store/static/assets/img/logo.png",
		"contactPoint": {
			"@type": "ContactPoint",
			"telephone": "+84938787735",
			"contactType": "customer service",
			"email": "hi@lemondigital.vn"
		},
		"address": {
			"@type": "PostalAddress",
			"addressCountry": "VN",
			"addressLocality": "Ho Chi Minh City"
		},
		"sameAs": [
			"https://facebook.com/lemondigital",
			"https://linkedin.com/company/lemondigital"
		]
	}`
}

func GenerateAboutPageSchema() string {
	return `{
		"@context": "https://schema.org",
		"@type": "AboutPage",
		"name": "Giới thiệu về Lemon Digital",
		"description": "Tìm hiểu về Lemon Digital - đội ngũ chuyên gia hàng đầu trong lĩnh vực thiết kế web và marketing digital",
		"url": "https://lemondigital.store/about"
	}`
}

func GenerateServicePageSchema() string {
	return `{
		"@context": "https://schema.org",
		"@type": "Service",
		"name": "Dịch vụ Digital Marketing & Thiết kế Website",
		"description": "Các dịch vụ chuyên nghiệp về thiết kế website, SEO, social media marketing",
		"provider": {
			"@type": "Organization",
			"name": "Lemon Digital"
		}
	}`
}

func GenerateBlogPageSchema() string {
	return `{
		"@context": "https://schema.org",
		"@type": "Blog",
		"name": "Lemon Digital Blog",
		"description": "Blog về digital marketing, thiết kế web và xu hướng công nghệ",
		"url": "https://lemondigital.store/blog"
	}`
}

func GenerateArticleSchema(post models.BlogPost) string {
	return fmt.Sprintf(`{
		"@context": "https://schema.org",
		"@type": "Article",
		"headline": "%s",
		"description": "%s",
		"author": {
			"@type": "Person",
			"name": "%s"
		},
		"datePublished": "%s",
		"image": "%s",
		"publisher": {
			"@type": "Organization",
			"name": "Lemon Digital",
			"logo": "https://lemondigital.store/static/assets/img/logo.png"
		}
	}`, post.Title, post.Excerpt, post.Author, post.PublishedAt.Format("2006-01-02"), post.FeaturedImg)
}

func GeneratePortfolioPageSchema() string {
	return `{
		"@context": "https://schema.org",
		"@type": "CollectionPage",
		"name": "Portfolio Lemon Digital",
		"description": "Các dự án tiêu biểu của Lemon Digital",
		"url": "https://lemondigital.store/portfolio"
	}`
}

func GenerateContactPageSchema() string {
	return `{
		"@context": "https://schema.org",
		"@type": "ContactPage",
		"name": "Liên hệ Lemon Digital",
		"description": "Thông tin liên hệ và tư vấn miễn phí",
		"url": "https://lemondigital.store/contact"
	}`
}

func GeneratePricingPageSchema() string {
	return `{
		"@context": "https://schema.org",
		"@type": "PriceSpecification",
		"name": "Bảng giá dịch vụ Lemon Digital",
		"description": "Bảng giá các dịch vụ digital marketing và thiết kế website",
		"url": "https://lemondigital.store/pricing",
		"priceCurrency": "VND"
	}`
}
