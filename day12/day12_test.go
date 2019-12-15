package main

import (
	"reflect"
	"testing"
)

var exampleState = []object{
	object{position{-1, 0, 2}, velocity{}},
	object{position{2, -10, -7}, velocity{}},
	object{position{4, -8, 8}, velocity{}},
	object{position{3, 5, -1}, velocity{}},
}

func TestSimulateMotion(t *testing.T) {

	examples := map[int][]object{
		0: exampleState,
		1: []object{
			object{position{2, -1, 1}, velocity{3, -1, -1}},
			object{position{3, -7, -4}, velocity{1, 3, 3}},
			object{position{1, -7, 5}, velocity{-3, 1, -3}},
			object{position{2, 2, 0}, velocity{-1, -3, 1}},
		},
		2: []object{
			object{position{5, -3, -1}, velocity{3, -2, -2}},
			object{position{1, -2, 2}, velocity{-2, 5, 6}},
			object{position{1, -4, -1}, velocity{0, 3, -6}},
			object{position{1, -4, 2}, velocity{-1, -6, 2}},
		},
	}

	for k, v := range examples {
		result := simulateMotion(exampleState, k)
		if !reflect.DeepEqual(result, v) {
			t.Errorf("Expected to get %v for %d iterations but got %v", v, k, result)
		}
	}
}

func TestFindRepeatedState(t *testing.T) {
	expected := 2772
	result := findRepeatedState(exampleState)
	if result != expected {
		t.Errorf("Expected to get %d from %v but got %d", expected, exampleState, result)
	}
}
