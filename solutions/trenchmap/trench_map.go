package trenchmap

import (
	"advent/solutions/utils"
	"fmt"
	"log"
	"strconv"
)

type Image struct {
	grid        [][]bool
	enhancement string
	emptyValue  bool
}

func newImage(lines []string, enhancement string) *Image {
	grid := make([][]bool, len(lines))
	for i, l := range lines {
		grid[i] = make([]bool, len(l))
		for j, r := range l {
			if r == '#' {
				grid[i][j] = true
			}
		}
	}
	return &Image{
		grid:        grid,
		enhancement: enhancement,
		emptyValue:  false,
	}
}

func (img *Image) iterate() error {
	newY, newX := len(img.grid)+2, len(img.grid[0])+2
	newGrid := make([][]bool, newY)
	for i := range newGrid {
		newGrid[i] = make([]bool, newX)
		for j := range newGrid[i] {
			val, err := img.enhancePixel(i, j, newY, newX)
			if err != nil {
				return err
			}
			newGrid[i][j] = val
		}
	}
	img.grid = newGrid
	if img.emptyValue {
		img.emptyValue = img.enhancement[511] == '#'
	} else {
		img.emptyValue = img.enhancement[0] == '#'
	}
	return nil
}

func (img *Image) enhancePixel(y, x, yMax, xMax int) (bool, error) {
	s := make([]rune, 9)
	possible := getCoords(y, x)
	for i, c := range possible {
		yc, xc := c[0], c[1]
		if yc < 1 || yc > yMax-2 || xc < 1 || xc > yMax-2 {
			if img.emptyValue {
				s[i] = '1'
			} else {
				s[i] = '0'
			}
		} else {
			if img.grid[yc-1][xc-1] {
				s[i] = '1'
			} else {
				s[i] = '0'
			}
		}
	}
	val, err := strconv.ParseInt(string(s), 2, 64)
	if err != nil {
		return false, err
	}
	return img.enhancement[val] == '#', nil
}

func (img *Image) lit() int {
	count := 0
	for i := range img.grid {
		for j := range img.grid[i] {
			if img.grid[i][j] {
				count++
			}
		}
	}
	return count
}

func (img *Image) print() {
	for i := range img.grid {
		s := make([]rune, len(img.grid[i]))
		for j := range img.grid[i] {
			if img.grid[i][j] {
				s[j] = '#'
			} else {
				s[j] = ' '
			}
		}
		fmt.Println(string(s))
	}
}

func getCoords(y, x int) [][]int {
	return [][]int{
		{y - 1, x - 1},
		{y - 1, x},
		{y - 1, x + 1},
		{y, x - 1},
		{y, x},
		{y, x + 1},
		{y + 1, x - 1},
		{y + 1, x},
		{y + 1, x + 1},
	}
}

func EnhanceImage(image []string, enhancement string) (int, int, error) {
	img := newImage(image, enhancement)
	for i := 0; i < 2; i++ {
		err := img.iterate()
		if err != nil {
			return 0, 0, err
		}
	}
	second := img.lit()
	for i := 2; i < 50; i++ {
		err := img.iterate()
		if err != nil {
			return 0, 0, err
		}
	}
	return second, img.lit(), nil
}

func Challenge(path string) (int, int) {
	lines, err := utils.ReadStrings(path)
	if err != nil {
		log.Fatal(err)
	}
	a, b, err := EnhanceImage(lines[2:], lines[0])
	if err != nil {
		log.Fatal(err)
	}
	return a, b
}
