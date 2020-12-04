package rules

import (
	"testing"

	"github.com/jlevitt/katas/advent/advent2020/day2/internal/parse"
	"github.com/stretchr/testify/assert"
)

func TestCharPositionedAt(t *testing.T) {
	for name, tc := range map[string]struct {
		line  parse.Line
		valid bool
	}{
		"valid": {
			line:  parse.Line{A: 1, B: 3, Char: "a", Password: "abcde"},
			valid: true,
		},
		"neither position matches": {
			line:  parse.Line{A: 1, B: 3, Char: "b", Password: "cdefg"},
			valid: false,
		},
		"both positions match": {
			line:  parse.Line{A: 2, B: 9, Char: "c", Password: "ccccccccc"},
			valid: false,
		},
	} {
		t.Run(t.Name()+"_"+name, func(t *testing.T) {
			assert.Equal(t, tc.valid, CharPositionedAt(tc.line))
		})
	}
}
