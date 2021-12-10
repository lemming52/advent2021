package main

import (
	"advent/solutions/binarydiagnostic"
	"advent/solutions/dive"
	"advent/solutions/giantsquid"
	"advent/solutions/hydrothermalventure"
	"advent/solutions/lanternfish"
	"advent/solutions/sevensegmentsearch"
	"advent/solutions/smokebasin"
	"advent/solutions/sonarsweep"
	"advent/solutions/syntaxscoring"
	"advent/solutions/treacheryofwhales"
	"flag"
	"fmt"
	"time"
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
		"lanternfish",
		"treacheryofwhales",
		"sevensegmentsearch",
		"smokebasin",
		"syntaxscoring",
	}
	if *all {
		previous := time.Now()
		fmt.Println("Start Time: ", time.Now())
		for _, c := range completed {
			current := time.Now()
			fmt.Println(RunChallenge(c), " Duration/ms: ", float64(current.Sub(previous).Microseconds())/1000)
			previous = current
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
	case "lanternfish":
		A, B := lanternfish.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "treacheryofwhales":
		A, B := treacheryofwhales.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "sevensegmentsearch":
		A, B := sevensegmentsearch.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "smokebasin":
		A, B := smokebasin.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "syntaxscoring":
		A, B := syntaxscoring.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	}
	return res
}
