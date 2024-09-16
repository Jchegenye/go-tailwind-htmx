package handlers

import (
	"go-tailwind-htmx/internal/contextkeys"
	"go-tailwind-htmx/internal/templates"
	"net/http"
)

// HandleHideBanner handles hiding the banner
func HandleHideBanner(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Set a cookie to hide the banner
		http.SetCookie(w, &http.Cookie{
			Name:  "hideBanner",
			Value: "true",
			Path:  "/",
		})
		// Redirect back to the referring page
		http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
	}
}

// HandleIndex handles the index page
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// Fetch context values
	hideBanner := contextkeys.GetHideBanner(r.Context())
	theme := contextkeys.GetTheme(r.Context())

	data := map[string]interface{}{
		"Title":       "Home",
		"SiteName":    "Buybot",
		"Year":        2024,
		"Description": "Welcome to Buybot!",
		"ThemeClass":  theme,
		"HideBanner":  hideBanner,
	}

	// Render the template with the provided data
	templates.RenderTemplate(w, "index", data)
}

// HandleAbout handles the about page
func HandleAbout(w http.ResponseWriter, r *http.Request) {
	// Fetch context values
	hideBanner := contextkeys.GetHideBanner(r.Context())
	theme := contextkeys.GetTheme(r.Context())

	data := map[string]interface{}{
		"Title":       "About Us",
		"SiteName":    "Buybot",
		"Year":        2024,
		"Description": "We value excellence and innovation.",
		"ThemeClass":  theme,
		"HideBanner":  hideBanner,
	}

	templates.RenderTemplate(w, "about", data)
}
