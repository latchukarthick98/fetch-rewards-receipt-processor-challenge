/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 09/27/2023
 */

package helpers

import (
	"fetch-rewards-receipt-processor-challenge/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAlphaNumeric(t *testing.T) {
	assert.Equal(t, CountAlphaNumeric("Target"), 6)
	assert.Equal(t, CountAlphaNumeric("M&M"), 2)
	assert.Equal(t, CountAlphaNumeric("M&M Corner Market"), 14)
	assert.Equal(t, CountAlphaNumeric("  M&M Corner Market "), 14)
	assert.Equal(t, CountAlphaNumeric("7 Eleven!"), 7)
	assert.Equal(t, CountAlphaNumeric("7 Eleven! "), 7)
	assert.Equal(t, CountAlphaNumeric(" "), 0)
}

func TestRoundedDollar(t *testing.T) {
	assert.False(t, IsRounded("33.45"), "Expected to return False")
	assert.True(t, IsRounded("2.00"), "Expected to return True")
	assert.True(t, IsRounded("02.00"), "Expected to return True")
	assert.True(t, IsRounded("0.00"), "Expected to return True")
	assert.False(t, IsRounded("21.12"), "Expected to return False")
}

func TestIsMultipleOfQuarter(t *testing.T) {
	testValues := []string{"9.00", "21.00", "21.25", "21.5", "21.75", "21.1", "21.12", "20.123"}
	results := []bool{true, true, true, true, true, false, false, false}
	for i, value := range testValues {
		if results[i] {
			assert.True(t, IsMultipleOfQuarter(value), "Expected to return True")
		} else {
			assert.False(t, IsMultipleOfQuarter(value), "Expected to return False")
		}
	}

}

func TestCountPairs(t *testing.T) {
	item1 := models.Item{
		ShortDescription: "Gatorade",
		Price:            "2.25",
	}
	item2 := models.Item{
		ShortDescription: "Pepsi",
		Price:            "2.50",
	}
	item3 := models.Item{
		ShortDescription: "Knorr Creamy Chicken",
		Price:            "1.26",
	}

	items1 := []models.Item{item1, item2, item3, item1, item2}
	items2 := []models.Item{item1, item2, item3, item1, item2, item3}
	items3 := []models.Item{item1}

	assert.Equal(t, CountPairs(items1), 2)
	assert.Equal(t, CountPairs(items2), 3)
	assert.Equal(t, CountPairs(items3), 0)
}

func TestTrimmedLength(t *testing.T) {
	assert.Equal(t, TrimmedLength("   Klarbrunn 12-PK 12 FL OZ  "), 24)
	assert.Equal(t, TrimmedLength("Klarbrunn 12-PK 12 FL OZ"), 24)
	assert.Equal(t, TrimmedLength("Pepsi 100 ML"), 12)
	assert.Equal(t, TrimmedLength("Pepsi 100 ML "), 12)
}

func TestIsDayOdd(t *testing.T) {
	assert.True(t, IsDayOdd("2022-01-01"), "Expected to return true")
	assert.False(t, IsDayOdd("2023-01-02"), "Expected to return false")
	assert.True(t, IsDayOdd("2023-03-03"), "Expected to return true")
	assert.True(t, IsDayOdd("2023-03-21"), "Expected to return true")
	assert.False(t, IsDayOdd("2023-01-10"), "Expected to return false")

}

func TestIsTimeInRange(t *testing.T) {
	startTime := "14:00"
	endTime := "16:00"
	assert.True(t, IsTimeInRange("14:01", startTime, endTime), "Expected to return true")
	assert.False(t, IsTimeInRange("13:01", startTime, endTime), "Expected to return false")
	assert.False(t, IsTimeInRange("03:31", startTime, endTime), "Expected to return false")
	assert.False(t, IsTimeInRange("16:01", startTime, endTime), "Expected to return false")

}

func TestIsValidDate(t *testing.T) {
	assert.True(t, IsValidDate("2022-06-26"), "Expected to return true")
	assert.False(t, IsValidDate("2023-02-31"), "Expected to return false")
	assert.False(t, IsValidDate("2023-02-35"), "Expected to return false")
	assert.False(t, IsValidDate("2023-13-31"), "Expected to return false")
}

func TestIsValidTime(t *testing.T) {
	assert.True(t, IsValidTime("13:01"), "Expected to return true")
	assert.True(t, IsValidTime("00:00"), "Expected to return true")
	assert.False(t, IsValidTime("24:04"), "Expected to return false")
	assert.False(t, IsValidTime("25:00"), "Expected to return false")
	assert.False(t, IsValidTime("00:61"), "Expected to return false")
	assert.False(t, IsValidTime("10:61"), "Expected to return false")

}

func TestIsValidDollarValue(t *testing.T) {
	assert.True(t, IsValidDollarValue("3.35"), "Expected to return true")
	assert.True(t, IsValidDollarValue("3.00"), "Expected to return true")
	assert.True(t, IsValidDollarValue("334.35"), "Expected to return true")
	assert.True(t, IsValidDollarValue("100.00"), "Expected to return true")
	assert.True(t, IsValidDollarValue("0.00"), "Expected to return true")
	assert.True(t, IsValidDollarValue("00.00"), "Expected to return true")

	assert.False(t, IsValidDollarValue(".25"), "Expected to return false")
	assert.False(t, IsValidDollarValue("1.2"), "Expected to return false")
	assert.False(t, IsValidDollarValue("3"), "Expected to return false")
	assert.False(t, IsValidDollarValue("1.254"), "Expected to return false")
}
