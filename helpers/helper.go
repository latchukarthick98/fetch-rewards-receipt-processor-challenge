/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 09/26/2023
 */

package helpers

import (
	"fetch-rewards-receipt-processor-challenge/models"
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

// check if amount is a multiple of 0.25
func IsMultipleOfQuarter(amount string) bool {
	// If string is an integer
	if !strings.Contains(amount, ".") {
		return true
	}

	// Split the string into whole and fractional parts
	parts := strings.Split(amount, ".")
	fraction := parts[1]

	// If there's more than 2 decimal places, it's not a multiple of 0.25
	if len(fraction) > 2 {
		return false
	}

	// If there's only one decimal place, pad with a zero
	if len(fraction) == 1 {
		fraction += "0"
	}

	// Check if the fractional part is in the set of valid quarter multiples
	switch fraction {
	case "00", "25", "50", "75":
		return true
	default:
		return false
	}
}

// Counts the number of pairs in the list of items
func CountPairs(items []models.Item) int {
	return len(items) / 2
}

func TrimmedLength(s string) int {
	return len(strings.TrimSpace(s))
}
