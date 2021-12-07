package treacheryofwhales

import (
	"bufio"

	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type FuelCostCounter struct {
	costs map[float64]int
}

func (c *FuelCostCounter) calculateCost(f float64) int {
	if val, ok := c.costs[f]; ok {
		return val
	}
	c.costs[f] = int((f*f + f) / 2)
	return c.costs[f]
}

func CrabSubmarines(s string) (int, int, error) {
	crabs := strings.Split(s, ",")
	count := len(crabs)
	crabPositions := make([]float64, count)
	crabIndex := 0
	total := float64(0)
	for _, c := range crabs {
		val, err := strconv.ParseFloat(c, 64)
		if err != nil {
			return 0, 0, err
		}
		total += val
		if crabIndex == 0 {
			crabIndex++
			crabPositions[0] = val
			continue
		}
		sortIndex := sort.Search(crabIndex, func(j int) bool { return crabPositions[j] >= val })
		copy(crabPositions[sortIndex+1:], crabPositions[sortIndex:])
		crabPositions[sortIndex] = val
		crabIndex++
	}
	mean := total / float64(count)
	meanLower, meanHigher := math.Round(mean), math.Trunc(mean)
	complexTotal := computeComplexTotal(crabPositions, meanLower, meanHigher)

	var lowerTotal float64
	medianIndex := count / 2
	if count%2 == 0 && crabPositions[medianIndex-1] != crabPositions[medianIndex] {
		lowerTotal = computeTotal(crabPositions, -1)
	}
	upperTotal := computeTotal(crabPositions, count/2)

	if lowerTotal != 0 && lowerTotal < upperTotal {
		return int(lowerTotal), complexTotal, nil
	}
	return int(upperTotal), complexTotal, nil
}

func computeTotal(crabs []float64, medianIndex int) float64 {
	total := float64(0)
	median := crabs[medianIndex]
	for _, c := range crabs {
		total += math.Abs(c - median)
	}
	return total
}

func computeComplexTotal(crabs []float64, meanLower, meanHigher float64) int {
	lowerTotal, higherTotal := 0, 0
	calc := &FuelCostCounter{costs: map[float64]int{}}
	for _, c := range crabs {
		deltaLower, deltaHigher := math.Abs(c-meanLower), math.Abs(c-meanHigher)
		lowerTotal += calc.calculateCost(deltaLower)
		higherTotal += calc.calculateCost(deltaHigher)
	}
	if lowerTotal < higherTotal {
		return lowerTotal
	}
	return higherTotal
}

func Challenge(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fish := scanner.Text()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	res1, res2, err := CrabSubmarines(fish)
	if err != nil {
		log.Fatal(err)
	}
	return res1, res2
}
