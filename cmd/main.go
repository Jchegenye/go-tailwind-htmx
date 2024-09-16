package main

import (
	"go-tailwind-htmx/internal/handlers"
	"go-tailwind-htmx/internal/middleware"
	"go-tailwind-htmx/internal/templates"
	"log"
	"net/http"
)

func main() {
	templates.LoadTemplates()

	// Serve static files
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("public/dist"))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	// http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("public/icons"))))

	// Route handlers with middleware
	http.Handle("/", middleware.GlobalMiddleware(http.HandlerFunc(handlers.HandleIndex)))
	http.Handle("/about", middleware.GlobalMiddleware(http.HandlerFunc(handlers.HandleAbout)))
	http.Handle("/hide-banner", middleware.GlobalMiddleware(http.HandlerFunc(handlers.HandleHideBanner))) //Hidden from browser URL

	// Start the server
	log.Println("Server starting on :8083")
	http.ListenAndServe(":8083", nil)
}
