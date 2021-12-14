package extendedpolymerization

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoldOrigami(t *testing.T) {
	tests := []struct {
		name         string
		polymer      string
		instructions []string
		steps        int
		expected     int
	}{
		{
			name:    "first",
			polymer: "NNCB",
			instructions: []string{
				"CH -> B",
				"HH -> N",
				"CB -> H",
				"NH -> C",
				"HB -> C",
				"HC -> B",
				"HN -> C",
				"NN -> C",
				"BH -> H",
				"NC -> B",
				"NB -> B",
				"BN -> B",
				"BB -> N",
				"BC -> B",
				"CC -> N",
				"CN -> C",
			},
			steps:    10,
			expected: 1588,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res, _, err := ExtendPolymer(tt.polymer, tt.instructions, tt.steps)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res, "returned value should match expected")
		})
	}
}
