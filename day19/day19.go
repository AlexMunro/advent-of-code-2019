package main

import (
	"fmt"

	"../intcode"
	"../utils"
)

func countAffectedPoints(program []int) int {
	count := 0

	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			output := make(chan int)
			// Program halts after one input, so we have to re-execute every time
			go intcode.ExecuteProgram(program, []int{x, y}, intcode.Channels{Output: output})
			count += <-output
		}
	}

	return count
}

func main() {
	input := utils.GetCommaSeparatedInts("input.txt")
	affectedPoints := countAffectedPoints(input)
	fmt.Printf("The solution to part one is: %v\n", affectedPoints)
}
