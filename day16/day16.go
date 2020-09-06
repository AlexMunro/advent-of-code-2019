package main

import (
	"fmt"
	"strconv"
	"strings"

	"../utils"
)

func basePattern() []int {
	return []int{0, 1, 0, -1}
}

// Generates pattern for the nth position - does not skip the first entry
func pattern(pos int) []int {
	base := basePattern()

	output := make([]int, len(base)*(pos+1))

	for i, n := range base {
		for j := 0; j < pos+1; j++ {
			output[i*(pos+1)+j] = n
		}
	}

	return output
}

func nextFFSPhase(input []int) []int {
	output := make([]int, len(input))

	for i := range input {
		p := pattern(i)

		nextDigit := 0
		patternPos := 1
		for _, n := range input {
			nextDigit += p[patternPos%len(p)] * n
			patternPos++
		}
		output[i] = utils.Abs(nextDigit) % 10
	}

	return output
}

func abbreviatedFFS(input []int, phases int) []int {
	signal := utils.CopyInts(input)
	for i := 0; i < phases; i++ {
		signal = nextFFSPhase(signal)
	}
	return signal[:8]
}

func main() {
	rawInput := utils.GetInputSingleString("input.txt")
	input := make([]int, len(rawInput))
	for i, n := range strings.Split(rawInput, "") {
		input[i], _ = strconv.Atoi(n)
	}

	answer1 := abbreviatedFFS(input, 100)

	fmt.Print("The answer to part one is ")
	for _, i := range answer1 {
		fmt.Printf("%d", i)
	}
	fmt.Print("\n")
}
