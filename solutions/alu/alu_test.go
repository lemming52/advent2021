package alu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindModelNumber(t *testing.T) {
	tests := []struct {
		name         string
		instructions []string
		expected     int
	}{
		{
			name: "base",
			instructions: []string{
				"inp x",
				"mul x -1",
			},
			expected: -1,
		}, {
			name: "long",
			instructions: []string{
				"inp z",
				"inp x",
				"mul z 3",
				"eql z x",
			},
			expected: -1,
		}, {
			name: "longer",
			instructions: []string{
				"inp w",
				"add z w",
				"mod z 2",
				"div w 2",
				"add y w",
				"mod y 2",
				"div w 2",
				"add x w",
				"mod x 2",
				"div w 2",
				"mod w 2",
			},
			expected: -1,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			expected, err := FindModelNumber(tt.instructions)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, expected, "returned value should match expected")
		})
	}
}
