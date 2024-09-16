package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

var cache struct {
	templates map[string]*template.Template
	mutex     sync.RWMutex
}

func init() {
	// cache.templates = make(map[string]*template.Template)
	loadTemplates()
}

func loadTemplates() {
	viewsDir := "views"

	// Load layouts, pages, and partials
	layouts, err := filepath.Glob(filepath.Join(viewsDir, "layout", "*.templ"))
	if err != nil {
		log.Fatal("Error loading layouts:", err)
	}
	pages, err := filepath.Glob(filepath.Join(viewsDir, "pages", "*.templ"))
	if err != nil {
		log.Fatal("Error loading pages:", err)
	}
	partials, err := filepath.Glob(filepath.Join(viewsDir, "partials", "*.templ"))
	if err != nil {
		log.Fatal("Error loading partials:", err)
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
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			log.Fatalf("Error parsing template %s: %v", page, err)
		}

		// Cache the template with the base name (e.g., "about" or "index")
		name := filepath.Base(page)
		name = name[:len(name)-len(filepath.Ext(name))] // Remove the ".templ" extension

		log.Printf(">>> >>>>: %v\n", name)

		// Lock the cache and store the parsed template
		cache.mutex.Lock()
		cache.templates[name] = tmpl
		// cache.templates["index"] = tmpl
		// cache.templates["about"] = tmpl
		cache.mutex.Unlock()
	}
}

// Render templates
func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	cache.mutex.RLock()
	tmpl, exists := cache.templates[name]
	cache.mutex.RUnlock()

	if !exists {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	// Execute the layout template and inject page content into it
	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Function to handle theme switching based on URL parameters
func handleThemeSwitch(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("theme") != "" {
		newTheme := r.URL.Query().Get("theme")
		if newTheme == "dark" || newTheme == "light" {
			// Set the cookie for the new theme
			http.SetCookie(w, &http.Cookie{
				Name:  "theme",
				Value: newTheme,
				Path:  "/",
			})
			// Redirect to apply the new theme
			http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
			return
		}
	}
}

func getThemeFromRequest(r *http.Request) string {
	theme := "light" // TODO: Move this to a global front-end configs file, Default theme

	// Check for theme in cookies
	cookie, err := r.Cookie("theme")
	if err == nil {
		theme = cookie.Value
	}
	return theme
}

func main() {
	// Serve static files from /public/dist
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("public/dist"))))

	// Route handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleThemeSwitch(w, r) // Handle theme switching

		theme := getThemeFromRequest(r)
		// Prepare the data to pass to the template
		data := map[string]interface{}{
			"Title":       "Home Page",
			"SiteName":    "My Website",
			"Year":        2024,
			"Description": "Welcome to the homepage!",
			"ThemeClass":  theme,
		}

		// Render the template with the provided data
		renderTemplate(w, "index", data)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		handleThemeSwitch(w, r) // Handle theme switching

		theme := getThemeFromRequest(r)
		data := map[string]interface{}{
			"Title":       "About Us",
			"SiteName":    "My Website",
			"Year":        2024,
			"Description": "We value excellence and innovation.",
			"ThemeClass":  theme,
		}
		renderTemplate(w, "about", data)
	})

	// Start the server
	http.ListenAndServe(":8083", nil)
}
