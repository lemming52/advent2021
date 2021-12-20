package beaconscanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanBeacons(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		expected int
		count    int
	}{
		{
			name: "base",
			input: [][]string{
				{
					"0,2,0",
					"4,1,0",
					"3,3,0",
				},
				{
					"-1,-1,0",
					"-5,0,0",
					"-2,1,0",
				},
			},
			expected: 45,
			count:    112,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res, count, err := ScanBeacons(tt.input)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res, "returned value should match expected")
			assert.Equal(t, tt.count, count, "returned value should match expected")

		})
	}
}
