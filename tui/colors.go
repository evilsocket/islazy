package tui

// https://misc.flogisoft.com/bash/tip_colors_and_formatting
var (
	// effects
	bold  = "\033[1m"
	dim   = "\033[2m"
	reset = "\033[0m"
	// colors
	red    = "\033[31m"
	green  = "\033[32m"
	blue   = "\033[34m"
	yellow = "\033[33m"
)

// w for wrap
func w(e, s string) string {
	return e + s + reset
}

// Bold makes the string bold.
func Bold(s string) string {
	return w(bold, s)
}

// Dim makes the string diminished.
func Dim(s string) string {
	return w(dim, s)
}

// Red makes the string red.
func Red(s string) string {
	return w(red, s)
}

// Green makes the string green.
func Green(s string) string {
	return w(green, s)
}

// Blue makes the string green.
func Blue(s string) string {
	return w(blue, s)
}

// Yellow makes the string green.
func Yellow(s string) string {
	return w(yellow, s)
}
