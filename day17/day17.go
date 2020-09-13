package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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

func robotPosAndOrientation(image [][]rune) (int, int, rune) {
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			current := image[i][j]
			for _, dir := range []rune{'^', 'v', '<', '>', 'X'} {
				if current == dir {
					return i, j, dir
				}
			}
		}
	}

	panic("Robot implausibly not found!")
}

// Use Dijkstra's algorithm to find the shortest path
func findPath(image [][]rune) string {
	y, x, orientation := robotPosAndOrientation(image)
	path := ""

	for true {
		var dir rune

		switch orientation {
		case '^':
			if x > 0 && image[y][x-1] == '#' {
				dir = 'L'
				orientation = '<'
			} else if x < len(image[y]) && image[y][x+1] == '#' {
				dir = 'R'
				orientation = '>'
			} else {
				return path
			}
		case 'v':
			if x > 0 && image[y][x-1] == '#' {
				dir = 'R'
				orientation = '<'
			} else if x < len(image[y]) && image[y][x+1] == '#' {
				dir = 'L'
				orientation = '>'
			} else {
				return path
			}
		case '<':
			if y > 0 && image[y-1][x] == '#' {
				dir = 'R'
				orientation = '^'
			} else if y < len(image) && image[y+1][x] == '#' {
				dir = 'L'
				orientation = 'v'
			} else {
				return path
			}
		case '>':
			if y > 0 && image[y-1][x] == '#' {
				dir = 'L'
				orientation = '^'
			} else if y < len(image) && image[y+1][x] == '#' {
				dir = 'R'
				orientation = 'v'
			} else {
				return path
			}
		}
		dist := 0

		switch orientation {
		case '^':
			for y > 0 && image[y-1][x] == '#' {
				y--
				dist++
			}
		case 'v':
			for y < len(image)-1 && image[y+1][x] == '#' {
				y++
				dist++
			}
		case '<':
			for x > 0 && image[y][x-1] == '#' {
				x--
				dist++
			}
		case '>':
			for x < len(image[y])-1 && image[y][x+1] == '#' {
				x++
				dist++
			}
		}

		path += string(dir)
		path += ","
		path += strconv.Itoa(dist)
		path += ","
	}

	panic("I'm literally just doing this to shut the compiler up. The above loop is infinite for goodness' sake!")
}

// Very naive check of whether a string is a substring of an entire slice
func isGlobalSubstring(s string, strs []string) bool {
	for _, str := range strs {
		if len(strings.ReplaceAll(str, s, "")) != 0 {
			return false
		}
	}

	return true
}

func generateInput(image [][]rune) string {
	path := findPath(image)

	// First group, then
	// (first group,)*
	// Second group then
	// (first || second group,)*
	// Third group then
	// (first || second || third,)*

	step := regexp.MustCompile("^[LR],\\d+,")

	for i := len(step.FindString(path)); i < 20 && i < len(path); {
		// Only bother checking at commas
		a := path[:i]

		startPoints := []int{i}
		remainingPath := path[i:]

		bStart := i

		for strings.HasPrefix(remainingPath, a) {
			bStart += len(a)
			startPoints = append(startPoints, bStart)
			remainingPath = remainingPath[len(a):]
		}

		for _, bStart = range startPoints {
			for j := bStart + len(step.FindString(path[bStart:])); j < bStart+20 && j < len(path); {
				b := path[bStart:j]

				if b == a {
					j += len(step.FindString(path[j:]))
					continue
				}

				cStartPoints := []int{j}
				cStart := j

				remainingPath := path[j:]

				// Ignoring the edge case where A and B are both prefixes because lazy
				for strings.HasPrefix(remainingPath, a) || strings.HasPrefix(remainingPath, b) {
					if strings.HasPrefix(remainingPath, a) {
						cStart += len(a)
						remainingPath = remainingPath[len(a):]
					} else {
						cStart += len(b)
						remainingPath = remainingPath[len(b):]
					}
					cStartPoints = append(cStartPoints, cStart)
				}

				for _, cStart = range cStartPoints {
					for k := cStart + len(step.FindString(path[cStart:])); k < cStart+20 && k < len(path); {
						c := path[cStart:k]

						if c == a || c == b {
							k += len(step.FindString(path[k:]))
							continue
						}

						confirmationPath := path
						order := ""

						// Again, ignoring potential edge cases with multiple valid substitutions
						for strings.HasPrefix(confirmationPath, a) || strings.HasPrefix(confirmationPath, b) || strings.HasPrefix(confirmationPath, c) {
							if strings.HasPrefix(confirmationPath, a) {
								confirmationPath = confirmationPath[len(a):]
								order += "A,"
							}

							if strings.HasPrefix(confirmationPath, b) {
								confirmationPath = confirmationPath[len(b):]
								order += "B,"
							}

							if strings.HasPrefix(confirmationPath, c) {
								confirmationPath = confirmationPath[len(c):]
								order += "C,"
							}

							if len(confirmationPath) == 0 {
								return order[:len(order)-1] + "\n" + a[:len(a)-1] + "\n" + b[:len(b)-1] + "\n" + c[:len(c)-1] + "\n" + "y" + "\n"
							}
						}
						k += len(step.FindString(path[k:]))
					}
				}
				j += len(step.FindString(path[j:]))
			}
		}
		i += len(step.FindString(path[i:]))
	}

	panic("Failed to find valid A, B and C")
}

func guideRobot(program []int, image [][]rune) int {
	input := make(chan int)
	output := make(chan int)

	code := generateInput(image)
	codeInts := make([]int, len(code))

	for i, c := range code {
		codeInts[i] = int(rune(c))
	}

	go intcode.ExecuteProgram(program, codeInts, intcode.Channels{Input: input, Output: output})

	var result int

	for result < 255 {
		result = <-output
	}

	return result
}

func main() {
	input := utils.GetCommaSeparatedInts("input.txt")

	cameraImage, _ := intcode.ExecuteProgram(utils.CopyInts(input), []int{}, intcode.Channels{})

	parsedImage := parseCameraImage(cameraImage)

	firstAnswer := alignmentParemeterSum(parsedImage)

	fmt.Printf("The answer to part one is %d\n", firstAnswer)

	input[0] = 2

	secondAnswer := guideRobot(input, parsedImage)
	fmt.Printf("The answer to part two is %v\n", secondAnswer)
}
