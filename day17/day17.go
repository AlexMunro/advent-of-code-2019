package main

import (
	"fmt"

	"../intcode"
	"../utils"
)

func parseCameraImage(input []int) [][]rune {
	lines := [][]rune{[]rune{}}

	for _, char := range input[:len(input)-2] { // Filter out extra newlines
		nextRune := rune(char)
		if nextRune == '\n' {
			lines = append(lines, []rune{})
		} else {
			lines[len(lines)-1] = append(lines[len(lines)-1], nextRune)
		}
	}

	return lines
}

func alignmentParemeterSum(image [][]rune) int {
	sum := 0

	isIntersection := func(x int, y int) bool {
		return image[x][y] == '#' &&
			image[x-1][y] == '#' && image[x+1][y] == '#' &&
			image[x][y-1] == '#' && image[x][y+1] == '#'
	}

	for i := 1; i < len(image)-1; i++ {
		for j := 1; j < len(image[i])-1; j++ {
			if isIntersection(i, j) {
				sum += i * j
			}
		}
	}

	return sum
}

func main() {
	input := utils.GetCommaSeparatedInts("input.txt")

	cameraImage, _ := intcode.ExecuteProgram(utils.CopyInts(input), []int{}, intcode.Channels{})

	parsedImage := parseCameraImage(cameraImage)

	firstAnswer := alignmentParemeterSum(parsedImage)

	fmt.Printf("The answer to part one is %d\n", firstAnswer)
}
