package main

import (
	"../utils/location"
)

type graph = map[string]map[string]int

func parsePointLocs(input []string) map[location.Location]string {
	pointLocs := map[location.Location]string{}

	visited := map[location.Location]struct{}{}

	for y, line := range input {
		for x, c := range line {
			if c < 'A' || c > 'Z' {
				continue
			}

			if _, present := visited[location.Location{X: x, Y: y}]; present {
				continue
			}

			visited[location.Location{X: x, Y: y}] = struct{}{}

			var point location.Location
			var otherChar rune

			// We know that the next letter will be either to the right or down from here
			if x < len(line)-1 && line[x+1] >= 'A' && line[x+1] <= 'Z' {
				otherChar = rune(line[x+1])
				visited[location.Location{X: x + 1, Y: y}] = struct{}{}

				// Establish whether the dot is to the right or left of the label
				if x > 0 && line[x-1] == '.' {
					point = location.Location{X: x - 1, Y: y}
				} else {
					point = location.Location{X: x + 2, Y: y}
				}
			} else {
				otherChar = rune(input[y+1][x])
				visited[location.Location{X: x, Y: y + 1}] = struct{}{}

				// Establish whether the dot is above or below the label
				if y > 0 && input[y-1][x] == '.' {
					point = location.Location{X: x, Y: y - 1}
				} else {
					point = location.Location{X: x, Y: y + 2}
				}
			}

			visited[point] = struct{}{}
			if c < otherChar {
				pointLocs[point] = string([]rune{c, otherChar})
			} else {
				pointLocs[point] = string([]rune{otherChar, c})
			}
		}
	}

	return pointLocs
}

type parseNode struct {
	position location.Location
	history  location.LocationSet
	distance int
}

func parseGraph(input []string) graph {
	pointLocs := parsePointLocs(input)

	graph := graph{}

	for loc, locName := range pointLocs {
		if _, gPresent := graph[locName]; !gPresent {
			graph[locName] = map[string]int{}
		}
		queue := []parseNode{
			parseNode{position: loc, history: location.LocationSet{}, distance: 0},
		}

		for len(queue) > 0 {
			currentNode := queue[0]
			currentNode.history.AddLoc(currentNode.position)
			queue = queue[1:]

			for _, adjacent := range currentNode.position.AdjacentLocations() {
				if _, present := currentNode.history[adjacent]; present {
					continue
				}

				if adjacent.X < 0 || adjacent.X > len(input[0]) || adjacent.Y < 0 || adjacent.Y > len(input) {
					continue
				}

				if input[adjacent.Y][adjacent.X] != '.' {
					continue
				}

				distance := currentNode.distance + 1

				if otherLocName, present := pointLocs[adjacent]; present {
					if existingDistance, gPresent := graph[locName][otherLocName]; gPresent {
						if distance < existingDistance {
							graph[locName][otherLocName] = distance
						}
					} else {
						graph[locName][otherLocName] = distance
					}
					continue
				} else {
					nextSearchNode := parseNode{
						position: adjacent,
						history:  *currentNode.history.Clone(),
						distance: distance,
					}

					queue = append(queue, nextSearchNode)
				}
			}
		}
	}
	return graph
}
