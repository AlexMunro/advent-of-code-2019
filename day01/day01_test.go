package main

import (
	"testing"
)

func TestFuelForMass(t *testing.T) {
	examples := map[int]int{
		12:     2,
		14:     2,
		1969:   654,
		100756: 33583,
	}

	for k, v := range examples {
		if fuelForMass(k) != v {
			t.Errorf("Expected to need %d fuel for %d mass but calculated %d", v, k, fuelForMass(k))
		}
	}
}

func TestSumFuelForMass(t *testing.T) {
	example := []int{12, 14, 1969, 100756}
	expected := 2 + 2 + 654 + 33583

	answer := sumFuelForMass(example, fuelForMass)
	if answer != expected {

	}
}

func TestStableFuelForMass(t *testing.T) {
	examples := map[int]int{
		14:     2,
		1969:   966,
		100756: 50346,
	}

	for k, v := range examples {
		if stableFuelForMass(k) != v {
			t.Errorf("Expected to need %d fuel for %d mass but calculated %d", v, k, stableFuelForMass(k))
		}
	}
}
