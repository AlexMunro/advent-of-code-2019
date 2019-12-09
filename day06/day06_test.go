package main

import (
	"testing"
)

func TestOrbitCount(t *testing.T) {
	examples := map[*[]string]int{
		&([]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}): 42,
	}

	for k, v := range examples {
		result := countOrbits(buildDirectOrbiterMap(*k))
		if result != v {
			t.Errorf("Expected to get %d from %v but got %d", v, *k, result)
		}
	}
}

func TestDistToSanta(t *testing.T) {
	examples := map[*[]string]int{
		&([]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}): 4,
	}

	for k, v := range examples {
		result := distToSanta(buildDirectOrbitsMap(*k))
		if result != v {
			t.Errorf("Expected to get %d from %v but got %d", v, *k, result)
		}
	}
}
