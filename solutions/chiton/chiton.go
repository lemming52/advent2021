package chiton

import (
	"bufio"
	"log"
	"os"
)

const runeOffset = 48

type Cave struct {
	chiton   []string
	minimums [][]int
	yMax     int
	xMax     int
}

func (c *Cave) explore(y, x, total, min int) int {
	total = total + val(c.chiton[y][x])
	if y == c.yMax && x == c.xMax {
		return total
	}
	if total > min {
		return total
	}
	coords := c.getCoords(y, x)
	for _, yx := range coords {
		newTotal := c.explore(yx[0], yx[1], total, min)
		if newTotal < min {
			min = newTotal
		}
	}
	return min
}

func (c *Cave) minimumFromPoint(y, x int) {
	value := val(c.chiton[y][x])
	if y == c.yMax && x == c.xMax {
		c.minimums[y][x] = value
		return
	}
	minimum := 0
	coords := c.getCoords(y, x)
	for _, yx := range coords {
		adjacentMinimum := c.minimums[yx[0]][yx[1]]
		if adjacentMinimum == 0 {
			continue
		}
		if adjacentMinimum+value < minimum || minimum == 0 {
			minimum = adjacentMinimum + value
		}
	}
	c.minimums[y][x] = minimum
}

func (c *Cave) getCoords(y, x int) [][]int {
	coords := [][]int{
		[]int{y - 1, x},
		[]int{y, x - 1},
		[]int{y + 1, x},
		[]int{y, x + 1},
	}
	correct := [][]int{}
	for _, yx := range coords {
		if yx[0] >= 0 && yx[1] >= 0 && yx[0] <= c.yMax && yx[1] <= c.xMax {
			correct = append(correct, yx)
		}
	}
	return correct
}

func (c *Cave) correctPoint(y, x int) {
	coords := c.getCoords(y, x)
	for _, yx := range coords {
		adjacentMinimum := c.minimums[yx[0]][yx[1]]
		if adjacentMinimum > c.minimums[y][x]+val(c.chiton[yx[0]][yx[1]]) {
			c.minimums[yx[0]][yx[1]] = c.minimums[y][x] + val(c.chiton[yx[0]][yx[1]])
			c.correctPoint(yx[0], yx[1])
		}
	}
}

func val(r byte) int {
	return int(r) - runeOffset
}

func AvoidChiton(chiton []string) int {
	cave := &Cave{
		chiton: chiton,
		yMax:   len(chiton) - 1,
		xMax:   len(chiton[0]) - 1,
	}
	mins := make([][]int, cave.yMax+1)
	for i := 0; i <= cave.xMax; i++ {
		mins[i] = make([]int, cave.xMax+1)
	}
	cave.minimums = mins
	for i := cave.yMax; i >= 0; i-- {
		for j := cave.xMax; j >= 0; j-- {
			cave.minimumFromPoint(i, j)
		}
	}

	for i := cave.xMax; i >= 0; i-- {
		for j := cave.yMax; j >= 0; j-- {
			cave.correctPoint(j, i)
		}
	}
	return cave.minimums[0][0] - val(chiton[0][0])
}

func buildBiggerChiton(chiton []string, factor int) []string {
	newChiton := make([]string, len(chiton)*factor)
	for i := 0; i < len(chiton)*factor; i++ {
		layer := make([]rune, len(chiton[0])*factor)
		for j := 0; j < len(chiton[0])*factor; j++ {
			value := (val(chiton[i%len(chiton)][j%len(chiton[0])]) + i/len(chiton) + j/len(chiton)) % 9
			if value == 0 {
				value = 9
			}
			layer[j] = rune(value + runeOffset)
		}
		newChiton[i] = string(layer)
	}
	return newChiton
}

func Challenge(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	chiton := []string{}
	for scanner.Scan() {
		chiton = append(chiton, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	bigger := buildBiggerChiton(chiton, 5)
	return AvoidChiton(chiton), AvoidChiton(bigger)
}
