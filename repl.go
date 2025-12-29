package main

import "strings"

func cleanInput(text string) []string {
	formattedText := strings.TrimSpace(strings.ToLower(text))
	return strings.Fields(formattedText)
}
