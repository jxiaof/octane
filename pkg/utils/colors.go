package utils

import "fmt"

// ANSI color codes
const (
	Reset     = "\033[0m"
	Bold      = "\033[1m"
	Dim       = "\033[2m"
	Underline = "\033[4m"

	// Foreground colors
	Black  = "\033[30m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"

	// Bright colors
	BrightBlack  = "\033[90m"
	BrightRed    = "\033[91m"
	BrightGreen  = "\033[92m"
	BrightYellow = "\033[93m"
	BrightBlue   = "\033[94m"
	BrightPurple = "\033[95m"
	BrightCyan   = "\033[96m"
	BrightWhite  = "\033[97m"

	// Background colors
	BgBlack  = "\033[40m"
	BgRed    = "\033[41m"
	BgGreen  = "\033[42m"
	BgYellow = "\033[43m"
	BgBlue   = "\033[44m"
	BgPurple = "\033[45m"
	BgCyan   = "\033[46m"
	BgWhite  = "\033[47m"
)

// ColorPrint prints text with specified color
func ColorPrint(text string, color string) {
	fmt.Printf("%s%s%s", color, text, Reset)
}

// ColorPrintln prints text with specified color and newline
func ColorPrintln(text string, color string) {
	fmt.Printf("%s%s%s\n", color, text, Reset)
}

// Colorize returns a string wrapped in the specified color
func Colorize(text string, color string) string {
	return fmt.Sprintf("%s%s%s", color, text, Reset)
}

// ColorizeWithStyle returns a string with color and style
func ColorizeWithStyle(text string, color string, style string) string {
	return fmt.Sprintf("%s%s%s%s", style, color, text, Reset)
}

// Header creates a styled header
func Header(text string) string {
	return ColorizeWithStyle(text, BrightCyan, Bold)
}

// Success creates success text
func Success(text string) string {
	return Colorize(text, BrightGreen)
}

// Warning creates warning text
func Warning(text string) string {
	return Colorize(text, BrightYellow)
}

// Error creates error text
func Error(text string) string {
	return Colorize(text, BrightRed)
}

// Info creates info text
func Info(text string) string {
	return Colorize(text, BrightBlue)
}

// Value creates value text (for displaying values)
func Value(text string) string {
	return Colorize(text, BrightWhite)
}

// Label creates label text (for displaying labels)
func Label(text string) string {
	return Colorize(text, Cyan)
}

// Highlight creates highlighted text
func Highlight(text string) string {
	return ColorizeWithStyle(text, BrightYellow, Bold)
}
