package packetdecoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePackets(t *testing.T) {
	tests := []struct {
		name     string
		packet   string
		expected int
	}{
		{
			name:     "literal",
			packet:   "D2FE28",
			expected: 6,
		}, {
			name:     "bit length",
			packet:   "38006F45291200",
			expected: 9,
		}, {
			name:     "packet count",
			packet:   "EE00D40C823060",
			expected: 14,
		}, {
			name:     "alpha",
			packet:   "8A004A801A8002F478",
			expected: 16,
		}, {
			name:     "bravo",
			packet:   "620080001611562C8802118E34",
			expected: 12,
		}, {
			name:     "gamma",
			packet:   "C0015000016115A2E0802F182340",
			expected: 23,
		}, {
			name:     "delta",
			packet:   "A0016C880162017C3686B18A3D4780",
			expected: 31,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res, _, err := ParsePackets(tt.packet)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res, "returned value should match expected")
		})
	}
}

func TestParsePacketsValue(t *testing.T) {
	tests := []struct {
		name     string
		packet   string
		expected int
	}{
		{
			name:     "C200B40A82",
			packet:   "C200B40A82",
			expected: 3,
		}, {
			name:     "04005AC33890",
			packet:   "04005AC33890",
			expected: 54,
		}, {
			name:     "880086C3E88112",
			packet:   "880086C3E88112",
			expected: 7,
		}, {
			name:     "CE00C43D881120",
			packet:   "CE00C43D881120",
			expected: 9,
		}, {
			name:     "D8005AC2A8F0",
			packet:   "D8005AC2A8F0",
			expected: 1,
		}, {
			name:     "F600BC2D8F",
			packet:   "F600BC2D8F",
			expected: 0,
		}, {
			name:     "9C005AC2F8F0",
			packet:   "9C005AC2F8F0",
			expected: 0,
		}, {
			name:     "9C0141080250320F1802104A08",
			packet:   "9C0141080250320F1802104A08",
			expected: 1,
		}, {
			name:     "literal",
			packet:   "D2FE28",
			expected: 2021,
		}, {
			name:     "bit length",
			packet:   "38006F45291200",
			expected: 1,
		}, {
			name:     "packet count",
			packet:   "EE00D40C823060",
			expected: 3,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			_, res, err := ParsePackets(tt.packet)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res, "returned value should match expected")
		})
	}
}
