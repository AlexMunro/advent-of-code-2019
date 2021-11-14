package main

import (
	"reflect"
	"testing"
)

type graphExample struct {
	maze  []string
	graph graph
}

func TestBuildSimpleGraph(t *testing.T) {
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
		actual := buildSimpleGraph(example.maze)
		if !reflect.DeepEqual(example.graph, actual) {
			t.Errorf("Expected graph to be %v but was %v", example.graph, actual)
		}
	}
}

func TestBuildQuarterGraphs(t *testing.T) {
	examples := []graphExample{
		graphExample{
			maze: []string{
				"#######",
				"#a.#Cd#",
				"##...##",
				"##.@.##",
				"##...##",
				"#cB#Ab#",
				"#######",
			},
			graph: graph{
				'a': {'0': 2}, '0': {'a': 2},
				'1': {'C': 1}, 'C': {'1': 1, 'd': 1}, 'd': {'C': 1},
				'2': {'B': 1}, 'B': {'2': 1, 'c': 1}, 'c': {'B': 1},
				'3': {'A': 1}, 'A': {'3': 1, 'b': 1}, 'b': {'A': 1},
			},
		},
		graphExample{
			maze: []string{
				"###############",
				"#d.ABC.#.....a#",
				"######.#.######",
				"#######@#######",
				"######..#######",
				"#b.....#.....c#",
				"###############",
			},
			graph: graph{
				'0': {'C': 2}, 'C': {'0': 2, 'B': 1}, 'B': {'C': 1, 'A': 1}, 'A': {'B': 1, 'd': 2}, 'd': {'A': 2},
				'1': {'a': 6}, 'a': {'1': 6},
				'2': {'b': 6}, 'b': {'2': 6},
				'3': {'c': 6}, 'c': {'3': 6},
			},
		},
	}

	for _, example := range examples {
		if !reflect.DeepEqual(buildQuarterGraphs(example.maze), example.graph) {
			t.Errorf("Expected quartered graph to be %v but was %v", example.graph, buildQuarterGraphs(example.maze))
		}
	}

}
