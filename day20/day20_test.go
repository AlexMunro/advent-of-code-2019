package main

import (
	"reflect"
	"testing"

	"../utils/location"
)

func testMaze() []string {
	return []string{
		"         A           ",
		"         A           ",
		"  #######.#########  ",
		"  #######.........#  ",
		"  #######.#######.#  ",
		"  #######.#######.#  ",
		"  #######.#######.#  ",
		"  #####  B    ###.#  ",
		"BC...##  C    ###.#  ",
		"  ##.##       ###.#  ",
		"  ##...DE  F  ###.#  ",
		"  #####    G  ###.#  ",
		"  #########.#####.#  ",
		"DE..#######...###.#  ",
		"  #.#########.###.#  ",
		"FG..#########.....#  ",
		"  ###########.#####  ",
		"             Z       ",
		"             Z       ",
	}
}

func TestParsePointLocs(t *testing.T) {
	expected := map[location.Location]string{
		location.Location{X: 9, Y: 2}:   "AA",
		location.Location{X: 9, Y: 6}:   "BC",
		location.Location{X: 2, Y: 8}:   "BC",
		location.Location{X: 6, Y: 10}:  "DE",
		location.Location{X: 2, Y: 13}:  "DE",
		location.Location{X: 2, Y: 15}:  "FG",
		location.Location{X: 11, Y: 12}: "FG",
		location.Location{X: 13, Y: 16}: "ZZ",
	}

	actual := parsePointLocs(testMaze())

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestParseGraph(t *testing.T) {
	expected := graph{
		"AA": map[string]int{"BC": 4, "FG": 30, "ZZ": 26},
		"BC": map[string]int{"AA": 4, "DE": 6, "FG": 32, "ZZ": 28},
		"DE": map[string]int{"BC": 6, "FG": 4},
		"FG": map[string]int{"AA": 30, "BC": 32, "DE": 4, "ZZ": 6},
		"ZZ": map[string]int{"AA": 26, "BC": 28, "FG": 6},
	}
	actual := parseGraph(testMaze())

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func complexTestMaze() []string {
	return []string{
		"                   A               ",
		"                   A               ",
		"  #################.#############  ",
		"  #.#...#...................#.#.#  ",
		"  #.#.#.###.###.###.#########.#.#  ",
		"  #.#.#.......#...#.....#.#.#...#  ",
		"  #.#########.###.#####.#.#.###.#  ",
		"  #.............#.#.....#.......#  ",
		"  ###.###########.###.#####.#.#.#  ",
		"  #.....#        A   C    #.#.#.#  ",
		"  #######        S   P    #####.#  ",
		"  #.#...#                 #......VT",
		"  #.#.#.#                 #.#####  ",
		"  #...#.#               YN....#.#  ",
		"  #.###.#                 #####.#  ",
		"DI....#.#                 #.....#  ",
		"  #####.#                 #.###.#  ",
		"ZZ......#               QG....#..AS",
		"  ###.###                 #######  ",
		"JO..#.#.#                 #.....#  ",
		"  #.#.#.#                 ###.#.#  ",
		"  #...#..DI             BU....#..LF",
		"  #####.#                 #.#####  ",
		"YN......#               VT..#....QG",
		"  #.###.#                 #.###.#  ",
		"  #.#...#                 #.....#  ",
		"  ###.###    J L     J    #.#.###  ",
		"  #.....#    O F     P    #.#...#  ",
		"  #.###.#####.#.#####.#####.###.#  ",
		"  #...#.#.#...#.....#.....#.#...#  ",
		"  #.#####.###.###.#.#.#########.#  ",
		"  #...#.#.....#...#.#.#.#.....#.#  ",
		"  #.###.#####.###.###.#.#.#######  ",
		"  #.#.........#...#.............#  ",
		"  #########.###.###.#############  ",
		"           B   J   C               ",
		"           U   P   P               ",
	}
}

func TestAAToZZ(t *testing.T) {
	expected := 23
	actual := aaToZZ(parseGraph(testMaze()))

	if expected != actual {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}

	expected = 58
	actual = aaToZZ(parseGraph(complexTestMaze()))

	if expected != actual {
		t.Errorf("Expected %v but instead got %v", expected, actual)
	}
}
