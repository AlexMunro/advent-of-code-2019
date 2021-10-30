package main

import (
	"fmt"
	"unicode"

	"../utils"
)

type pathNode struct {
	position rune
	distance int
	visited  map[rune]struct{}
	keyCount int
	index    int // used by container/heap
}

func (p pathNode) hash() string {
	encodedKeys := encodeKeys(p.visited)

	return fmt.Sprint(p.position, '_', encodedKeys)
}

func (p *pathNode) getDistance() int {
	return p.distance
}

func (p *pathNode) setIndex(i int) {
	p.index = i
}

func encodeKeys(keys map[rune]struct{}) int {
	encodedKeys := 0

	for key := range keys {
		if unicode.IsLower(key) {
			position := int(key - 'a')
			shiftedPosition := 1 << position
			encodedKeys = encodedKeys | shiftedPosition
		}
	}
	return encodedKeys
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
		visited:  map[rune]struct{}{'@': struct{}{}},
		keyCount: 0,
	}

	frontier := newNodeQueue(initialNode)

	// Prune our search
	previouslyConsidered := map[string]int{
		initialNode.hash(): 0,
	}

	for frontier.Len() > 0 {
		currentNode := frontier.next().(*pathNode)
		if totalNodesToVisit == currentNode.keyCount {
			return currentNode.distance
		}

		// Generate new nodes for all immediate neighbours
		for neighbour, distance := range graph[currentNode.position] {
			var newNode *pathNode

			if _, present := currentNode.visited[neighbour]; present {
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

				newVisitedNodes := map[rune]struct{}{}
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

			frontier.add(newNode)
		}
	}

	panic("Failed to find a solution")
}

func main() {
	input := utils.GetInputLines("input.txt")

	fmt.Printf("The answer to part one is %v\n", shortestPath(input))
}
