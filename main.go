package main

import (
	"advent/solutions/sonarsweep"
	"flag"
	"fmt"
)

// res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
// res = fmt.Sprintf("%s Results A: %d", challenge, A)

func main() {
	var challenge string
	flag.StringVar(&challenge, "challenge", "sonarsweep", "name of challenge")
	all := flag.Bool("all", false, "display all results")
	flag.Parse()

	completed := []string{
		"sonarsweep",
	}
	if *all {
		for _, c := range completed {
			fmt.Println(RunChallenge(c))
		}
	} else {
		fmt.Println(RunChallenge(challenge))
	}

}

func RunChallenge(challenge string) string {
	var res string
	input := fmt.Sprintf("inputs/%s.txt", challenge)
	switch challenge {
	case "sonarsweep":
		A, B := sonarsweep.LoadSonar(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	}
	return res
}
