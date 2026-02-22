package main

import "unicode"

func is_palindrome(input string) bool {
	var cleaned []rune
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, unicode.ToLower(r))
		}
	}

	i, j := 0, len(cleaned)-1
	for i < j {
		if cleaned[i] != cleaned[j] {
			return false
		}
		i++
		j--
	}
	return true
}
