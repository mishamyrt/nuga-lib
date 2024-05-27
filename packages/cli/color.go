package cli

import "fmt"

// Attribute represents text visual attributes
type Attribute int

// Foreground text colors
const (
	fgBlack Attribute = iota + 30
	fgRed
	fgGreen
	fgYellow
	fgBlue
	fgMagenta
	fgCyan
	fgWhite
)

// Color applies text attribute
func Color(a Attribute, v string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", a, v)
}

// Red returns text in red
func Red(v string) string {
	return Color(fgRed, v)
}

// Green returns text in green
func Green(v string) string {
	return Color(fgGreen, v)
}

// Yellow returns text in yellow
func Yellow(v string) string {
	return Color(fgYellow, v)
}

// Blue returns text in blue
func Blue(v string) string {
	return Color(fgBlue, v)
}

// Purple returns text in purple
func Purple(v string) string {
	return Color(fgMagenta, v)
}

// Cyan returns text in cyan
func Cyan(v string) string {
	return Color(fgCyan, v)
}

// White returns text in white
func White(v string) string {
	return Color(fgWhite, v)
}
