package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"../utils"
	"../utils/intset"
)

type location struct {
	x int
	y int
}

// Retrieves or initialises an *IntSet from a map of *IntSet
func getOrCreateIntSet(intSetMap map[int]*intset.IntSet, key int) *intset.IntSet {
	if intSetMap[key] == nil {
		intSetMap[key] = intset.New()
	}
	return intSetMap[key]
}

func intersectionPoints(path1, path2 string) []location {
	x, y := 0, 0
	points := map[int]*intset.IntSet{}

	// Build up maps from the first path that can be quickly checked
	for _, segment := range strings.Split(path1, ",") {
		dir := segment[0]
		dist, _ := strconv.Atoi(segment[1:])

		switch dir {
		case 'L':
			for i := 0; i < dist; i++ {
				getOrCreateIntSet(points, x-i).Add(y)
			}
			x -= dist
		case 'R':
			for i := 0; i < dist; i++ {
				getOrCreateIntSet(points, x+i).Add(y)
			}
			x += dist
		case 'D':
			for i := 0; i < dist; i++ {
				getOrCreateIntSet(points, x).Add(y - i)
			}
			y -= dist
		case 'U':
			for i := 0; i < dist; i++ {
				getOrCreateIntSet(points, x).Add(y + i)
			}
			y += dist
		}
	}

	x, y = 0, 0
	intersections := []location{}

	// For each point in the second path, check for any intersections
	for _, segment := range strings.Split(path2, ",") {
		dir := segment[0]
		dist, _ := strconv.Atoi(segment[1:])

		switch dir {
		case 'L':
			for i := 0; i < dist; i++ {
				if points[x-i].Contains(y) {
					intersections = append(intersections, location{x - i, y})
				}
			}
			x -= dist
		case 'R':
			for i := 0; i < dist; i++ {
				if points[x+i].Contains(y) {
					intersections = append(intersections, location{x + i, y})
				}
			}
			x += dist
		case 'D':
			for i := 0; i < dist; i++ {
				if points[x].Contains(y - i) {
					intersections = append(intersections, location{x, y - i})
				}
			}
			y -= dist
		case 'U':
			for i := 0; i < dist; i++ {
				if points[x].Contains(y + i) {
					intersections = append(intersections, location{x, y + i})
				}
			}
			y += dist
		}
	}
	return intersections
}

func manhattanDist(first, second location) int {
	return utils.Abs(first.x-second.x) + utils.Abs(first.y-second.y)
}

func nearestIntersectionDistance(firstPath, secondPath string) int {
	origin := location{}
	intersections := intersectionPoints(firstPath, secondPath)[1:]

	minDist := manhattanDist(origin, intersections[0])
	for _, point := range intersections {
		nextDist := manhattanDist(origin, point)
		if nextDist < minDist {
			minDist = nextDist
		}
	}
	return minDist
}

func stepsToPoint(path string, point location) int {
	x, y, steps := 0, 0, 0

	for _, segment := range strings.Split(path, ",") {
		dir := segment[0]
		dist, _ := strconv.Atoi(segment[1:])

		switch dir {
		case 'L':
			if point.y == y && point.x < x && point.x >= x-dist {
				return steps - point.x + x
			}
			x -= dist
		case 'R':
			if point.y == y && point.x > x && point.x <= x+dist {
				return steps + point.x - x
			}
			x += dist
		case 'D':
			if point.x == x && point.y < y && point.y >= y-dist {
				return steps - point.y + y
			}
			y -= dist
		case 'U':
			if point.x == x && point.y > y && point.y <= y+dist {
				return steps + point.y - y
			}
			y += dist
		}
		steps += dist
	}
	panic("Did not encounter the point")
}

func fewestIntersectionSteps(firstPath, secondPath string) int {
	intersections := intersectionPoints(firstPath, secondPath)[1:]

	minSteps := stepsToPoint(firstPath, intersections[0]) + stepsToPoint(secondPath, intersections[0])
	for _, point := range intersections[1:] {
		nextSteps := stepsToPoint(firstPath, point) + stepsToPoint(secondPath, point)
		if nextSteps < minSteps {
			minSteps = nextSteps
		}
	}
	return minSteps
}

func main() {
	filename, _ := filepath.Abs("./input.txt")
	paths := utils.GetInputLines(filename)

	nearestIntersection := nearestIntersectionDistance(paths[0], paths[1])
	fmt.Printf("The answer for part one is %d\n", nearestIntersection)

	fewestIntersectionSteps := fewestIntersectionSteps(paths[0], paths[1])
	fmt.Printf("The answer for part two is %d\n", fewestIntersectionSteps)
}
