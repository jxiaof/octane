package utils

import "fmt"

const (
    Reset  = "\033[0m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
    Purple = "\033[35m"
    Cyan   = "\033[36m"
    White  = "\033[37m"
)

// Colorize returns a string wrapped in the specified color.
func Colorize(text string, color string) string {
    return fmt.Sprintf("%s%s%s", color, text, Reset)
}