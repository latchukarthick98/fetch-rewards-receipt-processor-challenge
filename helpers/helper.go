/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 09/26/2023
 */

package helpers

import (
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
