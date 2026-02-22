package main

import (
	"regexp"
	"strings"
)

func word_frequency_count(input string) map[string]int {
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	processed := re.ReplaceAllString(input, " ")
	processed = strings.ToLower(processed)
	words := strings.Split(processed, " ")

	counts := make(map[string]int)

	for _, word := range words {
		if word != "" {
			counts[word]++
		}
	}
	return counts
}
