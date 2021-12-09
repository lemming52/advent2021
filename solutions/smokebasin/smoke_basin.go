package smokebasin

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const runeOffset = 48

type RiskManagement struct {
	depths         []string
	risk           int
	marked         [][]bool
	yMax           int
	xMax           int
	basinIndicator [][]int
	biggestBasin   []int
}

func (r *RiskManagement) findLowPoints() {
	for y, depth := range r.depths {
		for x, value := range depth {
			if r.isLowestOfAdjacent(y, x, value) {
				r.risk += int(value) + 1 - runeOffset
			}
		}
	}
}

func (r *RiskManagement) isLowestOfAdjacent(y, x int, value rune) bool {
	if r.marked[y][x] {
		return false
	}

	correct := r.getCoords(y, x)
	if value != '0' {
		for _, c := range correct {
			if value >= rune(r.depths[c[0]][c[1]]) {
				r.marked[y][x] = true
				return false
			}
		}
	}
	for _, c := range correct {
		r.marked[c[0]][c[1]] = true
	}
	return true
}

func (r *RiskManagement) getCoords(y, x int) [][]int {
	coords := [][]int{
		[]int{y - 1, x},
		[]int{y + 1, x},
		[]int{y, x - 1},
		[]int{y, x + 1},
	}
	correct := [][]int{}
	for _, c := range coords {
		if c[0] >= 0 && c[1] >= 0 && c[0] < r.yMax && c[1] < r.xMax {
			correct = append(correct, c)
		}
	}
	return correct
}

func (r *RiskManagement) findBasins() int {
	basinCounter := 0
	for y, depth := range r.depths {
		for x, value := range depth {
			if value == '9' {
				continue
			}
			currentBasin := r.basinIndicator[y][x]
			if currentBasin != 0 {
				continue
			}
			coords := r.getCoords(y, x)
			print(".")
			for _, c := range coords {
				val := r.depths[c[0]][c[1]]
				if val != '9' && r.basinIndicator[c[0]][c[1]] != 0 {
					currentBasin = r.basinIndicator[c[0]][c[1]]
					break
				}
			}
			if currentBasin != 0 {
				r.basinIndicator[y][x] = currentBasin
				for _, c := range coords {
					val := r.depths[c[0]][c[1]]
					if val != '9' && r.basinIndicator[c[0]][c[1]] == 0 {
						r.basinIndicator[c[0]][c[1]] = currentBasin
					}
				}
				continue
			}
			basinCounter++
			r.basinIndicator[y][x] = basinCounter
			for _, c := range coords {
				val := r.depths[c[0]][c[1]]
				if val != '9' && r.basinIndicator[c[0]][c[1]] == 0 {
					r.basinIndicator[c[0]][c[1]] = basinCounter
				}
			}
		}
	}
	basins := map[int]int{}
	for _, b := range r.basinIndicator {
		for _, i := range b {
			if i == 0 {
				continue
			}
			basins[i]++
		}
	}
	sizes := make([]int, len(basins))
	for i := range sizes {
		count := basins[i+1]
		sortIndex := sort.Search(i, func(j int) bool { return sizes[j] > count })
		copy(sizes[sortIndex+1:], sizes[sortIndex:])
		sizes[sortIndex] = count
	}
	total := 1
	fmt.Println(sizes[len(sizes)-3:])
	for _, b := range sizes[len(sizes)-3:] {
		total *= b
	}
	return total
}

func (r *RiskManagement) checkSize(size int) {
	if r.biggestBasin[2] == 0 {
		r.biggestBasin[2] = size
		return
	}
	sortIndex := sort.Search(3, func(j int) bool { return r.biggestBasin[j] >= size })
	switch sortIndex {
	case 1:
		r.biggestBasin[0] = size
	case 2:
		r.biggestBasin[0] = r.biggestBasin[1]
		r.biggestBasin[1] = size
	case 3:
		r.biggestBasin = []int{
			r.biggestBasin[1],
			r.biggestBasin[2],
			size,
		}
	}

}

func FindLowPoints(depths []string) (int, int) {
	marked := make([][]bool, len(depths))
	basins := make([][]int, len(depths))
	for i := range marked {
		marked[i] = make([]bool, len(depths[0]))
		basins[i] = make([]int, len(depths[0]))
	}
	manager := &RiskManagement{
		depths:         depths,
		risk:           0,
		marked:         marked,
		yMax:           len(depths),
		xMax:           len(depths[0]),
		basinIndicator: basins,
		biggestBasin:   make([]int, 3),
	}
	manager.findLowPoints()
	return manager.risk, manager.findBasins()
}

func Challenge(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	depths := []string{}
	for scanner.Scan() {
		depths = append(depths, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return FindLowPoints(depths)
}
