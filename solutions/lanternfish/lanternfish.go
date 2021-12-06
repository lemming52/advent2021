package lanternfish

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func SpawnFish(s string) (map[int]int, error) {
	fishAges := strings.Split(s, ",")
	counts := map[int]int{}
	for i := 0; i < 8; i++ {
		counts[i] = 0
	}
	for _, f := range fishAges {
		age, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
		counts[age]++
	}

	return counts, nil
}

func CircleOfLife(s string, cycles int) (int, error) {
	counts, err := SpawnFish(s)
	if err != nil {
		return 0, err
	}
	for i := 0; i < cycles; i++ {
		newFish := counts[0]
		for j := 1; j <= 8; j++ {
			counts[j-1] = counts[j]
		}
		counts[8] = newFish
		counts[6] += newFish
	}
	count := 0
	for _, total := range counts {
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
