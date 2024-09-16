package middleware

import (
	"context"
	"go-tailwind-htmx/internal/contextkeys"
	"net/http"
)

func GlobalMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle theme switching and banner visibility
		theme := "light" // Default theme
		if cookie, err := r.Cookie("theme"); err == nil {
			theme = cookie.Value
		}
		if r.URL.Query().Get("theme") != "" {
			newTheme := r.URL.Query().Get("theme")
			if newTheme == "dark" || newTheme == "light" {
				theme = newTheme
				http.SetCookie(w, &http.Cookie{
					Name:  "theme",
					Value: newTheme,
					Path:  "/",
				})
				// Redirect to remove theme parameter from URL
				http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
				return
			}
		}

		hideBanner := false
		if _, err := r.Cookie("hideBanner"); err == nil {
			hideBanner = true
		}

		// Set context values
		ctx := r.Context()
		ctx = context.WithValue(ctx, contextkeys.HideBannerKey, hideBanner)
		ctx = context.WithValue(ctx, contextkeys.ThemeKey, theme)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
