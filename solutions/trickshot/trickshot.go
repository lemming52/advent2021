package trickshot

import (
	"advent/solutions/utils"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
)

const targetPattern = `target area: x=(\d{1,3})\.\.(\d{1,3}), y=(\-{0,1}\d{1,3})\.\.(\-{0,1}\d{1,3})`

type velocity struct {
	x, y, t int
}

func Trickshot(target string) (int, int, error) {
	pattern, err := regexp.Compile(targetPattern)
	if err != nil {
		return 0, 0, err
	}
	components := pattern.FindStringSubmatch(target)
	bounds, err := getBounds(components[1:])
	if err != nil {
		return 0, 0, err
	}
	uy := findYVelocity(bounds[2], bounds[3])
	count := findVelocities(bounds)
	return findYMax(uy), count, nil
}

func getBounds(c []string) ([]int, error) {
	bounds := make([]int, 4)
	for i, c := range c {
		val, err := strconv.Atoi(c)
		if err != nil {
			return nil, err
		}
		bounds[i] = int(val)
	}
	if bounds[0] > bounds[1] {
		bounds[0], bounds[1] = bounds[1], bounds[0]
	}
	if bounds[3] > bounds[2] {
		bounds[2], bounds[3] = -1*bounds[3], -1*bounds[2]
	} else {
		bounds[2], bounds[3] = -1*bounds[2], -1*bounds[3]
	}
	return bounds, nil
}

func findYVelocity(ya, yb int) int {
	return yb * -1
}

func findYMax(uy int) int {
	return uy*uy - uy*(uy-1)/2
}

func findVelocities(bounds []int) int {
	count := 0
	xSingle, xArcing, xStopped := findXVelocities(bounds[0], bounds[1])
	ySingle, yArcing, yStopped := findYVelocities(bounds[2], bounds[3])
	count += len(xSingle) * len(ySingle)
	arcedVelocities := map[string]bool{}
	for _, y := range yArcing {
		for _, x := range xArcing {
			if x.t == y.t {
				key := fmt.Sprintf("%d, %d", x.x, y.y)
				ok := arcedVelocities[key]
				if !ok {
					arcedVelocities[key] = true
					count++
				}
			}
		}
		for _, x := range xStopped {
			if y.t >= x {
				key := fmt.Sprintf("%d, %d", x, y.y)
				ok := arcedVelocities[key]
				if !ok {
					arcedVelocities[key] = true
					count++
				}
			}
		}
	}
	for _, x := range xStopped {
		for _, y := range yStopped {
			key := fmt.Sprintf("%d, %d", x, y)
			ok := arcedVelocities[key]
			if !ok {
				arcedVelocities[key] = true
				count++
			}
		}
	}
	return count
}

func findXVelocities(xa, xb int) ([]int, []*velocity, []int) {
	singleStep := make([]int, xb-xa+1)
	for x := xa; x <= xb; x++ {
		singleStep[x-xa] = x
	}
	arcing := findXArcingVelocities(xa, xb)
	stopped := findXStoppedVelocities(xa, xb)
	return singleStep, arcing, stopped
}

func findXArcingVelocities(xa, xb int) []*velocity {
	velocities := []*velocity{}
	// maximum possible arcing speed; i.e. 2 steps at least
	arcMax := xb/2 + 1
	for i := arcMax; i > 0; i-- {
		total := 0
		vel := i
		step := 0
		for total < xa {
			if vel == 0 {
				break
			}
			total += vel
			step++
			vel--
		}
		for total <= xb {
			if vel == 0 {
				break
			}
			if total <= xb && total >= xa {
				velocities = append(velocities, &velocity{
					x: i,
					t: step,
				})
			}
			total += vel
			step++
			vel--
		}

	}
	return velocities
}

func findXStoppedVelocities(xa, xb int) []int {
	xMin := int(math.Ceil(math.Sqrt(1+8*float64(xa))/2 - 0.5))
	xMax := int(math.Trunc(math.Sqrt(1+8*float64(xb))/2 - 0.5))
	velocities := make([]int, xMax-xMin+1)
	for i := xMin; i <= xMax; i++ {
		velocities[i-xMin] = i
	}
	return velocities
}

func findYVelocities(ya, yb int) ([]int, []*velocity, []int) {
	singleStep := make([]int, yb-ya+1)
	for y := ya; y <= yb; y++ {
		singleStep[y-ya] = -y
	}
	singleStepWayDown := make([]int, yb-ya+1)
	for y := ya; y <= yb; y++ {
		singleStepWayDown[y-ya] = y - 1
	}
	arcing := findYArcingVelocities(ya, yb)
	return singleStep, arcing, singleStepWayDown
}

func findYArcingVelocities(ya, yb int) []*velocity {
	velocities := []*velocity{}
	// maximum possible arcing speed; i.e. 2 steps at least
	arcMax := yb/2 + 1
	for i := -arcMax; i <= arcMax; i++ {
		total := 0
		vel := i
		step := 0
		for total > -ya {
			total += vel
			step++
			vel--
		}
		for total >= -yb {
			if total <= -ya && total >= -yb {
				velocities = append(velocities, &velocity{
					y: i,
					t: step,
				})
			}
			total += vel
			step++
			vel--
		}
	}
	return velocities
}

func Challenge(path string) (int, int) {
	packets, err := utils.ReadStrings(path)
	if err != nil {
		log.Fatal(err)
	}
	a, b, err := Trickshot(packets[0])
	if err != nil {
		log.Fatal(err)
	}
	return a, b
}
