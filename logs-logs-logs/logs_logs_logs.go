package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		unicode := fmt.Sprintf("%U", char)

		if unicode == "U+2757" {
			return "recommendation"
		} else if unicode == "U+1F50D" {
			return "search"
		} else if unicode == "U+2600" {
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	oldRuneValue := fmt.Sprintf("%c", oldRune)
	newRuneValue := fmt.Sprintf("%c", newRune)

	return strings.ReplaceAll(log, oldRuneValue, newRuneValue)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}