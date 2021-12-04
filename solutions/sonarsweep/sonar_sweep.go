package sonarsweep

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// SonarSweep counts the number of times a ping depth is higher than the preceding ping
func SonarSweep(pings []int64, window int) int {
	count := 0
	for i, p := range pings[window:] {
		if p > pings[i] {
			count++
		}
	}
	return count
}

// LoadSonar loads an input text file and executes the sonar sweep
func LoadSonar(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	depths := []int64{}
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, int64(n))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return SonarSweep(depths, 1), SonarSweep(depths, 3)
}
