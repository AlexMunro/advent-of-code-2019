package main

import "testing"

type mockChecker struct {
	grid [][]int
}

func (m mockChecker) check(x int, y int) int {
	return m.grid[y][x]
}

func buildMockChecker(lines []string) mockChecker {
	grid := make([][]int, len(lines[0]))
	for i := range grid {
		grid[i] = make([]int, len(lines))
	}

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == '.' {
				grid[x][y] = 0
			} else {
				grid[x][y] = 1
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

	actualX, actualY := nearestTractorBeamSquare(checker)

	if expectedX != actualX || expectedY != actualY {
		t.Errorf("Expected (%v, %v) but got (%v, %v)", expectedX, expectedY, actualX, actualY)
	}
}
