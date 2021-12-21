package beaconscanner

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type transformFunc func(x, y, z int) (int, int, int)

type Beacon struct {
	x, y, z int
	name    string
}

type Displacement struct {
	magnitude int
	a, b      *Beacon
}

func sharedPoint(a, b *Displacement) *Beacon {
	if a.a == b.a {
		return a.a
	}
	if a.a == b.b {
		return a.a
	}
	if a.b == b.a {
		return a.b
	}
	if a.b == b.b {
		return a.b
	}
	return nil
}

func (d *Displacement) otherEnd(beacon *Beacon) *Beacon {
	if beacon == d.a {
		return d.b
	}
	if beacon == d.b {
		return d.a
	}
	return nil
}

func magnitude(a, b *Beacon) *Displacement {
	dx, dy, dz := (b.x - a.x), (b.y - a.y), (b.z - a.z)
	return &Displacement{
		magnitude: dx*dx + dy*dy + dz*dz,
		a:         a,
		b:         b,
	}
}

type Scanner struct {
	name        int
	x, y, z     int
	rotations   []transformFunc
	spacings    []*Displacement
	beacons     map[string]*Beacon
	beaconCount int
}

func newScanner(name int, beacons []string) (*Scanner, error) {
	scanner := &Scanner{
		name:     name,
		spacings: make([]*Displacement, len(beacons)),
		beacons:  make(map[string]*Beacon, len(beacons)),
	}
	for _, b := range beacons {
		beacon, err := parseBeacon(b)
		if err != nil {
			return nil, err
		}
		scanner.beaconCount++
		for _, delta := range scanner.beacons {
			scanner.addDisplacement(beacon, delta)
		}
		scanner.beacons[beacon.name] = beacon
	}
	return scanner, nil
}

func (s *Scanner) addDisplacement(a, b *Beacon) {
	mag := magnitude(a, b)
	if s.beaconCount == 2 {
		s.spacings[0] = mag
		return
	}
	sortIndex := sort.Search(s.beaconCount-2, func(i int) bool { return s.spacings[i].magnitude >= mag.magnitude })
	copy(s.spacings[sortIndex+1:], s.spacings[sortIndex:])
	s.spacings[sortIndex] = mag
}

func (s *Scanner) rotate(x, y, z int) (int, int, int) {
	for _, t := range s.rotations {
		x, y, z = t(x, y, z)
	}
	return x, y, z
}

func parseBeacon(s string) (*Beacon, error) {
	coords := strings.Split(s, ",")
	b := &Beacon{
		name: s,
	}
	for i, c := range coords {
		val, err := strconv.Atoi(c)
		if err != nil {
			return nil, err
		}
		switch i {
		case 0:
			b.x = val
		case 1:
			b.y = val
		case 2:
			b.z = val
		}
	}
	return b, nil
}

func locateScanner(origin, target *Scanner) bool {
	originIndex, targetIndex, maxO, maxT := 0, 0, len(origin.spacings), len(target.spacings)
	shared := map[string][]*Beacon{}
	candidates := map[string][]*Displacement{}
	for originIndex < maxO && targetIndex < maxT {
		if len(shared) >= 12 { // below this the results get weird
			findTargetOrigin(origin, target, shared)
			return true
		}
		oDisplacement := origin.spacings[originIndex]
		tDisplacement := target.spacings[targetIndex]
		if oDisplacement.magnitude == tDisplacement.magnitude {
			aShared, aFound := shared[oDisplacement.a.name]
			bShared, bFound := shared[oDisplacement.b.name]
			if aFound {
				if !bFound {
					shared[oDisplacement.b.name] = []*Beacon{oDisplacement.b, tDisplacement.otherEnd(aShared[0])}
				}
				originIndex++
				targetIndex++
				continue
			}
			if bFound {
				shared[oDisplacement.a.name] = []*Beacon{oDisplacement.a, tDisplacement.otherEnd(bShared[0])}
				originIndex++
				targetIndex++
				continue
			}
			aCandidate, aFound := candidates[oDisplacement.a.name]
			bCandidate, bFound := candidates[oDisplacement.b.name]
			if aFound {
				aPrime := sharedPoint(aCandidate[1], tDisplacement)
				shared[oDisplacement.a.name] = []*Beacon{oDisplacement.a, aPrime}
				shared[aCandidate[0].otherEnd(oDisplacement.a).name] = []*Beacon{aCandidate[0].otherEnd(oDisplacement.a), aCandidate[1].otherEnd(aPrime)}
				shared[oDisplacement.b.name] = []*Beacon{oDisplacement.b, tDisplacement.otherEnd(aPrime)}
				originIndex++
				targetIndex++
				continue
			}
			if bFound {
				bPrime := sharedPoint(bCandidate[1], tDisplacement)
				shared[oDisplacement.b.name] = []*Beacon{oDisplacement.b, bPrime}
				shared[bCandidate[0].otherEnd(oDisplacement.b).name] = []*Beacon{bCandidate[0].otherEnd(oDisplacement.b), bCandidate[1].otherEnd(bPrime)}
				shared[oDisplacement.a.name] = []*Beacon{oDisplacement.a, tDisplacement.otherEnd(bPrime)}
				originIndex++
				targetIndex++
				continue
			}
			candidates[oDisplacement.a.name] = []*Displacement{oDisplacement, tDisplacement}
			candidates[oDisplacement.b.name] = []*Displacement{oDisplacement, tDisplacement}
			originIndex++
			targetIndex++
			continue
		}
		if oDisplacement.magnitude < tDisplacement.magnitude {
			originIndex++
		} else {
			targetIndex++
		}
	}
	return false
}

