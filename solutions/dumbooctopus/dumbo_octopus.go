package dumbooctopus

import (
	"bufio"
	"log"
	"os"
)

const runeOffset = 48

type Octopus struct {
	level int
	step  int
}

type Configuration struct {
	octopodes        [][]*Octopus
	step             int
	flashes          int
	markedFlashCount int
	flashesPerStep   int
	yMax             int
	xMax             int
}

func (c *Configuration) Cycle(step int) bool {
	c.flashesPerStep = 0
	for y := 0; y < c.yMax; y++ {
		for x := 0; x < c.xMax; x++ {
			flashed := c.powerUp(y, x, step)
			neighbours := [][]int{}
			if flashed {
				neighbours = c.getCoords(y, x)
			}
			for len(neighbours) != 0 {
				newNeighbours := [][]int{}
				for _, n := range neighbours {
					flashed = c.powerUp(n[0], n[1], step)
					if flashed {
						newNeighbours = append(newNeighbours, c.getCoords(n[0], n[1])...)
					}
				}
				neighbours = newNeighbours
			}
		}
	}
	return c.flashesPerStep == c.xMax*c.yMax
}

func (c *Configuration) powerUp(y, x, step int) bool {
	octopus := c.octopodes[y][x]
	if octopus.level == 0 {
		if octopus.step == step {
			return false
		}
	}
	octopus.level++
	octopus.step = step
	if octopus.level == 10 {
		octopus.level = 0
		c.flashes++
		c.flashesPerStep++
		return true
	}
	return false
}

func (c *Configuration) getCoords(y, x int) [][]int {
	coords := [][]int{
		[]int{y - 1, x},
		[]int{y + 1, x},
		[]int{y, x - 1},
		[]int{y, x + 1},
		[]int{y - 1, x - 1},
		[]int{y + 1, x - 1},
		[]int{y - 1, x + 1},
		[]int{y + 1, x + 1},
	}
	correct := [][]int{}
	for _, co := range coords {
		if co[0] >= 0 && co[1] >= 0 && co[0] < c.yMax && co[1] < c.xMax {
			correct = append(correct, co)
		}
	}
	return correct
}

func ModelOctopus(octopodes []string, steps int) (int, int) {
	yMax, xMax := len(octopodes), len(octopodes[0])
	conf := &Configuration{
		octopodes: make([][]*Octopus, yMax),
		yMax:      yMax,
		xMax:      xMax,
	}
	for y, o := range octopodes {
		conf.octopodes[y] = make([]*Octopus, xMax)
		for x, level := range o {
			conf.octopodes[y][x] = &Octopus{
				level: int(level - runeOffset),
				step:  -1,
			}
		}
	}
	i := 0
	for {
		if i == steps {
			conf.markedFlashCount = conf.flashes
		}
		synced := conf.Cycle(i)
		if synced {
			return conf.markedFlashCount, i + 1
		}
		i++
	}
}

func Challenge(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	octopodes := []string{}
	for scanner.Scan() {
		octopodes = append(octopodes, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ModelOctopus(octopodes, 100)
}
