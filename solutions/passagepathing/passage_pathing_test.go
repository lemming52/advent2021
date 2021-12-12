package passagepathing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapCaves(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		paths int
		twice int
	}{
		{
			name: "base",
			input: []string{
				"start-A",
				"start-b",
				"A-c",
				"A-b",
				"b-d",
				"A-end",
				"b-end",
			},
			paths: 10,
			twice: 36,
		}, {
			name: "bigger",
			input: []string{
				"dc-end",
				"HN-start",
				"start-kj",
				"dc-start",
				"dc-HN",
				"LN-dc",
				"HN-end",
				"kj-sa",
				"kj-HN",
				"kj-dc",
			},
			paths: 19,
			twice: 103,
		}, {
			name: "biggest",
			input: []string{
				"fs-end",
				"he-DX",
				"fs-he",
				"start-DX",
				"pj-DX",
				"end-zg",
				"zg-sl",
				"zg-pj",
				"pj-he",
				"RW-he",
				"fs-DX",
				"pj-RW",
				"zg-RW",
				"start-pj",
				"he-WI",
				"zg-he",
				"pj-fs",
				"start-RW",
			},
			paths: 226,
			twice: 3509,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			paths, incomplete := MapCaves(tt.input)
			assert.Equal(t, tt.paths, paths, "returned value should match expected")
			assert.Equal(t, tt.twice, incomplete, "returned value should match expected")
		})
	}
}
