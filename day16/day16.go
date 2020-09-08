package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"../utils"
)

func basePattern() []int {
	return []int{0, 1, 0, -1}
}

var memoizedPatterns map[int][]int = map[int][]int{}

// Generates pattern for the nth position - does not skip the first entry
func pattern(pos int) []int {
	base := basePattern()

	if memoizedPattern, present := memoizedPatterns[pos]; present {
		return memoizedPattern
	}

	output := make([]int, len(base)*(pos+1))

	for i, n := range base {
		for j := 0; j < pos+1; j++ {
			output[i*(pos+1)+j] = n
		}
	}

	memoizedPatterns[pos] = output
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

func parseInts(signal []int) int {
	total := 0
	for i, n := range signal {
		total += n * int(math.Pow10(len(signal)-i-1))
	}
	return total
}

func findMessage(signal []int) []int {
	messageLoc := parseInts(signal[:7])

	if messageLoc < len(signal)/2 {
		panic("This hack only works if you're beyond half way through the input")
	}

	fullSignal := utils.RepeatInts(signal, 10_000)

	relevantSignal := fullSignal[messageLoc:]

	for phase := 0; phase < 100; phase++ {
		// Skip final element since it never changes
		for pos := len(relevantSignal) - 2; pos >= 0; pos-- {
			relevantSignal[pos] += relevantSignal[pos+1]
			relevantSignal[pos] %= 10
		}
	}

	return relevantSignal[:8]
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

	answer2 := findMessage(input)

	fmt.Print("The answer to part two is ")
	for _, i := range answer2 {
		fmt.Printf("%d", i)
	}
	fmt.Print("\n")
}
