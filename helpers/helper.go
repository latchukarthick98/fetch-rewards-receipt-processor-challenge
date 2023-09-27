/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 09/26/2023
 */

package helpers

import (
	"strings"
	"unicode"
)

// countAlphaNumeric returns the number of alphanumeric characters in a string.
func CountAlphaNumeric(s string) int {
	count := 0
	for _, runeValue := range s {
		if unicode.IsLetter(runeValue) || unicode.IsDigit(runeValue) {
			count++
		}
	}
	return count
}

// Checks if the given dollar amount is rounded to whole dollars.
func IsRounded(amount string) bool {
	parts := strings.Split(amount, ".")
	if len(parts) != 2 {
		// Not a valid dollar amount format
		return false
	}
	// Check if the decimal part is "00"
	return parts[1] == "00"
}
