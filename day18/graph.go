package main

import (
	. "../utils/location"
)

type graph = map[rune]map[rune]int

func findStartPoint(maze []string) Location {
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			if maze[y][x] == '@' {
				return Location{X: x, Y: y}
			}
		}
	}
	panic("Could not find starting point in maze!")
}

func buildGraph(maze []string) graph {
	g := graph{}

	locs := map[rune]Location{
		'@': findStartPoint(maze),
	}

	verticesToVisit := []rune{'@'}
	visitedVertices := map[rune]struct{}{}

	for len(verticesToVisit) > 0 {
		// BFS from each node to immediate neighbours
		nextVerticesToVisit := []rune{}
		for _, vertex := range verticesToVisit {
			g[vertex] = map[rune]int{}
			visitedVertices[vertex] = struct{}{}
			exploredLocs := map[Location]struct{}{}
			locsToExplore := map[Location]int{locs[vertex]: 0}

			for len(locsToExplore) > 0 {
				nextLocsToExplore := map[Location]int{}

				for currentLoc, depth := range locsToExplore {
					if _, present := exploredLocs[currentLoc]; present {
						continue
					}
					exploredLocs[currentLoc] = struct{}{}

					for _, dir := range []Direction{North, South, West, East} {
						nextLoc := currentLoc.Head(dir)
						if nextLoc == locs[vertex] {
							continue
						}

						// The maze is enclosed by walls, so we don't have to worry about out of bounds indexing
						nextTile := rune(maze[nextLoc.Y][nextLoc.X])
						switch nextTile {
						case '#':
						case '.':
							if _, present := nextLocsToExplore[nextLoc]; !present {
								nextLocsToExplore[nextLoc] = depth + 1
							}
						default:
							if _, present := visitedVertices[nextTile]; !present {
								nextVerticesToVisit = append(nextVerticesToVisit, nextTile)
							}

							g[vertex][nextTile] = depth + 1
							locs[nextTile] = nextLoc
						}
					}
				}

				locsToExplore = nextLocsToExplore
			}

		}
		verticesToVisit = nextVerticesToVisit
	}

	return g
}
