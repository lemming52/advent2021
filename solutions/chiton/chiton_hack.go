package chiton

func (c *Cave) minimumFromPoint(y, x int) {
	value := val(c.chiton[y][x])
	if y == c.yMax && x == c.xMax {
		c.minimums[y][x] = value
		return
	}
	minimum := 0
	coords := c.getHackCoords(y, x)
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

func (c *Cave) getHackCoords(y, x int) [][]int {
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
	coords := c.getHackCoords(y, x)
	for _, yx := range coords {
		adjacentMinimum := c.minimums[yx[0]][yx[1]]
		if adjacentMinimum > c.minimums[y][x]+val(c.chiton[yx[0]][yx[1]]) {
			c.minimums[yx[0]][yx[1]] = c.minimums[y][x] + val(c.chiton[yx[0]][yx[1]])
			c.correctPoint(yx[0], yx[1])
		}
	}
}

func AvoidChitonHack(chiton []string) int {
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
