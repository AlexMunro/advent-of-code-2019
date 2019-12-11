package main

import (
	"fmt"
	"sort"

	"../utils"
	. "../utils/location"
)

type visibleLocations struct {
	contents map[Location]*LocationSet
}

func parseAsteroids(input []string, height, width int) []Location {
	asteroids := []Location{}
	for y, line := range input {
		for x, char := range line {
			if char == '#' {
				asteroids = append(asteroids, Location{x, y})
			}
		}
	}
	return asteroids
}

func newVisibleLocations() *visibleLocations {
	vl := visibleLocations{}
	vl.contents = map[Location]*LocationSet{}
	return &vl
}

func (vl *visibleLocations) addVisibleLocation(from, to Location) {
	if locs, present := vl.contents[from]; present {
		locs.AddLoc(to)
	} else {
		vl.contents[from] = New(to)
	}

	// This is -reciprocal
	if locs, present := vl.contents[to]; present {
		locs.AddLoc(from)
	} else {
		vl.contents[to] = New(from)
	}
}

func (vl *visibleLocations) visibleFrom(from, to Location) bool {
	if locs, present := vl.contents[from]; present {
		return locs.Contains(to)
	}
	return false
}

func containsLoc(locs []Location, loc Location) bool {
	for _, l := range locs {
		if l == loc {
			return true
		}
	}
	return false
}

func visibleAsteroids(asteroids []Location, height, width int) visibleLocations {
	visibleFrom := newVisibleLocations()
	notVisibleFrom := newVisibleLocations()

	for _, fromAsteroid := range asteroids {
		notVisibleFrom.addVisibleLocation(fromAsteroid, fromAsteroid)

		for _, toAsteroid := range asteroids {
			if notVisibleFrom.visibleFrom(fromAsteroid, toAsteroid) {
				continue
			}

			// Handling vertical line of sight separately because
			// I don't know how to divide by zero in Go

			if fromAsteroid.X == toAsteroid.X {
				foundAbove, foundBelow := false, false
				for y := fromAsteroid.Y - 1; y >= 0; y-- {
					loc := Location{fromAsteroid.X, y}
					if containsLoc(asteroids, loc) {
						if foundAbove {
							notVisibleFrom.addVisibleLocation(fromAsteroid, loc)
						} else {
							foundAbove = true
							visibleFrom.addVisibleLocation(fromAsteroid, loc)
						}
					}
				}

				for y := fromAsteroid.Y + 1; y < height; y++ {
					loc := Location{fromAsteroid.X, y}
					if containsLoc(asteroids, loc) {
						if foundBelow {
							notVisibleFrom.addVisibleLocation(fromAsteroid, loc)
						} else {
							foundBelow = true
							visibleFrom.addVisibleLocation(fromAsteroid, loc)
						}
					}
				}
			} else {
				// And now for all non-verticals. Yay gradients!
				gradient := Gradient(fromAsteroid, toAsteroid)
				foundLeft, foundRight := false, false

				for x := fromAsteroid.X - 1; x >= 0; x-- {
					if x == fromAsteroid.X {
						continue
					}

					// Can only cross an asteroid if y is an integer within bounds
					y := (gradient * float64(x-fromAsteroid.X)) + float64(fromAsteroid.Y)
					if y < 0 || int(y) >= height || y != float64(int(y)) {
						continue
					}

					loc := Location{x, int(y)}
					if containsLoc(asteroids, loc) {
						if foundLeft {
							notVisibleFrom.addVisibleLocation(fromAsteroid, loc)
						} else {
							foundLeft = true
							visibleFrom.addVisibleLocation(fromAsteroid, loc)
						}
					}
				}

				for x := fromAsteroid.X + 1; x < width; x++ {
					if x == fromAsteroid.X {
						continue
					}

					// Can only cross an asteroid if y is an integer within bounds
					y := (gradient * float64(x-fromAsteroid.X)) + float64(fromAsteroid.Y)
					if y < 0 || int(y) >= height || y != float64(int(y)) {
						continue
					}

					loc := Location{x, int(y)}
					if containsLoc(asteroids, loc) {
						if foundRight {
							notVisibleFrom.addVisibleLocation(fromAsteroid, loc)
						} else {
							foundRight = true
							visibleFrom.addVisibleLocation(fromAsteroid, loc)
						}
					}
				}
			}
		}
	}

	return *visibleFrom
}

// Most visible asteroid and the number of asteroids it can see
func mostVisibleAsteroid(asteroids []Location, height, width int) (Location, int) {
	linesOfSight := visibleAsteroids(asteroids, height, width)
	maxLines := 0
	var station Location
	for asteroid, visible := range linesOfSight.contents {
		if visible.Size() > maxLines {
			maxLines = visible.Size()
			station = asteroid
		}
	}
	return station, maxLines
}

func nthDestroyedAsteroid(centre Location, asteroids []Location, height, width, n int) Location {
	remainingAsteroids := n
	asteroidSet := FromSlice(asteroids)
	currentLayer := visibleAsteroids(asteroids, height, width).contents[centre]
	for remainingAsteroids-currentLayer.Size() > 0 {
		remainingAsteroids -= currentLayer.Size()
		asteroidSet.Difference(currentLayer)
		currentLayer = visibleAsteroids(asteroidSet.ToSlice(), height, width).contents[centre]
	}

	// There is probably a smarter way to sort these, but my high school trig has its limits

	above := currentLayer.Filter(func(loc Location) bool { return loc.X == centre.X && loc.Y < centre.Y })
	if remainingAsteroids > len(above) {
		remainingAsteroids -= len(above)
	} else {
		return above[remainingAsteroids-1]
	}

	right := currentLayer.Filter(func(loc Location) bool { return loc.X > centre.X })

	if remainingAsteroids > len(right) {
		remainingAsteroids -= len(right)
	} else {
		sort.Slice(right, func(i, j int) bool { return Gradient(right[i], centre) < Gradient(right[j], centre) })
		return right[remainingAsteroids-1]
	}

	below := currentLayer.Filter(func(loc Location) bool { return loc.X == centre.X && loc.Y > centre.Y })

	if remainingAsteroids > len(below) {
		remainingAsteroids -= len(below)
	} else {
		return below[remainingAsteroids-1]
	}

	left := currentLayer.Filter(func(loc Location) bool { return loc.X < centre.X })
	sort.Slice(left, func(i, j int) bool { return Gradient(left[i], centre) < Gradient(left[j], centre) })
	return left[remainingAsteroids-1]
}

func main() {
	input := utils.GetInputLines("input.txt")

	height := len(input)
	width := len(input[0])
	asteroids := parseAsteroids(input, height, width)

	station, maxLines := mostVisibleAsteroid(asteroids, height, width)
	fmt.Printf("The answer to part one is %d\n", maxLines)

	asteroid200 := nthDestroyedAsteroid(station, asteroids, height, width, 200)
	bet := (asteroid200.X * 100) + asteroid200.Y
	fmt.Printf("The answer to part two is %d\n", bet)
}
