package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLine_Basic(t *testing.T) {
	parsedLine, err := NewLine("dull black bags contain 2 shiny brown bags, 3 plaid bronze bags, 3 wavy teal bags, 3 dull chartreuse bags.")

	if assert.NoError(t, err) {
		assert.Equal(t, parsedLine.BagColor, "dull black")
		assert.Equal(t, parsedLine.Children, []ChildBag{
			{
				BagColor: "shiny brown",
				Count:    2,
			},
			{
				BagColor: "plaid bronze",
				Count:    3,
			},
			{
				BagColor: "wavy teal",
				Count:    3,
			},
			{
				BagColor: "dull chartreuse",
				Count:    3,
			},
		})
	}
}

func TestNewLine_Empty(t *testing.T) {
	parsedLine, err := NewLine("muted crimson bags contain no other bags.")

	if assert.NoError(t, err) {
		assert.Equal(t, parsedLine.BagColor, "muted crimson")
		assert.Empty(t, parsedLine.Children)
	}
}
