package main

import (
	"reflect"
	"testing"

	"../utils"
)

type example struct {
	input  []int
	output []int
	phases int
}

func TestPattern(t *testing.T) {
	examples := map[int][]int{
		0: {0, 1, 0, -1},
		1: {0, 0, 1, 1, 0, 0, -1, -1},
		2: {0, 0, 0, 1, 1, 1, 0, 0, 0, -1, -1, -1},
	}

	for pos, pat := range examples {
		result := pattern(pos)

		if !reflect.DeepEqual(result, pat) {
			t.Errorf("Expected to get %v for position %d but got %v", pat, pos, result)
		}
	}
}

func TestSinglePhaseFFT(t *testing.T) {
	examples := []example{
		example{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			output: []int{4, 8, 2, 2, 6, 1, 5, 8},
		},
		example{
			input:  []int{4, 8, 2, 2, 6, 1, 5, 8},
			output: []int{3, 4, 0, 4, 0, 4, 3, 8},
		},
		example{
			input:  []int{3, 4, 0, 4, 0, 4, 3, 8},
			output: []int{0, 3, 4, 1, 5, 5, 1, 8},
		},
		example{
			input:  []int{0, 3, 4, 1, 5, 5, 1, 8},
			output: []int{0, 1, 0, 2, 9, 4, 9, 8},
		},
	}

	for _, example := range examples {
		input := utils.CopyInts(example.input)
		result := nextFFSPhase(input)
		if !reflect.DeepEqual(result, example.output) {
			t.Errorf("Expected to get %v from %v but got %v", example.output, example.input, result)
		}
	}
}

func TestMultiPhaseFFT(t *testing.T) {
	examples := []example{
		example{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			output: []int{0, 1, 0, 2, 9, 4, 9, 8},
			phases: 4,
		},
		example{
			input:  []int{8, 0, 8, 7, 1, 2, 2, 4, 5, 8, 5, 9, 1, 4, 5, 4, 6, 6, 1, 9, 0, 8, 3, 2, 1, 8, 6, 4, 5, 5, 9, 5},
			output: []int{2, 4, 1, 7, 6, 1, 7, 6},
			phases: 100,
		},
		example{
			input:  []int{1, 9, 6, 1, 7, 8, 0, 4, 2, 0, 7, 2, 0, 2, 2, 0, 9, 1, 4, 4, 9, 1, 6, 0, 4, 4, 1, 8, 9, 9, 1, 7},
			output: []int{7, 3, 7, 4, 5, 4, 1, 8},
			phases: 100,
		},
		example{
			input:  []int{6, 9, 3, 1, 7, 1, 6, 3, 4, 9, 2, 9, 4, 8, 6, 0, 6, 3, 3, 5, 9, 9, 5, 9, 2, 4, 3, 1, 9, 8, 7, 3},
			output: []int{5, 2, 4, 3, 2, 1, 3, 3},
			phases: 100,
		},
	}

	for _, example := range examples {
		input := utils.CopyInts(example.input)
		result := abbreviatedFFS(input, example.phases)
		if !reflect.DeepEqual(result, example.output) {
			t.Errorf("Expected to get %v from %v over %d phases, but got %v",
				example.output, example.input, example.phases, result)
		}
	}
}
