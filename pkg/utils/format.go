package utils

import (
    "fmt"
    "strings"
)

// FormatDuration formats a duration in seconds to a human-readable string.
func FormatDuration(seconds int) string {
    if seconds < 60 {
        return fmt.Sprintf("%d seconds", seconds)
    } else if seconds < 3600 {
        minutes := seconds / 60
        return fmt.Sprintf("%d minutes", minutes)
    } else {
        hours := seconds / 3600
        return fmt.Sprintf("%d hours", hours)
    }
}

// FormatScore formats a score to a string with a percentage sign.
func FormatScore(score float64) string {
    return fmt.Sprintf("%.2f%%", score)
}

// FormatList formats a list of strings into a comma-separated string.
func FormatList(items []string) string {
    return strings.Join(items, ", ")
}

// FormatKeyValue formats a key-value pair into a string.
func FormatKeyValue(key string, value interface{}) string {
    return fmt.Sprintf("%s: %v", key, value)
}