func findTargetOrigin(origin, target *Scanner, shared map[string][]*Beacon) {
	transforms := allowedTransformations()
	origins := map[int][]int{}
	for _, v := range shared {
		if origins[0] == nil {
			for k, t := range transforms {
				rx, ry, rz := t(v[1].x, v[1].y, v[1].z)
				origins[k] = []int{rx - v[0].x, ry - v[0].y, rz - v[0].z}
			}
			continue
		}
		if len(transforms) == 1 {
			break
		}
		for k, t := range transforms {
			rx, ry, rz := t(v[1].x, v[1].y, v[1].z)
			if origins[k][0] != rx-v[0].x || origins[k][1] != ry-v[0].y || origins[k][2] != rz-v[0].z {
				delete(transforms, k)
			}
		}
	}
	for k, v := range transforms {
		originCoords := origins[k]
		x, y, z := origin.rotate(originCoords[0], originCoords[1], originCoords[2])
		target.x = origin.x + x
		target.y = origin.y + y
		target.z = origin.z + z
		target.rotations = append([]transformFunc{v}, origin.rotations...)
	}
}

func allowedTransformations() map[int]transformFunc {
	return map[int]transformFunc{
		0: func(x, y, z int) (int, int, int) {
			return x, y, z
		},
		1: func(x, y, z int) (int, int, int) {
			return x, -z, y
		},
		2: func(x, y, z int) (int, int, int) {
			return x, -y, -z
		},
		3: func(x, y, z int) (int, int, int) {
			return x, z, -y
		},
		4: func(x, y, z int) (int, int, int) {
			return -x, -y, z // - +
		},
		5: func(x, y, z int) (int, int, int) {
			return -x, -z, -y // - -
		},
		6: func(x, y, z int) (int, int, int) {
			return -x, y, -z // + -
		},
		7: func(x, y, z int) (int, int, int) {
			return -x, z, y // + +
		},
		8: func(x, y, z int) (int, int, int) {
			return y, -x, z
		},
		9: func(x, y, z int) (int, int, int) {
			return y, -z, -x
		},
		10: func(x, y, z int) (int, int, int) {
			return y, x, -z
		},
		11: func(x, y, z int) (int, int, int) {
			return y, z, x
		},
		12: func(x, y, z int) (int, int, int) {
			return -y, x, z
		},
		13: func(x, y, z int) (int, int, int) {
			return -y, -z, x
		},
		14: func(x, y, z int) (int, int, int) {
			return -y, -x, -z
		},
		15: func(x, y, z int) (int, int, int) {
			return -y, z, -x
		},
		16: func(x, y, z int) (int, int, int) {
			return z, y, -x
		},
		17: func(x, y, z int) (int, int, int) {
			return z, x, y
		},
		18: func(x, y, z int) (int, int, int) {
			return z, -y, x
		},
		19: func(x, y, z int) (int, int, int) {
			return z, -y, -x
		},
		20: func(x, y, z int) (int, int, int) {
			return -z, y, x
		},
		21: func(x, y, z int) (int, int, int) {
			return -z, -x, y
		},
		22: func(x, y, z int) (int, int, int) {
			return -z, -y, -x
		},
		23: func(x, y, z int) (int, int, int) {
			return -z, x, -y
		},
	}
}

func ScanBeacons(scannedBeacons [][]string) (int, int, error) {
	toLocate := make(map[int]*Scanner, len(scannedBeacons))
	located := map[int]*Scanner{}

	for i, beacons := range scannedBeacons {
		scan, err := newScanner(i, beacons)
		if err != nil {
			return 0, 0, err
		}
		toLocate[scan.name] = scan
	}
	located[0] = toLocate[0]
	delete(toLocate, 0)

	for len(toLocate) != 0 {
		for _, target := range toLocate {
			for _, origin := range located {
				if target.name == origin.name {
					continue
				}
				overlap := locateScanner(origin, target)

				if overlap {
					located[target.name] = target
					delete(toLocate, target.name)
					break
				}
			}
		}

	}
	beacons := map[string]bool{}
	maxManhattan := float64(0)
	for i, l := range located {
		for _, b := range l.beacons {
			x, y, z := l.rotate(b.x, b.y, b.z)
			beacons[fmt.Sprintf("%d,%d,%d", x-l.x, y-l.y, z-l.z)] = true
		}
		for j, k := range located {
			if i == j {
				continue
			}
			manhattan := math.Abs(float64(k.x-l.x)) + math.Abs(float64(k.y-l.y)) + math.Abs(float64(k.z-l.z))
			if manhattan > maxManhattan {
				maxManhattan = manhattan
			}
		}
	}
	return len(beacons), int(maxManhattan), nil
}

func Challenge(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanners := [][]string{}
	scan := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else if strings.HasPrefix(line, "---") {
			scanners = append(scanners, scan)
			scan = []string{}

		} else {
			scan = append(scan, line)
		}
	}
	scanners = append(scanners, scan)
	a, b, err := ScanBeacons(scanners[1:])
	if err != nil {
		log.Fatal(err)
	}
	return a, b
}
