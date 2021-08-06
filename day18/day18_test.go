package main

import (
	"reflect"
	"testing"
)

type graphExample struct {
	maze  []string
	graph graph
}

func TestBuildGraph(t *testing.T) {
	examples := []graphExample{
		graphExample{
			maze: []string{
				"#########",
				"#b.A.@.a#",
				"#########",
			},
			graph: graph{
				'b': {
					'A': 2,
				},
				'A': {
					'b': 2,
					'@': 2,
				},
				'@': {
					'A': 2,
					'a': 2,
				},
				'a': {
					'@': 2,
				},
			},
		},
		graphExample{
			maze: []string{
				"#################",
				"#i.G..c...e..H.p#",
				"########.########",
				"#j.A..b...f..D.o#",
				"########@########",
				"#k.E..a...g..B.n#",
				"########.########",
				"#l.F..d...h..C.m#",
				"#################",
			},
			graph: graph{
				'i': {
					'G': 2,
				},
				'G': {
					'i': 2,
					'c': 3,
				},
				'c': {
					'G': 3,
					'e': 4,
					'b': 6,
					'f': 6,
					'@': 5,
				},
				'e': {
					'c': 4,
					'H': 3,
					'b': 6,
					'f': 6,
					'@': 5,
				},
				'H': {
					'e': 3,
					'p': 2,
				},
				'p': {
					'H': 2,
				},
				'j': {
					'A': 2,
				},
				'A': {
					'j': 2,
					'b': 3,
				},
				'b': {
					'A': 3,
					'f': 4,
					'c': 6,
					'e': 6,
					'@': 3,
				},
				'f': {
					'D': 3,
					'b': 4,
					'c': 6,
					'e': 6,
					'@': 3,
				},
				'D': {
					'f': 3,
					'o': 2,
				},
				'o': {
					'D': 2,
				},
				'@': {
					'c': 5,
					'e': 5,
					'b': 3,
					'f': 3,
					'a': 3,
					'g': 3,
					'd': 5,
					'h': 5,
				},
				'k': {
					'E': 2,
				},
				'E': {
					'k': 2,
					'a': 3,
				},
				'a': {
					'E': 3,
					'@': 3,
					'g': 4,
					'd': 6,
					'h': 6,
				},
				'g': {
					'B': 3,
					'@': 3,
					'a': 4,
					'd': 6,
					'h': 6,
				},
				'B': {
					'g': 3,
					'n': 2,
				},
				'n': {
					'B': 2,
				},
				'l': {
					'F': 2,
				},
				'F': {
					'l': 2,
					'd': 3,
				},
				'd': {
					'F': 3,
					'h': 4,
					'a': 6,
					'g': 6,
					'@': 5,
				},
				'h': {
					'd': 4,
					'a': 6,
					'g': 6,
					'@': 5,
					'C': 3,
				},
				'C': {
					'h': 3,
					'm': 2,
				},
				'm': {
					'C': 2,
				},
			},
		},
	}

	for _, example := range examples {
		actual := buildGraph(example.maze)
		if !reflect.DeepEqual(example.graph, actual) {
			t.Errorf("Expected graph to be %v but was %v", example.graph, actual)
		}
	}
}

func TestNodeQueue(t *testing.T) {
	queue := newNodeQueue(&pathNode{distance: 5})

	queue.add(&pathNode{distance: 2})
	queue.add(&pathNode{distance: 8})
	queue.add(&pathNode{distance: 5})

	two := queue.next().(*pathNode)
	if two.distance != 2 {
		t.Errorf("Expected 2, got %v", two.distance)
	}

	five := queue.next().(*pathNode)
	if five.distance != 5 {
		t.Errorf("Expected 5, got %v", five.distance)
	}

	fiveAgain := queue.next().(*pathNode)
	if fiveAgain.distance != 5 {
		t.Errorf("Expected 5, got %v", fiveAgain.distance)
	}

	eight := queue.next().(*pathNode)
	if eight.distance != 8 {
		t.Errorf("Expected 8, got %v", eight.distance)
	}
}

type example struct {
	maze               []string
	shortestPathLength int
}

func TestShortestPath(t *testing.T) {
	examples := []example{
		example{
			maze: []string{
				"#########",
				"#b.A.@.a#",
				"#########",
			},
			shortestPathLength: 8,
		},
		example{
			maze: []string{
				"########################",
				"#f.D.E.e.C.b.A.@.a.B.c.#",
				"######################.#",
				"#d.....................#",
				"########################",
			},
			shortestPathLength: 86,
		},
		example{
			maze: []string{
				"########################",
				"#...............b.C.D.f#",
				"#.######################",
				"#.....@.a.B.c.d.A.e.F.g#",
				"########################",
			},
			shortestPathLength: 132,
		},
		example{
			maze: []string{
				"#################",
				"#i.G..c...e..H.p#",
				"########.########",
				"#j.A..b...f..D.o#",
				"########@########",
				"#k.E..a...g..B.n#",
				"########.########",
				"#l.F..d...h..C.m#",
				"#################",
			},
			shortestPathLength: 136,
		},
		example{
			maze: []string{
				"########################",
				"#@..............ac.GI.b#",
				"###d#e#f################",
				"###A#B#C################",
				"###g#h#i################",
				"########################",
			},
			shortestPathLength: 81,
		},
	}

	for _, example := range examples {
		actual := shortestPath(example.maze)
		if example.shortestPathLength != actual {
			t.Errorf("Expected maze to have length %d but had length %d\n", example.shortestPathLength, actual)
		}
	}
}
