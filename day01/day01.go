package main

import (
	"fmt"
	"path/filepath"

	"../utils"
)

func fuelForMass(mass int) int {
	return (mass / 3) - 2
}

func stableFuelForMass(mass int) int {
	totalFuel := 0
	remainingMass := mass
	for fuelForMass(remainingMass) > 0 {
		nextFuel := fuelForMass(remainingMass)
		remainingMass = nextFuel
		totalFuel += nextFuel
	}
	return totalFuel
}

func sumFuelForMass(masses []int, fuelFormula func(int) int) int {
	sum := 0
	for _, m := range masses {
		sum += fuelFormula(m)
	}
	return sum
}

func main() {
	filename, _ := filepath.Abs("./input.txt")
	input := utils.GetInputInts(filename)
	fmt.Printf("The answer to part one is %d\n", sumFuelForMass(input, fuelForMass))
	fmt.Printf("The answer to part two is %d\n", sumFuelForMass(input, stableFuelForMass))
}
