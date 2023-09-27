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
