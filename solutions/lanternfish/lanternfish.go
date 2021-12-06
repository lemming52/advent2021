package lanternfish

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type LanternFish struct {
	age int
}

func (f *LanternFish) Live() *LanternFish {
	if f.age == 0 {
		f.age = 6
		return &LanternFish{age: 8}
	}
	f.age--
	return nil
}

func SpawnFish(s string) (map[int][]*LanternFish, map[int]int, error) {
	fishAges := strings.Split(s, ",")
	counts := map[int]int{}
	for _, f := range fishAges {
		age, err := strconv.Atoi(f)
		if err != nil {
			return nil, nil, err
		}
		counts[age]++
	}
	fish := make(map[int][]*LanternFish, len(counts))
	for age := range counts {
		fish[age] = []*LanternFish{&LanternFish{age: age}}
	}
	return fish, counts, nil
}

func CircleOfLife(s string, cycles int) (int, error) {
	fishes, counts, err := SpawnFish(s)
	if err != nil {
		return 0, err
	}
	totals := map[int]int{}
	for age := range fishes {
		for i := 0; i < cycles; i++ {
			newFish := []*LanternFish{}
			for _, f := range fishes[age] {
				newF := f.Live()
				if newF != nil {
					newFish = append(newFish, newF)
				}
			}
			fishes[age] = append(fishes[age], newFish...)
		}
		totals[age] = len(fishes[age]) * counts[age]
		fishes[age] = nil
	}
	count := 0
	for _, total := range totals {
		count += total
	}
	return count, nil
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
	res, err := CircleOfLife(fish, 80)
	if err != nil {
		log.Fatal(err)
	}
	res2, err := CircleOfLife(fish, 256)
	if err != nil {
		log.Fatal(err)
	}
	return res, res2
}
