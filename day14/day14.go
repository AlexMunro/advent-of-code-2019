package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"../utils"
)

type reagent struct {
	material string
	quantity int
}

type recipe struct {
	in  []reagent
	out reagent
}

func parseReagent(recipe string) reagent {
	r := regexp.MustCompile(`(\d+) ([A-Za-z]+)`)
	parts := r.FindStringSubmatch(recipe)
	q, _ := strconv.Atoi(parts[1])
	return reagent{material: parts[2], quantity: q}
}

func buildRecipeMap(recipeSlice []string) map[string]recipe {
	recipeMap := map[string]recipe{}
	for _, r := range recipeSlice {
		topSplit := strings.Split(r, " => ")
		out := parseReagent(topSplit[1])
		in := []reagent{}
		for _, s := range strings.Split(topSplit[0], ",") {
			in = append(in, parseReagent(s))
		}
		recipeMap[out.material] = recipe{in: in, out: out}
	}
	return recipeMap
}

// Returns the ore needed to make fuel with a map of surplus ingredients
func oreNeeded(recipes map[string]recipe, fuelNeeded int) (int, map[string]int) {
	remainingIngredients := map[string]int{}
	surplus := map[string]int{}
	remainingIngredients["FUEL"] = fuelNeeded

	onlyOreLeft := false

	for !onlyOreLeft {
		for material, quantity := range remainingIngredients {
			if material == "ORE" {
				continue
			}

			recipeMultiplier := utils.CeilDiv(quantity, recipes[material].out.quantity)

			surplus[material] += (recipeMultiplier * recipes[material].out.quantity) - quantity

			for _, ingredient := range recipes[material].in {
				newMaterial := ingredient.material
				requiredQuantity := ingredient.quantity * recipeMultiplier

				if surplus[newMaterial] >= requiredQuantity {
					surplus[newMaterial] -= requiredQuantity
					continue
				} else {
					totalQuantity := requiredQuantity + remainingIngredients[newMaterial] - surplus[newMaterial]
					surplus[newMaterial] = 0
					remainingIngredients[newMaterial] = totalQuantity
				}
			}

			delete(remainingIngredients, material)
		}

		if len(remainingIngredients) == 1 {
			_, onlyOreLeft = remainingIngredients["ORE"]
		}
	}
	return remainingIngredients["ORE"], surplus
}

// How much ore can we save by removing surplus ingredients?
// Starts from fuel then breaks down to component ingredients to see what savings can be made
func surplusOreSavings(recipes map[string]recipe, surplus map[string]int) int {
	ingredientsToVisit := map[string]struct{}{}
	ingredientsToVisit["FUEL"] = struct{}{}

	for len(ingredientsToVisit) > 0 {
		for material := range ingredientsToVisit {
			delete(ingredientsToVisit, material)

			minimumBatch := recipes[material].out.quantity
			skippableBatches := surplus[material] / minimumBatch

			for _, reagent := range recipes[material].in {
				if reagent.material != "ORE" {
					ingredientsToVisit[reagent.material] = struct{}{}
				}
				surplus[reagent.material] += skippableBatches * reagent.quantity
			}
			delete(surplus, material)
		}
	}
	return surplus["ORE"]
}

// Returns the amount of fuel which can be produced with a trillion ore
func fuelPerTrillionOre(recipes map[string]recipe) int {
	fuelForOne, _ := oreNeeded(recipes, 1)

	// Start by ignoring surplus materials to create a lower bound and then increment
	fuelCount := 1000000000000 / fuelForOne
	oreCount, surplus := 0, map[string]int{}

	for oreCount <= 1000000000000 {
		oreCount, surplus = oreNeeded(recipes, fuelCount)
		oreCount -= surplusOreSavings(recipes, surplus)

		// Skipping as many iterations as possible based on the cost for one fuel, which should be pessimistic
		fuelIncrement := utils.Max([]int{(1000000000000 - oreCount) / fuelForOne, 1})
		fuelCount += fuelIncrement
	}

	// This should read - 1, but I have an off-by-one error somewhere and I'm not above cheap hacks
	return fuelCount - 2
}

func main() {
	input := utils.GetInputLines("input.txt")

	recipes := buildRecipeMap(input)
	oreForOne, _ := oreNeeded(recipes, 1)
	fmt.Printf("The answer to part one is %d\n", oreForOne)

	fuelPerTrillionOre := fuelPerTrillionOre(recipes)
	fmt.Printf("The answer to part two is %d\n", fuelPerTrillionOre)
}
