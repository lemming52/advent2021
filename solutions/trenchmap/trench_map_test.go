package trenchmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnhanceImage(t *testing.T) {
	tests := []struct {
		name        string
		image       []string
		enhancement string
		second      int
		fifty       int
	}{
		{
			name: "backtrack",
			image: []string{
				"#..#.",
				"#....",
				"##..#",
				"..#..",
				"..###",
			},
			enhancement: "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#",
			second:      35,
			fifty:       3351,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			second, fifty, err := EnhanceImage(tt.image, tt.enhancement)
			assert.Nil(t, err)
			assert.Equal(t, tt.second, second, "returned value should match expected")
			assert.Equal(t, tt.fifty, fifty, "returned value should match expected")
		})
	}
}
