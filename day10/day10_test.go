package main

import (
	"testing"

	. "../utils/location"
)

var exampleInput []string = []string{
	".#..##.###...#######",
	"##.############..##.",
	".#.######.########.#",
	".###.#######.####.#.",
	"#####.##.#.##.###.##",
	"..#####..#.#########",
	"####################",
	"#.####....###.#.#.##",
	"##.#################",
	"#####.##.###..####..",
	"..######..##.#######",
	"####.##.####...##..#",
	".#####..#.######.###",
	"##...#.##########...",
	"#.##########.#######",
	".####.#.###.###.#.##",
	"....##.##.###..#####",
	".#.#.###########.###",
	"#.#.#.#####.####.###",
	"###.##.####.##.#..##",
}

var height, width = len(exampleInput), len(exampleInput[0])

var testAsteroids []Location = parseAsteroids(exampleInput, height, width)

func TestMostVisibleAsteroid(t *testing.T) {
	_, result := mostVisibleAsteroid(testAsteroids, height, width)
	if result != 210 {
		t.Errorf("Expected to get 210 but got %d", result)
	}
}

func TestNthDestroyedAsteroid(t *testing.T) {
	examples := map[int]Location{
		1:   Location{11, 12},
		2:   Location{12, 1},
		3:   Location{12, 2},
		10:  Location{12, 8},
		20:  Location{16, 0},
		50:  Location{16, 9},
		100: Location{10, 16},
		199: Location{9, 6},
		200: Location{8, 2},
		201: Location{10, 9},
		299: Location{11, 1},
	}

	for k, v := range examples {
		result := nthDestroyedAsteroid(Location{11, 13}, testAsteroids, height, width, k)
		if result != v {
			t.Errorf("Expected asteroid %d to be %d but got %d", k, v, result)
		}
	}
}
