package smokebasin

import (
	"bufio"
	"log"
	"os"
	"sort"
)

const runeOffset = 48

type RiskManagement struct {
	depths  []string
	risk    int
	yMax    int
	xMax    int
	marked  [][]int
	basins  map[int]int
	largest []int
}

func (r *RiskManagement) exploreBasins() {
	basinCounter := 1
	for y, depth := range r.depths {
		for x := range depth {
			if r.explorePoint(y, x, basinCounter) {
				r.checkSize(basinCounter)
				basinCounter++
			}
		}
	}
}

func (r *RiskManagement) explorePoint(y, x, marker int) bool {
	if r.marked[y][x] != 0 || r.depths[y][x] == '9' {
		return false
	}
	r.marked[y][x] = marker
	r.basins[marker]++
	coords := r.getCoords(y, x)
	anyAdjacentLower := false
	for _, c := range coords {
		val := r.depths[c[0]][c[1]]
		if val != '9' {
			anyAdjacentLower = anyAdjacentLower || val < r.depths[y][x]
			r.explorePoint(c[0], c[1], marker)
		}
	}
	if !anyAdjacentLower {
		r.risk += int(r.depths[y][x]) + 1 - runeOffset
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

func (r *RiskManagement) checkSize(marker int) {
	size := r.basins[marker]
	if r.largest[2] == 0 {
		r.largest[2] = size
		return
	}
	sortIndex := sort.Search(3, func(j int) bool { return r.largest[j] >= size })
	switch sortIndex {
	case 1:
		r.largest[0] = size
	case 2:
		r.largest[0] = r.largest[1]
		r.largest[1] = size
	case 3:
		r.largest = []int{
			r.largest[1],
			r.largest[2],
			size,
		}
	}
}

func (r *RiskManagement) basinProduct() int {
	return r.largest[0] * r.largest[1] * r.largest[2]
}

func FindLowPoints(depths []string) (int, int) {
	marked := make([][]int, len(depths))
	for i := range marked {
		marked[i] = make([]int, len(depths[0]))
	}
	manager := &RiskManagement{
		depths:  depths,
		risk:    0,
		yMax:    len(depths),
		xMax:    len(depths[0]),
		marked:  marked,
		basins:  map[int]int{},
		largest: make([]int, 3),
	}
	manager.exploreBasins()
	return manager.risk, manager.basinProduct()
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
