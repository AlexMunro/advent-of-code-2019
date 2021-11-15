package main

import (
	"fmt"

	"../intcode"
	"../utils"
)

// Allows us to replace the intcode computer with hard-coded responses in tests
type checker interface {
	check(x int, y int) bool
}

type intcodeChecker struct {
	program []int
	cache   map[int]map[int]bool
}

func intcodeCheckerWithCache(program []int) intcodeChecker {
	return intcodeChecker{program: program, cache: map[int]map[int]bool{}}
}

func (i intcodeChecker) check(x int, y int) bool {
	if cachedLine, present := i.cache[x]; present {
		if cachedVal, present := cachedLine[y]; present {
			return cachedVal
		}
	}

	// Program halts after one input, so we have to re-execute every time
	output := make(chan int)
	go intcode.ExecuteProgram(i.program, []int{x, y}, intcode.Channels{Output: output})

	if _, present := i.cache[x]; !present {
		i.cache[x] = map[int]bool{}
	}

	i.cache[x][y] = <-output == 1

	return i.cache[x][y]
}

func countAffectedPoints(startX int, startY int, size int, checker checker) int {
	count := 0

	for x := startX; x < startX+size; x++ {
		for y := startY; y < startY+size; y++ {
			if checker.check(x, y) {
				count++
			}
		}
	}

	return count
}

func validRow(startX int, y int, length int, checker checker) bool {
	return checker.check(startX, y) && checker.check(startX+length-1, y)
}

func validGrid(startX int, startY int, size int, checker checker) bool {
	return validRow(startX, startY, size, checker) && validRow(startX, startY+size-1, size, checker)
}

func nearestTractorBeamSquare(checker checker, size int) (int, int) {
	currentX, currentY := 1, 1

	for {
		if validGrid(currentX, currentY, size, checker) {
			return currentX, currentY
		} else {
			// Move along in the same row if there is room to do so
			if checker.check(currentX+size, currentY) {
				currentX++
			} else { // Else, move down once and along until a checked point
				currentY++
				for !checker.check(currentX, currentY) {
					currentX++
				}
			}
		}
	}
}

func main() {
	input := utils.GetCommaSeparatedInts("input.txt")
	checker := intcodeCheckerWithCache(input)

	affectedPoints := countAffectedPoints(0, 0, 50, checker)
	fmt.Printf("The solution to part one is: %v\n", affectedPoints)

	nearestX, nearestY := nearestTractorBeamSquare(checker, 100)
	fmt.Printf("The solution to part two is: %v \n", nearestX*10_000+nearestY)
}
