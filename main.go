package main

import (
	"advent/solutions/alu"
	"advent/solutions/amphipod"
	"advent/solutions/beaconscanner"
	"advent/solutions/binarydiagnostic"
	"advent/solutions/chiton"
	"advent/solutions/diracdie"
	"advent/solutions/dive"
	"advent/solutions/dumbooctopus"
	"advent/solutions/extendedpolymerization"
	"advent/solutions/giantsquid"
	"advent/solutions/hydrothermalventure"
	"advent/solutions/lanternfish"
	"advent/solutions/packetdecoder"
	"advent/solutions/passagepathing"
	"advent/solutions/reactorreboot"
	"advent/solutions/seacucumber"
	"advent/solutions/sevensegmentsearch"
	"advent/solutions/smokebasin"
	"advent/solutions/snailmaths"
	"advent/solutions/sonarsweep"
	"advent/solutions/syntaxscoring"
	"advent/solutions/transparentorigami"
	"advent/solutions/treacheryofwhales"
	"advent/solutions/trenchmap"
	"advent/solutions/trickshot"
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
		"dumbooctopus",
		"passagepathing",
		"transparentorigami",
		"extendedpolymerization",
		"chiton",
		"packetdecoder",
		"trickshot",
		"snailmaths",
		"beaconscanner",
		"trenchmap",
		"diracdie",
		"reactorreboot",
		"amphipod",
		"alu",
		"seacucumber",
	}
	if *all {
		previous := time.Now()
		fmt.Println("Start Time: ", time.Now())
		for _, c := range completed {
			s := RunChallenge(c)
			current := time.Now()
			fmt.Println(s, " Duration/ms: ", float64(current.Sub(previous).Microseconds())/1000)
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
	case "dumbooctopus":
		A, B := dumbooctopus.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "passagepathing":
		A, B := passagepathing.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "transparentorigami":
		A, B := transparentorigami.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "extendedpolymerization":
		A, B := extendedpolymerization.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "chiton":
		A, B := chiton.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "packetdecoder":
		A, B := packetdecoder.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "trickshot":
		A, B := trickshot.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "snailmaths":
		A, B := snailmaths.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "beaconscanner":
		A, B := beaconscanner.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "trenchmap":
		A, B := trenchmap.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "diracdie":
		A, B := diracdie.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "reactorreboot":
		A, B := reactorreboot.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "amphipod":
		A, B := amphipod.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "alu":
		A, B := alu.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "seacucumber":
		A, B := seacucumber.Challenge(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	}
	return res
}
