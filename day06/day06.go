package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"../utils"
)

func buildDirectOrbiterMap(input []string) map[string][]string {
	directOrbiters := map[string][]string{}

	for _, line := range input {
		objects := strings.Split(line, ")")
		inner, outer := objects[0], objects[1]
		if orbits, present := directOrbiters[inner]; present {
			directOrbiters[inner] = append(orbits, outer)
		} else {
			directOrbiters[inner] = []string{outer}
		}
	}

	return directOrbiters
}

func buildDirectOrbitsMap(input []string) map[string]string {
	directOrbits := map[string]string{}

	for _, line := range input {
		objects := strings.Split(line, ")")
		inner, outer := objects[0], objects[1]
		directOrbits[outer] = inner
	}

	return directOrbits
}

func countOrbitsRec(object string, directOrbiters map[string][]string, orbitCounts map[string]int) int {
	if count, present := orbitCounts[object]; present {
		return count
	}

	var sum int

	if orbiters, present := directOrbiters[object]; present {
		sum = len(orbiters)
	} else {
		return 0
	}

	for _, outer := range directOrbiters[object] {
		sum += countOrbitsRec(outer, directOrbiters, orbitCounts)
	}

	orbitCounts[object] = sum
	return sum
}

func countOrbits(directOrbiters map[string][]string) int {
	sum := 0

	orbitCounts := map[string]int{}

	countOrbitsRec("COM", directOrbiters, orbitCounts)

	for _, orbiters := range orbitCounts {
		sum += orbiters
	}

	return sum
}

// Since this is a tree, we can go straight up from ME to COM and record how
// long it takes to get to each point. We can then go straight up from SAN,
// which should eventually meet the ME->COM path.
func distToSanta(directOrbits map[string]string) int {
	currentMe, currentSan := directOrbits["YOU"], directOrbits["SAN"]

	distanceTravelled := 0
	distsFromMe := map[string]int{}
	distsFromSan := map[string]int{}

	for true {
		if dist, present := distsFromSan[currentMe]; present {
			return distanceTravelled + dist
		}
		distsFromMe[currentMe] = distanceTravelled

		if dist, present := distsFromMe[currentSan]; present {
			return distanceTravelled + dist
		}
		distsFromSan[currentSan] = distanceTravelled
		distanceTravelled++
		currentMe, currentSan = directOrbits[currentMe], directOrbits[currentSan]
	}
	panic("This is only for the compiler's benefit.")
}

func main() {
	filename, _ := filepath.Abs("./input.txt")
	directOrbiters := buildDirectOrbiterMap(utils.GetInputLines(filename))
	orbitCount := countOrbits(directOrbiters)
	fmt.Printf("The answer to part one is %d\n", orbitCount)

	directOrbits := buildDirectOrbitsMap(utils.GetInputLines(filename))
	dist := distToSanta(directOrbits)
	fmt.Printf("The answer to part two is %d\n", dist)
}
