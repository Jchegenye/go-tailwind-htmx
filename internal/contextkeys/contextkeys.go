package contextkeys

import "context"

type contextKey string

const (
	HideBannerKey contextKey = "hideBanner"
	ThemeKey      contextKey = "theme"
)

func GetHideBanner(ctx context.Context) bool {
	if v, ok := ctx.Value(HideBannerKey).(bool); ok {
		return v
	}
	return false
}

func GetTheme(ctx context.Context) string {
	if v, ok := ctx.Value(ThemeKey).(string); ok {
		return v
	}
	return "light"
}
