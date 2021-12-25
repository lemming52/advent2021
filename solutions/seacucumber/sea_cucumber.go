package seacucumber

import (
	"advent/solutions/utils"
	"fmt"
	"log"
)

type Cucumber struct {
	x, y int
	east bool
}

func newCucumber(y, x int, east bool) *Cucumber {
	return &Cucumber{
		x:    x,
		y:    y,
		east: east,
	}
}

type Seafloor struct {
	grid  [][]*Cucumber
	east  []*Cucumber
	south []*Cucumber
	xMax  int
	yMax  int
}

func newSeafloor(cucumber []string) *Seafloor {
	grid := make([][]*Cucumber, len(cucumber))
	east, south := []*Cucumber{}, []*Cucumber{}
	for y, c := range cucumber {
		grid[y] = make([]*Cucumber, len(c))
		for x, r := range c {
			switch r {
			case '>':
				grid[y][x] = newCucumber(y, x, true)
				east = append(east, grid[y][x])
			case 'v':
				grid[y][x] = newCucumber(y, x, false)
				south = append(south, grid[y][x])
			default:
				grid[y][x] = nil
			}
		}
	}
	return &Seafloor{
		grid:  grid,
		east:  east,
		south: south,
		yMax:  len(grid),
		xMax:  len(grid[0]),
	}
}

func (s *Seafloor) migrate() bool {
	newGrid := make([][]*Cucumber, len(s.grid))
	for i := range newGrid {
		newGrid[i] = make([]*Cucumber, len(s.grid[i]))
	}
	moved := false
	for _, c := range s.east {
		if s.moveCucumber(c, newGrid, true) {
			moved = true
		}
	}
	for _, c := range s.south {
		if s.moveCucumber(c, newGrid, false) {
			moved = true
		}
	}
	if moved {
		s.grid = newGrid
	}
	return moved
}

func (s *Seafloor) moveCucumber(c *Cucumber, grid [][]*Cucumber, east bool) bool {
	y, x := s.getNextSpace(c)
	if east {
		if s.grid[y][x] != nil {
			grid[c.y][c.x] = c
			return false
		}
		grid[y][x] = c
		c.x = x
		c.y = y
		return true
	} else {
		if (s.grid[y][x] != nil && !s.grid[y][x].east) || grid[y][x] != nil {
			grid[c.y][c.x] = c
			return false
		}
		grid[y][x] = c
		c.x = x
		c.y = y
		return true
	}

}

func (s *Seafloor) getNextSpace(c *Cucumber) (int, int) {
	var y, x int
	if c.east {
		y = c.y
		x = c.x + 1
		if x >= s.xMax {
			x = 0
		}
	} else {
		y = c.y + 1
		x = c.x
		if y >= s.yMax {
			y = 0
		}
	}
	return y, x
}

func (s *Seafloor) print(step int) {
	fmt.Println(step)
	for _, g := range s.grid {
		s := ""
		for _, c := range g {
			if c == nil {
				s += " "
				continue
			}
			if c.east {
				s += ">"
			} else {
				s += "v"
			}
		}
		fmt.Println(s)
	}
}

func Migrate(cucumber []string) int {
	seafloor := newSeafloor(cucumber)
	moved := true
	count := 0
	for moved {
		moved = seafloor.migrate()
		count++
	}
	return count
}

func Challenge(path string) (int, int) {
	lines, err := utils.ReadStrings(path)
	if err != nil {
		log.Fatal(err)
	}
	val := Migrate(lines)
	if err != nil {
		log.Fatal(err)
	}
	return val, 0
}
