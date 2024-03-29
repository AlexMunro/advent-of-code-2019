package main

import (
	"testing"
)

type mockChecker struct {
	grid [][]bool
}

func (m mockChecker) check(x int, y int) bool {
	return m.grid[x][y]
}

func buildMockChecker(lines []string) mockChecker {
	grid := make([][]bool, len(lines[0]))
	for i := range grid {
		grid[i] = make([]bool, len(lines))
	}

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == '.' {
				grid[x][y] = false
			} else {
				grid[x][y] = true
			}
		}
	}

	return mockChecker{grid: grid}
}

func TestCountAffectedPoints(t *testing.T) {
	gridInput := []string{
		"#.........",
		".#........",
		"..##......",
		"...###....",
		"....###...",
		".....####.",
		"......####",
		"......####",
		".......###",
		"........##",
	}

	checker := buildMockChecker(gridInput)

	expected := 27
	result := countAffectedPoints(0, 0, 10, checker)

	if expected != result {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestNearestTractorBeamSquare(t *testing.T) {
	gridInput := []string{
		"#.......................................",
		".#......................................",
		"..##....................................",
		"...###..................................",
		"....###.................................",
		".....####...............................",
		"......#####.............................",
		"......######............................",
		".......#######..........................",
		"........########........................",
		".........#########......................",
		"..........#########.....................",
		"...........##########...................",
		"...........############.................",
		"............############................",
		".............#############..............",
		"..............##############............",
		"...............###############..........",
		"................###############.........",
		"................#################.......",
		".................########OOOOOOOOOO.....",
		"..................#######OOOOOOOOOO#....",
		"...................######OOOOOOOOOO###..",
		"....................#####OOOOOOOOOO#####",
		".....................####OOOOOOOOOO#####",
		".....................####OOOOOOOOOO#####",
		"......................###OOOOOOOOOO#####",
		".......................##OOOOOOOOOO#####",
		"........................#OOOOOOOOOO#####",
		".........................OOOOOOOOOO#####",
		"..........................##############",
		"..........................##############",
		"...........................#############",
		"............................############",
		".............................###########",
	}

	checker := buildMockChecker(gridInput)

	expectedX := 25
	expectedY := 20

	actualX, actualY := nearestTractorBeamSquare(checker, 10)

	if expectedX != actualX || expectedY != actualY {
		t.Errorf("Expected (%v, %v) but got (%v, %v)", expectedX, expectedY, actualX, actualY)
	}
}
