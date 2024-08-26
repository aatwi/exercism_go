package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	var message = "Welcome to the Tech Palace, "
	return message + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var line = strings.Repeat("*", numStarsPerLine)
	return line + "\n" + welcomeMsg + "\n" + line
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var noStars string = strings.ReplaceAll(oldMsg, "*", "")
	return strings.Trim(strings.Trim(noStars, "\n"), " ")
}
