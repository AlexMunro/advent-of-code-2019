package main

import (
	"fmt"

	"../intcode"
	"../utils"
)

// Allows us to replace the intcode computer with hard-coded responses in tests
type checker interface {
	check(x int, y int) int
}

type intcodeChecker struct {
	program []int
}

func (i intcodeChecker) check(x int, y int) int {
	// Program halts after one input, so we have to re-execute every time
	output := make(chan int)
	go intcode.ExecuteProgram(i.program, []int{x, y}, intcode.Channels{Output: output})
	return <-output
}

func countAffectedPoints(startX int, startY int, size int, checker checker) int {
	count := 0

	for x := startX; x < startX+size; x++ {
		for y := startY; y < startY+size; y++ {
			count += checker.check(x, y)
		}
	}

	return count
}

func nearestTractorBeamSquare(checker checker) (int, int) {
	return 0, 0
}

func main() {
	input := utils.GetCommaSeparatedInts("input.txt")
	affectedPoints := countAffectedPoints(0, 0, 50, intcodeChecker{program: input})
	fmt.Printf("The solution to part one is: %v\n", affectedPoints)
	fmt.Printf("The solution to part two is: \n")
}
