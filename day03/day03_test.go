package main

import (
	"testing"
)

func TestNearestIntersectionDistance(t *testing.T) {
	examples := map[*[]string]int{
		&([]string{"R8,U5,L5,D3", "U7,R6,D4,L4"}):                                                          6,
		&([]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}):               159,
		&([]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}): 135,
	}

	for k, v := range examples {
		result := nearestIntersectionDistance((*k)[0], (*k)[1])
		if result != v {
			t.Errorf("Expected to get %d from %v but got %d", v, *k, result)
		}
	}
}

func TestFewestIntersectionSteps(t *testing.T) {
	examples := map[*[]string]int{
		&([]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}):               610,
		&([]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}): 410,
	}

	for k, v := range examples {
		result := fewestIntersectionSteps((*k)[0], (*k)[1])
		if result != v {
			t.Errorf("Expected to get %d from %v but got %d", v, *k, result)
		}
	}
}
