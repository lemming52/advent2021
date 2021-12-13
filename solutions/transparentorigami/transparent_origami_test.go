package transparentorigami

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoldOrigami(t *testing.T) {
	tests := []struct {
		name  string
		dots []string
		folds []string
		first int
		all int
	}{
		{
			name: "first",
			dots: []string{
				"6,10",
				"0,14",
				"9,10",
				"0,3",
				"10,4",
				"4,11",
				"6,0",
				"6,12",
				"4,1",
				"0,13",
				"10,12",
				"3,4",
				"3,0",
				"8,4",
				"1,10",
				"2,14",
				"8,10",
				"9,0",
			},
			folds: []string{
				"fold along y=7",
			},
			first: 17,
			all: 17,
		},{
			name: "second",
			dots: []string{
				"6,10",
				"0,14",
				"9,10",
				"0,3",
				"10,4",
				"4,11",
				"6,0",
				"6,12",
				"4,1",
				"0,13",
				"10,12",
				"3,4",
				"3,0",
				"8,4",
				"1,10",
				"2,14",
				"8,10",
				"9,0",
			},
			folds: []string{
				"fold along y=7",
				"fold along x=5",
			},
			first: 17,
			all: 16,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			first, all, err := FoldOrigami(tt.dots, tt.folds)
			assert.Nil(t, err)
			assert.Equal(t, tt.first, first, "returned value should match expected")
			assert.Equal(t, tt.all, all, "returned value should match expected")
		})
	}
}
