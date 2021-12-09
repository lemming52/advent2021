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

func TestCheckSize(t *testing.T) {
	tests := []struct {
		name     string
		existing []int
		size     int
		expected []int
	}{
		{
			name: "smaller",
			existing: []int{
				2,
				4,
				6,
			},
			size: 1,
			expected: []int{
				2,
				4,
				6,
			},
		}, {
			name: "middle",
			existing: []int{
				2,
				4,
				6,
			},
			size: 5,
			expected: []int{
				4,
				5,
				6,
			},
		}, {
			name: "bigger",
			existing: []int{
				2,
				4,
				6,
			},
			size: 7,
			expected: []int{
				4,
				6,
				7,
			},
		}, {
			name: "equal low",
			existing: []int{
				2,
				4,
				6,
			},
			size: 2,
			expected: []int{
				2,
				4,
				6,
			},
		}, {
			name: "equal middle",
			existing: []int{
				2,
				4,
				6,
			},
			size: 4,
			expected: []int{
				4,
				4,
				6,
			},
		}, {
			name: "equal big",
			existing: []int{
				2,
				4,
				6,
			},
			size: 6,
			expected: []int{
				4,
				6,
				6,
			},
		}, {
			name:     "missing",
			existing: make([]int, 3),
			size:     1,
			expected: []int{
				1,
				0,
				0,
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			man := &RiskManagement{biggestBasin: tt.existing}
			man.checkSize(tt.size)
			assert.Equal(t, tt.expected, man.biggestBasin, "returned value should match expected")
		})
	}
}
