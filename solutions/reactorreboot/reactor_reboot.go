package reactorreboot

import (
	"advent/solutions/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

const instructionPattern = `(\w{2,3}) x=(\-?\d{1,6})\.\.(\-?\d{1,6})\,y=(\-?\d{1,6})\.\.(\-?\d{1,6})\,z=(\-?\d{1,6})\.\.(\-?\d{1,6})`

type Cuboid struct {
	x0, x1, y0, y1, z0, z1 int
	on                     bool
}

func (c *Cuboid) volume() int {
	return (c.x1 + 1 - c.x0) * (c.y1 + 1 - c.y0) * (c.z1 + 1 - c.z0)
}

func (c *Cuboid) print() {
	for i := c.x0; i <= c.x1; i++ {
		for j := c.y0; j <= c.y1; j++ {
			for k := c.z0; k <= c.z1; k++ {
				fmt.Println("[", i, j, k, "]")
			}
		}
	}
}

func newCuboid(s string, pattern *regexp.Regexp) (*Cuboid, error) {
	components := pattern.FindStringSubmatch(s)
	values := make([]int, 6)
	for i, c := range components[2:] {
		val, err := strconv.Atoi(c)
		if err != nil {
			return nil, err
		}
		values[i] = val
	}
	c := &Cuboid{
		on: components[1] == "on",
	}
	if values[0] < values[1] {
		c.x0, c.x1 = values[0], values[1]
	} else {
		c.x0, c.x1 = values[1], values[0]
	}
	if values[2] < values[3] {
		c.y0, c.y1 = values[2], values[3]
	} else {
		c.y0, c.y1 = values[3], values[2]
	}
	if values[4] < values[5] {
		c.z0, c.z1 = values[4], values[5]
	} else {
		c.z0, c.z1 = values[5], values[4]
	}
	return c, nil
}

func cuboidsOverlap(a, b *Cuboid) bool {
	return (overlappingRange(a.x0, a.x1, b.x0, b.x1) &&
		overlappingRange(a.y0, a.y1, b.y0, b.y1) &&
		overlappingRange(a.z0, a.z1, b.z0, b.z1))
}

func overlappingRange(a, b, c, d int) bool {
	return (c <= a && a <= d) || (c <= b && b <= d) || (c >= a && d <= b)
}

func evaluateNewCube(new *Cuboid, existing []*Cuboid) []*Cuboid {
	newCuboids := []*Cuboid{}
	for _, old := range existing {
		if cuboidsOverlap(new, old) {
			additional := splitCuboidAroundNew(new, old)
			newCuboids = append(newCuboids, additional...)
		} else {
			newCuboids = append(newCuboids, old)
		}
	}
	if new.on {
		newCuboids = append(newCuboids, new)
	}
	return newCuboids
}

func splitCuboidAroundNew(new, old *Cuboid) []*Cuboid {
	components := []*Cuboid{}
	x0Corrected, x1Corrected := old.x0, old.x1
	y0Corrected, y1Corrected := old.y0, old.y1
	if old.x0 < new.x0 {
		components = append(components, &Cuboid{
			on: true,
			x0: old.x0,
			x1: new.x0 - 1,
			y0: old.y0,
			y1: old.y1,
			z0: old.z0,
			z1: old.z1,
		})
		x0Corrected = new.x0
	}
	if old.x1 > new.x1 {
		components = append(components, &Cuboid{
			on: true,
			x0: new.x1 + 1,
			x1: old.x1,
			y0: old.y0,
			y1: old.y1,
			z0: old.z0,
			z1: old.z1,
		})
		x1Corrected = new.x1
	}
	if old.y0 < new.y0 {
		components = append(components, &Cuboid{
			on: true,
			x0: x0Corrected,
			x1: x1Corrected,
			y0: old.y0,
			y1: new.y0 - 1,
			z0: old.z0,
			z1: old.z1,
		})
		y0Corrected = new.y0
	}
	if old.y1 > new.y1 {
		components = append(components, &Cuboid{
			on: true,
			x0: x0Corrected,
			x1: x1Corrected,
			y0: new.y1 + 1,
			y1: old.y1,
			z0: old.z0,
			z1: old.z1,
		})
		y1Corrected = new.y1
	}
	if old.z0 < new.z0 {
		components = append(components, &Cuboid{
			on: true,
			x0: x0Corrected,
			x1: x1Corrected,
			y0: y0Corrected,
			y1: y1Corrected,
			z0: old.z0,
			z1: new.z0 - 1,
		})
	}
	if old.z1 > new.z1 {
		components = append(components, &Cuboid{
			on: true,
			x0: x0Corrected,
			x1: x1Corrected,
			y0: y0Corrected,
			y1: y1Corrected,
			z0: new.z1 + 1,
			z1: old.z1,
		})
	}
	return components
}

func Reboot(instructions []string) (int, error) {
	pattern, err := regexp.Compile(instructionPattern)
	if err != nil {
		return 0, err
	}
	var onCuboids []*Cuboid
	for _, i := range instructions {
		cuboid, err := newCuboid(i, pattern)
		if err != nil {
			return 0, err
		}
		if onCuboids == nil {
			if cuboid.on {
				onCuboids = []*Cuboid{cuboid}
			}
			continue
		}
		onCuboids = evaluateNewCube(cuboid, onCuboids)
	}
	total := 0
	for _, i := range onCuboids {
		total += i.volume()
	}
	return total, nil
}

func Challenge(path string) (int, int) {
	lines, err := utils.ReadStrings(path)
	if err != nil {
		log.Fatal(err)
	}
	a, err := Reboot(lines[:20])
	if err != nil {
		log.Fatal(err)
	}
	b, err := Reboot(lines)
	if err != nil {
		log.Fatal(err)
	}
	return a, b
}
