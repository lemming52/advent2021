package beaconscanner

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

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
	rx, ry, rz  int
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

func locateScanner(origin, target *Scanner) {
	originIndex, targetIndex := 0, 0
	var first, second, third []*Displacement
	for {
		oDisplacement := origin.spacings[originIndex]
		tDisplacement := target.spacings[targetIndex]
		if oDisplacement.magnitude == tDisplacement.magnitude {
			first = []*Displacement{oDisplacement, tDisplacement}
			break
		}
		if oDisplacement.magnitude < tDisplacement.magnitude {
			originIndex++
		} else {
			targetIndex++
		}
	}
	originIndex++
	targetIndex++
	var sharedBeacon *Beacon
	for {
		oDisplacement := origin.spacings[originIndex]
		sharedBeacon = sharedPoint(first[0], oDisplacement)
		if sharedBeacon == nil {
			originIndex++
			continue
		}
		tDisplacement := target.spacings[targetIndex]
		if oDisplacement.magnitude == tDisplacement.magnitude {
			second = []*Displacement{oDisplacement, tDisplacement}
			break
		}
		if oDisplacement.magnitude < tDisplacement.magnitude {
			originIndex++
		} else {
			targetIndex++
		}
	}
	thirdDisplacement := magnitude(first[0].otherEnd(sharedBeacon), second[0].otherEnd(sharedBeacon))
	for _, delta := range target.spacings {
		if delta.magnitude = thirdDisplacement.magnitude {
			third = []*Displacement{thirdDisplacement, delta}
		}
	}
	identifyTransform(origin, target, first, second, third)
}

func identifyTransform(origin, target *Scanner, f, s, t []*Displacement) {
	originA, targetA := sharedPoint(f[0], s[0]), sharedPoint(f[1], s[1])
	originB, targetB := sharedPoint(f[0], t[0]), sharedPoint(f[1], t[1])
	originA, targetA := sharedPoint(s[0], t[0]), sharedPoint(s[1], t[1])


}

func ScanBeacons(scannedBeacons [][]string) (int, int, error) {
	origin, err := newScanner(0, scannedBeacons[0])
	if err != nil {
		return 0, 0, nil
	}
	scanners := make([]*Scanner, len(scannedBeacons)-1)
	for i, beacons := range scannedBeacons[1:] {
		scan, err := newScanner(i+1, beacons)
		if err != nil {
			return 0, 0, err
		}
		scanners[i] = scan
	}
	for _, scanner := range scanners {
		locateScanner(origin, scanner)
	}
	for _, d := range origin.spacings {
		fmt.Println(d.magnitude, d.a.name, d.b.name)
	}
	return 0, 0, nil
}

/*
func Challenge(path string) (int, int) {
	packets, err := utils.ReadStrings(path)
	if err != nil {
		log.Fatal(err)
	}
	a, b, err := ScanBeacons(packets[0])
	if err != nil {
		log.Fatal(err)
	}
	return a, b
}
*/
