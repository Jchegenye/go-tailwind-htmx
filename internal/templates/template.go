package templates

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Template cache
var cache struct {
	templates map[string]*template.Template
}

// RenderTemplate renders the specified template with data
func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, exists := cache.templates[name]
	if !exists {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	// Render the template with the provided data
	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// LoadTemplates parses and caches all templates
func LoadTemplates() {
	viewsDir := "views"

	// Load layouts
	layouts, err := filepath.Glob(filepath.Join(viewsDir, "layout", "*.templ"))
	if err != nil {
		log.Fatal("Error loading layouts:", err)
	}

	// Load pages
	pages, err := filepath.Glob(filepath.Join(viewsDir, "pages", "*.templ"))
	if err != nil {
		log.Fatal("Error loading pages:", err)
	}

	// Load partials from known directories
	partials := []string{}
	partialDirs := []string{
		filepath.Join(viewsDir, "partials"),
		filepath.Join(viewsDir, "partials", "widgets"), // Add more partial directories if necessary here
	}

	for _, dir := range partialDirs {
		files, err := filepath.Glob(filepath.Join(dir, "*.templ"))
		if err != nil {
			log.Fatal("Error loading partials:", err)
		}
		partials = append(partials, files...)
	}

	cache.templates = make(map[string]*template.Template)

	// Include the index file as a "page"
	indexFile := filepath.Join(viewsDir, "index.templ")
	pages = append(pages, indexFile)

	// Parse and cache templates
	for _, page := range pages {
		// Combine layouts, partials, and the current page file
		files := append(layouts, partials...)
		files = append(files, page)

		// Parse all the combined files
		tmpl, err := template.New("").Funcs(template.FuncMap{}).ParseFiles(files...)
		if err != nil {
			log.Fatalf("Error parsing template %s: %v", page, err)
		}

		// Cache the template with the base name (e.g., "about" or "index")
		name := filepath.Base(page)
		name = name[:len(name)-len(filepath.Ext(name))] // Remove the ".templ" extension

		log.Printf("Template parsed: %s\n", name)

		// Lock the cache and store the parsed template
		cache.templates[name] = tmpl
	}
}
