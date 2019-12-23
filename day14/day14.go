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

// Returns the ore needed to make fuel
func oreNeeded(recipes map[string]recipe) int {
	remainingIngredients := map[string]int{}
	surplus := map[string]int{}

	for _, reagent := range recipes["FUEL"].in {
		remainingIngredients[reagent.material] = reagent.quantity
	}

	onlyOreLeft := false

	if len(remainingIngredients) == 1 {
		_, onlyOreLeft = remainingIngredients["ORE"]
	}

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
	return remainingIngredients["ORE"]
}

func main() {
	input := utils.GetInputLines("input.txt")

	recipes := buildRecipeMap(input)
	answer := oreNeeded(recipes)

	fmt.Printf("The answer to part one is %d\n", answer)
}
