package chiton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvoidChiton(t *testing.T) {
	tests := []struct {
		name     string
		chiton   []string
		expected int
	}{
		{
			name: "first",
			chiton: []string{
				"1163751742",
				"1381373672",
				"2136511328",
				"3694931569",
				"7463417111",
				"1319128137",
				"1359912421",
				"3125421639",
				"1293138521",
				"2311944581",
			},
			expected: 40,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := AvoidChiton(tt.chiton)
			assert.Equal(t, tt.expected, res, "returned value should match expected")
		})
	}
}
