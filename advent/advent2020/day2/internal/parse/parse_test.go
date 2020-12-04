package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	parsedLine, err := NewLine("10-14 k: vqgnlzbbkjkljk")

	if assert.NoError(t, err) {
		assert.Equal(t, 10, parsedLine.A)
		assert.Equal(t, 14, parsedLine.B)
		assert.Equal(t, "k", parsedLine.Char)
		assert.Equal(t, "vqgnlzbbkjkljk", parsedLine.Password)
	}
}
