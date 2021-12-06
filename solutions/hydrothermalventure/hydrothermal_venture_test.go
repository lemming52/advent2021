package hydrothermalventure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpecifyFullPath(t *testing.T) {
	tests := []struct {
		name     string
		input    *Vent
		expected []string
	}{
		{
			name: "vertical",
			input: &Vent{
				x1: 1,
				x2: 1,
				y1: 1,
				y2: 3,
			},
			expected: []string{"1,1", "1,2", "1,3"},
		}, {
			name: "horizontal",
			input: &Vent{
				x1: 9,
				x2: 7,
				y1: 7,
				y2: 7,
			},
			expected: []string{"9,7", "8,7", "7,7"},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := tt.input.specifyFullPath()
			assert.Equal(t, tt.expected, res, "returned value should match expected	")
		})
	}
}

func TestHazardousVents(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		flag     bool
		expected int
	}{
		{
			name: "not diagonal",
			input: []string{
				"0,9 -> 5,9",
				"8,0 -> 0,8",
				"9,4 -> 3,4",
				"2,2 -> 2,1",
				"7,0 -> 7,4",
				"6,4 -> 2,0",
				"0,9 -> 2,9",
				"3,4 -> 1,4",
				"0,0 -> 8,8",
				"5,5 -> 8,2",
			},
			flag:     false,
			expected: 5,
		}, {
			name: "diagonal",
			input: []string{
				"0,9 -> 5,9",
				"8,0 -> 0,8",
				"9,4 -> 3,4",
				"2,2 -> 2,1",
				"7,0 -> 7,4",
				"6,4 -> 2,0",
				"0,9 -> 2,9",
				"3,4 -> 1,4",
				"0,0 -> 8,8",
				"5,5 -> 8,2",
			},
			flag:     true,
			expected: 12,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res, err := HazardousVents(tt.input, tt.flag)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res, "returned value should match expected")
		})
	}
}
