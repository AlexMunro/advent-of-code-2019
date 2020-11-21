package main

import (
	"container/heap"
	"fmt"
	"unicode"

	"../utils"
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

type pathNode struct {
	currentlyAt rune
	distance    int
	history     []rune
	index       int
	visited     map[rune]struct{}
}

// Priority queue code mostly lifted from golang docs: https://golang.org/pkg/container/heap/
type nodeQueue []*pathNode

func (nq nodeQueue) Len() int {
	return len(nq)
}

func (nq nodeQueue) Less(i, j int) bool {
	return nq[i].distance < nq[j].distance
}

func (nq nodeQueue) Swap(i, j int) {
	nq[i], nq[j] = nq[j], nq[i]
	nq[i].index = i
	nq[j].index = j
}

func (nq *nodeQueue) Push(x interface{}) {
	n := len(*nq)
	node := x.(*pathNode)
	node.index = n
	*nq = append(*nq, node)
}

func (nq *nodeQueue) Pop() interface{} {
	old := *nq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
	*nq = old[0 : n-1]
	return node
}

func (nq *nodeQueue) update(node *pathNode) {
	heap.Fix(nq, node.distance)
}

// Search the space of valid paths in lowest cumulative distance order
// Hopefully most paths will eventually be invalid, which should restrict the space
// Returns as soon as a solution is found - this solution should be optimal
func shortestPath(maze []string) int {
	graph := buildGraph(maze)
	totalNodesToVisit := len(graph)

	initialNode := &pathNode{
		currentlyAt: '@',
		distance:    0,
		history:     []rune{'@'},
		visited:     map[rune]struct{}{'@': struct{}{}},
	}

	frontier := make(nodeQueue, 1)
	frontier[0] = initialNode
	heap.Init(&frontier)

	for len(frontier) > 0 {
		currentNode := heap.Pop(&frontier).(*pathNode)

		// Generate new nodes for all immediate neighbours
	NEIGHBOURS:
		for neighbour, distance := range graph[currentNode.currentlyAt] {
			if _, present := currentNode.visited[neighbour]; present {
				// For any visited nodes, we need to check that we aren't going in circles
				for _, char := range currentNode.history {
					if neighbour == char {
						continue NEIGHBOURS
					}
				}

				newNode := &pathNode{
					currentlyAt: neighbour,
					distance:    currentNode.distance + distance,
					history:     append(currentNode.history, neighbour),
					visited:     currentNode.visited,
				}
				heap.Push(&frontier, newNode)
			} else {
				// Skip doors we can't open yet
				if unicode.IsUpper(neighbour) {
					if _, present := currentNode.visited[unicode.ToLower(neighbour)]; !present {
						continue
					}
				}

				newVisitedNodes := map[rune]struct{}{}
				for k := range currentNode.visited {
					newVisitedNodes[k] = struct{}{}
				}
				newVisitedNodes[neighbour] = struct{}{}

				newNode := &pathNode{
					currentlyAt: neighbour,
					distance:    currentNode.distance + distance,
					history:     []rune{}, // cycles no longer relevant since we're in new territory
					visited:     newVisitedNodes,
				}
				heap.Push(&frontier, newNode)

				// Finally, check for goal condition being reached
				if totalNodesToVisit == len(newNode.visited) {
					return newNode.distance
				}
			}
		}
	}

	panic("Failed to find a solution")
}

func main() {
	input := utils.GetInputLines("input.txt")

	fmt.Printf("The answer to part one is %v", shortestPath(input))
}
