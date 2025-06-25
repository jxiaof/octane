package utils

import (
    "fmt"
)

// ASCIIArt generates a simple ASCII art representation of the given text.
func ASCIIArt(text string) string {
    return fmt.Sprintf(`
  ____  _        _   _             
 / __ \| |      | | (_)            
| |  | | |  ___ | |_ _ _ __   __ _ 
| |  | | | / _ \| __| | '_ \ / _` + "`" + ` |
| |__| | || (_) | |_| | | | | (_| |
 \____/|_| \___/ \__|_|_| |_|\__, |
                               __/ |
                              |___/ 
%s`, text)
}