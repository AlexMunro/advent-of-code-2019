package main

import (
	"fmt"

	"../utils"
	"../utils/priorityqueue"
)

type searchNode struct {
	position string
	distance int
	history  map[string]struct{}
	index    int // implementing PriorityElem
}

func (sn *searchNode) GetPriority() int {
	return sn.distance
}

func (sn *searchNode) SetIndex(n int) {
	sn.index = n
}

func aaToZZ(maze graph) int {
	initialSearchNode := searchNode{
		position: "AA",
		distance: 0,
		history:  map[string]struct{}{"AA": struct{}{}},
	}

	frontier := priorityqueue.New(&initialSearchNode)

	for frontier.Len() > 0 {
		current := frontier.Next().(*searchNode)

		if current.position == "ZZ" {
			return current.distance - 1 // subtract portal hop
		}

		for neighbour, distance := range maze[current.position] {
			if _, present := current.history[neighbour]; present {
				continue
			}

			clonedHistory := map[string]struct{}{}
			for k := range current.history {
				clonedHistory[k] = struct{}{}
			}
			clonedHistory[neighbour] = struct{}{}

			nextSearchNode := searchNode{
				position: neighbour,
				distance: current.distance + distance + 1, // add 1 for the portal hop
				history:  clonedHistory,
			}

			frontier.Add(&nextSearchNode)
		}
	}

	panic("No path found from AA to ZZ")
}

func main() {
	input := utils.GetInputLines("input.txt")
	fmt.Printf("The solution to part one is: %v\n", aaToZZ(parseGraph(input)))
}
