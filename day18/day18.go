package main

import (
	"fmt"
	"unicode"

	"../utils"
	"../utils/priorityqueue"
)

type runeset map[rune]struct{}

func contains(rs runeset, r rune) bool {
	if _, present := rs[r]; present {
		return true
	}
	return false
}

func keyCount(graph graph) int {
	count := 0
	for node := range graph {
		if unicode.IsLower(node) {
			count++
		}
	}
	return count
}

// Search the space of valid paths in lowest cumulative distance order
// Returns as soon as a solution is found - this solution should be optimal
func shortestPath(maze []string) int {
	graph := buildSimpleGraph(maze)
	totalNodesToVisit := keyCount(graph)

	initialNode := &pathNode{
		position: '@',
		distance: 0,
		visited:  runeset{'@': struct{}{}},
		keyCount: 0,
	}

	frontier := priorityqueue.New(initialNode)

	// Prune our search
	previouslyConsidered := map[string]int{
		initialNode.hash(): 0,
	}

	for frontier.Len() > 0 {
		currentNode := frontier.Next().(*pathNode)
		if totalNodesToVisit == currentNode.keyCount {
			return currentNode.distance
		}

		// Generate new nodes for all immediate neighbours
		for neighbour, distance := range graph[currentNode.position] {
			var newNode *pathNode

			if contains(currentNode.visited, neighbour) {
				newNode = &pathNode{
					position: neighbour,
					distance: currentNode.distance + distance,
					visited:  currentNode.visited,
					keyCount: currentNode.keyCount,
				}
			} else {
				// Skip doors we can't open yet
				if unicode.IsUpper(neighbour) {
					if _, present := currentNode.visited[unicode.ToLower(neighbour)]; !present {
						continue
					}
				}

				newVisitedNodes := runeset{}
				for k := range currentNode.visited {
					newVisitedNodes[k] = struct{}{}
				}
				newVisitedNodes[neighbour] = struct{}{}

				newKeyCount := currentNode.keyCount
				if unicode.IsLower(neighbour) {
					newKeyCount++
				}

				newNode = &pathNode{
					position: neighbour,
					distance: currentNode.distance + distance,
					visited:  newVisitedNodes,
					keyCount: newKeyCount,
				}
			}

			// Prune backtracking branches of search where there are no new keys
			if prevDist, present := previouslyConsidered[newNode.hash()]; present {
				if prevDist <= newNode.distance {
					continue
				}
			}

			previouslyConsidered[newNode.hash()] = newNode.distance

			frontier.Add(newNode)
		}
	}

	panic("Failed to find a solution")
}

// Search as above with four robots instead of one!
func shortestQuadPath(maze []string) int {
	graph := buildQuarterGraphs(maze)
	totalNodesToVisit := keyCount(graph)

	initialNode := &quadPathNode{
		positions: []rune{'0', '1', '2', '3'},
		distance:  0,
		visited: runeset{
			'0': struct{}{},
			'1': struct{}{},
			'2': struct{}{},
			'3': struct{}{},
		},
		keyCount: 0,
	}

	frontier := priorityqueue.New(initialNode)

	previouslyConsidered := map[string]int{
		initialNode.hash(): 0,
	}

	for frontier.Len() > 0 {
		currentNode := frontier.Next().(*quadPathNode)

		if previousDistance, present := previouslyConsidered[currentNode.hash()]; present {
			if previousDistance < currentNode.distance {
				continue
			}
		}

		if totalNodesToVisit == currentNode.keyCount {
			return currentNode.distance
		}

		// Generate new nodes for all immediate neighbours
		for index := range [4]int{} {
			for neighbour, distance := range graph[currentNode.positions[index]] {
				var newNode *quadPathNode
				neighbourPositions := make([]rune, len(currentNode.positions))
				copy(neighbourPositions, currentNode.positions)
				neighbourPositions[index] = neighbour

				if _, present := currentNode.visited[neighbour]; present {
					newNode = &quadPathNode{
						positions: neighbourPositions,
						distance:  currentNode.distance + distance,
						visited:   currentNode.visited,
						keyCount:  currentNode.keyCount,
					}
				} else {
					// Skip doors we can't open yet
					if unicode.IsUpper(neighbour) {
						if _, present := currentNode.visited[unicode.ToLower(neighbour)]; !present {
							continue
						}
					}

					newVisitedNodes := runeset{}
					for k := range currentNode.visited {
						newVisitedNodes[k] = struct{}{}
					}
					newVisitedNodes[neighbour] = struct{}{}

					newKeyCount := currentNode.keyCount
					if unicode.IsLower(neighbour) {
						newKeyCount++
					}

					newNode = &quadPathNode{
						positions: neighbourPositions,
						distance:  currentNode.distance + distance,
						visited:   newVisitedNodes,
						keyCount:  newKeyCount,
					}
				}

				// Prune backtracking branches of search where there are no new keys
				if prevDist, present := previouslyConsidered[newNode.hash()]; present {
					if prevDist <= newNode.distance {
						continue
					}
				}

				previouslyConsidered[newNode.hash()] = newNode.distance

				frontier.Add(newNode)
			}
		}
	}

	panic("Failed to find a solution")
}

func main() {
	input := utils.GetInputLines("input.txt")

	fmt.Printf("The answer to part one is %v\n", shortestPath(input))
	fmt.Printf("The answer to part two is %v\n", shortestQuadPath(input))
}
