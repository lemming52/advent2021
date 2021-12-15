package chiton

import (
	"bufio"
	"fmt"
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
	if y%5 == 0 {
		fmt.Println(y, x, total, min)
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

func (c *Cave) getCoords(y, x int) [][]int {
	coords := [][]int{
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

func (c *Cave) naiveMinimum() int {
	total := 0
	for i := 1; i <= c.xMax; i++ {
		total += val(c.chiton[0][i])
	}
	for i := 1; i <= c.yMax; i++ {
		total += val(c.chiton[i][c.xMax])
	}
	return total
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
	naiveMinimum := cave.naiveMinimum()
	return cave.explore(0, 0, -val(chiton[0][0]), naiveMinimum)
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
	a := AvoidChiton(chiton)
	return a, 0
}
