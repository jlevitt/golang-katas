package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	parsedLine, err := parseLine("10-14 k: vqgnlzbbkjkljk")

	if assert.NoError(t, err) {
		assert.Equal(t, 10, parsedLine.min)
		assert.Equal(t, 14, parsedLine.max)
		assert.Equal(t, "k", parsedLine.char)
		assert.Equal(t, "vqgnlzbbkjkljk", parsedLine.password)
	}
}
