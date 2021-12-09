package smokebasin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowPoints(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
		basins   int
	}{
		{
			name: "base",
			input: []string{
				"2199943210",
				"3987894921",
				"9856789892",
				"8767896789",
				"9899965678",
			},
			expected: 15,
			basins:   1134,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res1, res2 := FindLowPoints(tt.input)
			assert.Equal(t, tt.expected, res1, "returned value should match expected")
			assert.Equal(t, tt.basins, res2, "returned value should match expected")
		})
	}
}
