# Makefile for Go + Templ project
.PHONY: dev watch-templ watch-go build clean run install help fmt test vet fresh prod

# Default target
.DEFAULT_GOAL := help

# Development với auto-reload (chạy song song templ watch và air)
dev:
	@echo 🚀 Starting development server with auto-reload...
	@echo 📁 Watching .templ files and .go files...
	@echo ⚠️  Press Ctrl+C to stop
	@$(MAKE) -j2 watch-templ watch-go

# Watch templ files và auto-generate
watch-templ:
	@echo 👀 Watching .templ files...
	@templ generate --watch

# Watch Go files và restart server với air
watch-go:
	@echo 👀 Watching .go files...
	@air

# Generate templates một lần
generate:
	@echo 🔄 Generating templates...
	@templ generate
	@echo ✅ Templates generated

# Build project
build: generate
	@echo 🔨 Building project...
	@go build -o bin/main.exe ./cmd/main.go
	@echo ✅ Build completed: bin/main.exe

# Run một lần (không watch)
run: generate
	@echo ▶️ Running server...
	@go run ./cmd/main.go

# Install dependencies và tools
install:
	@echo 📦 Installing dependencies...
	@go mod tidy
	@echo 📦 Installing air (hot reload)...
	@go install github.com/cosmtrek/air@latest
	@echo 📦 Installing templ...
	@go install github.com/a-h/templ/cmd/templ@latest
	@echo ✅ All dependencies installed

# Clean build files
clean:
	@echo 🧹 Cleaning build files...
	@if exist bin rmdir /s /q bin 2>nul || echo bin directory not found
	@if exist tmp rmdir /s /q tmp 2>nul || echo tmp directory not found
	@for /r %%i in (*_templ.go) do del "%%i" 2>nul || echo No templ files to clean
	@echo ✅ Cleanup completed

# Format code
fmt: generate
	@echo 🎨 Formatting code...
	@go fmt ./...
	@templ fmt ./internal/templates/
	@echo ✅ Code formatted

# Test project
test: generate
	@echo 🧪 Running tests...
	@go test ./...

# Check for potential issues
vet: generate
	@echo 🔍 Vetting code...
	@go vet ./...

# Tidy dependencies
tidy:
	@echo 📦 Tidying dependencies...
	@go mod tidy
	@echo ✅ Dependencies tidied

# Start fresh (clean + install + dev)
fresh: clean install
	@echo 🆕 Starting fresh development environment...
	@$(MAKE) dev

# Production build
prod: clean generate
	@echo 🏭 Building for production...
	@set CGO_ENABLED=0 && go build -ldflags="-w -s" -o bin/main.exe ./cmd/main.go
	@echo ✅ Production build completed: bin/main.exe

# Quick development setup
setup: install generate
	@echo 🎯 Project setup completed!
	@echo Run 'make dev' to start development

# Check project status
status:
	@echo 📊 Project Status:
	@echo Go version:
	@go version
	@echo Templ version:
	@templ version 2>nul || echo Templ not installed
	@echo Air version:
	@air -v 2>nul || echo Air not installed
	@echo Dependencies:
	@go list -m all

# Show available endpoints
endpoints:
	@echo 🌐 Available endpoints when server is running:
	@echo Main pages:
	@echo   http://localhost:8080/         - Homepage
	@echo   http://localhost:8080/about    - About page
	@echo   http://localhost:8080/service  - Services page
	@echo   http://localhost:8080/portfolio - Portfolio page
	@echo   http://localhost:8080/blog     - Blog page
	@echo   http://localhost:8080/contact  - Contact page
	@echo   http://localhost:8080/pricing  - Pricing page
	@echo SEO endpoints:
	@echo   http://localhost:8080/sitemap.xml - Sitemap
	@echo   http://localhost:8080/robots.txt  - Robots.txt

# Show help
help:
	@echo 🎯 Lemon Digital - Go Studio Makefile
	@echo.
	@echo 📋 Available commands:
	@echo.
	@echo 🚀 Development:
	@echo   dev          Start development server with auto-reload
	@echo   run          Run the server once (no auto-reload)
	@echo   fresh        Clean + install + start development
	@echo.
	@echo 🔨 Build:
	@echo   build        Build the project
	@echo   prod         Production build (optimized)
	@echo   generate     Generate templates from .templ files
	@echo.
	@echo 📦 Dependencies:
	@echo   install      Install required dependencies
	@echo   setup        Complete project setup
	@echo   tidy         Tidy dependencies
	@echo.
	@echo 🧹 Maintenance:
	@echo   clean        Clean build files and generated templates
	@echo   fmt          Format Go and Templ code
	@echo   test         Run tests
	@echo   vet          Check code for potential issues
	@echo.
	@echo ℹ️ Information:
	@echo   status       Show project status and versions
	@echo   endpoints    Show available server endpoints
	@echo   help         Show this help message
	@echo.
	@echo 🚀 Quick start:
	@echo   make setup   - First time setup
	@echo   make dev     - Start development
	@echo.
	@echo 💡 Tips:
	@echo   - Use 'make dev' for daily development with auto-reload
	@echo   - Use 'make fresh' when you want to start clean
	@echo   - Use 'make prod' for production builds
