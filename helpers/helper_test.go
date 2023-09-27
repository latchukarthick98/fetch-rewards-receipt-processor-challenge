/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 09/27/2023
 */

package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAlphaNumeric(t *testing.T) {
	assert.Equal(t, CountAlphaNumeric("Target"), 6)
	assert.Equal(t, CountAlphaNumeric("M&M"), 2)
	assert.Equal(t, CountAlphaNumeric("M&M Corner Market"), 14)
	assert.Equal(t, CountAlphaNumeric("7 Eleven!"), 7)
}

func TestRoundedDollar(t *testing.T) {
	assert.False(t, IsRounded("33.45"), "Expected to return False")
	assert.True(t, IsRounded("2.00"), "Expected to return True")
	assert.True(t, IsRounded("0.00"), "Expected to return True")
	assert.False(t, IsRounded("21.12"), "Expected to return False")
}
