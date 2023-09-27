/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 09/26/2023
 */

package helpers

import (
	"fetch-rewards-receipt-processor-challenge/models"
	"fmt"
	"strconv"
	"strings"
	"time"
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

// Calculates the length of the string after trimming whitespaces
func TrimmedLength(s string) int {
	return len(strings.TrimSpace(s))
}

// Determines if the day is odd in a date passed as string
func IsDayOdd(date string) bool {
	parts := strings.Split(date, "-")
	day, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	return int(day)%2 != 0
}

// Check if a time string is in the given range. Time is passed as string.
func IsTimeInRange(tStr string, startTimeStr string, endTimeStr string) bool {
	layout := "15:04" // 24-hour time format

	t, err := time.Parse(layout, tStr)
	if err != nil {
		fmt.Println(err)
		return false
	}

	startTime, _ := time.Parse(layout, startTimeStr)
	endTime, _ := time.Parse(layout, endTimeStr)
	fmt.Printf("Time in range: %t\n", t.After(startTime) && t.Before(endTime))
	return t.After(startTime) && t.Before(endTime)
}
