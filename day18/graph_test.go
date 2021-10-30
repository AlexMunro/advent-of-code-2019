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
