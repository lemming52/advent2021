package binarydiagnostic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagnose(t *testing.T) {
	tests := []struct {
		name        string
		input       []string
		expectedOne int64
		expectedTwo int64
	}{
		{
			name: "part one",
			input: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			expectedOne: 198,
			expectedTwo: 230,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res1, res2 := Diagnose(tt.input)
			assert.Equal(t, tt.expectedOne, res1, "returned value should match expected")
			assert.Equal(t, tt.expectedTwo, res2, "returned value should match expected")
		})
	}
}
