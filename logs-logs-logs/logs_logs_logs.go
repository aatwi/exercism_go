package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	supportedRunes := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}

	for _, logChar := range log {
		val, exists := supportedRunes[logChar]
		if exists {
			return val
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	result := ""
	for _, logChar := range log {
		if logChar == oldRune {
			result += string(newRune)
		} else {
			result += string(logChar)
		}
	}
	return result
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	count := utf8.RuneCountInString(log)
	return count <= limit
}
