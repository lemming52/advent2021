package main

import (
	"advent/solutions/binarydiagnostic"
	"advent/solutions/dive"
	"advent/solutions/giantsquid"
	"advent/solutions/hydrothermalventure"
	"advent/solutions/sonarsweep"
	"flag"
	"fmt"
)

func main() {
	var challenge string
	flag.StringVar(&challenge, "challenge", "sonarsweep", "name of challenge")
	all := flag.Bool("all", false, "display all results")
	flag.Parse()

	completed := []string{
		"sonarsweep",
		"dive",
		"binarydiagnostic",
		"giantsquid",
		"hydrothermalventure",
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
	case "dive":
		A, B := dive.LoadDive(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "binarydiagnostic":
		A, B := binarydiagnostic.LoadBD(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "giantsquid":
		A, B := giantsquid.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "hydrothermalventure":
		A, B := hydrothermalventure.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	}
	return res
}
