package hydrothermalventure

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

// 959,103 -> 139,923
const ventPattern = `(\d+)\,(\d+) -> (\d+)\,(\d+)`

type Vent struct {
	x1, y1, x2, y2 int
}

func (v *Vent) isNotDiagonal() bool {
	return v.x1 == v.x2 || v.y1 == v.y2
}

func (v *Vent) specifyFullPath() []string {
	coords := []string{}
	xDelta, yDelta := v.x2-v.x1, v.y2-v.y1
	var stepCount int
	if xDelta != 0 {
		stepCount = int(math.Abs(float64(xDelta)))
	} else {
		stepCount = int(math.Abs(float64(yDelta)))
	}
	xDirection, yDirection := xDelta/stepCount, yDelta/stepCount
	xi, yi := v.x1, v.y1
	for i := 0; i <= stepCount; i++ {
		coords = append(coords, fmt.Sprintf("%d,%d", xi, yi))
		xi += xDirection
		yi += yDirection
	}
	return coords
}

type VentMap struct {
	ventPositions map[string]int
	ventPattern   *regexp.Regexp
	overlapCount  int
}

func HazardousVents(vents []string, considerDiagonal bool) (int, error) {
	pattern, err := regexp.Compile(ventPattern)
	if err != nil {
		return 0, err
	}
	ventMap := &VentMap{
		ventPositions: map[string]int{},
		ventPattern:   pattern,
	}
	for _, vs := range vents {
		vent, err := ventMap.extractVent(vs)
		if err != nil {
			return 0, nil
		}
		if !considerDiagonal && !vent.isNotDiagonal() {
			continue
		}
		for _, coords := range vent.specifyFullPath() {
			count, ok := ventMap.ventPositions[coords]
			if !ok {
				ventMap.ventPositions[coords] = 1
				continue
			} else if count == 1 {
				ventMap.overlapCount++
			}
			ventMap.ventPositions[coords]++
		}
	}
	return ventMap.overlapCount, nil
}

func (m *VentMap) extractVent(s string) (*Vent, error) {
	ventComponents := m.ventPattern.FindStringSubmatch(s)
	coords := make([]int, 4)
	for i, v := range ventComponents[1:] {
		val, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		coords[i] = val
	}
	return &Vent{
		x1: coords[0],
		x2: coords[2],
		y1: coords[1],
		y2: coords[3],
	}, nil
}

func Challenge(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	vents := []string{}
	for scanner.Scan() {
		vents = append(vents, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	horizontal, err := HazardousVents(vents, false)
	if err != nil {
		log.Fatal(err)
	}
	diagonal, err := HazardousVents(vents, true)
	if err != nil {
		log.Fatal(err)
	}
	return horizontal, diagonal
}
