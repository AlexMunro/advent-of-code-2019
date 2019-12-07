package intcode

import (
	"reflect"
	"testing"

	"../utils"
)

type testCase struct {
	beforeRegisters []int
	inputs          []int
	afterRegisters  []int
	outputs         []int
}

func TestAdditionMultiplicationProgram(t *testing.T) {
	examples := []testCase{
		// position mode
		{beforeRegisters: []int{1, 0, 0, 0, 99}, afterRegisters: []int{2, 0, 0, 0, 99}},
		{beforeRegisters: []int{2, 3, 0, 3, 99}, afterRegisters: []int{2, 3, 0, 6, 99}},
		{beforeRegisters: []int{2, 4, 4, 5, 99, 0}, afterRegisters: []int{2, 4, 4, 5, 99, 9801}},
		{beforeRegisters: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, afterRegisters: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},

		// immediate mode
		{beforeRegisters: []int{1002, 4, 3, 4, 33}, afterRegisters: []int{1002, 4, 3, 4, 99}},
		{beforeRegisters: []int{11002, 4, 3, 4, 33}, afterRegisters: []int{11002, 4, 3, 4, 99}},
		{beforeRegisters: []int{1101, 100, -1, 4, 0}, afterRegisters: []int{1101, 100, -1, 4, 99}},
		{beforeRegisters: []int{1002, 4, -3, 5, 99, 48}, afterRegisters: []int{1002, 4, -3, 5, 99, -297}},
	}

	for _, tc := range examples {
		afterRegisters := utils.CopyInts(tc.beforeRegisters)
		ExecuteProgram(afterRegisters, tc.inputs)
		if !(reflect.DeepEqual(afterRegisters, tc.afterRegisters)) {
			t.Errorf("Expected to get %v from %v but got %v", tc.afterRegisters, tc.beforeRegisters, afterRegisters)
		}
	}
}

func TestInputOutputProgram(t *testing.T) {
	examples := []testCase{
		{beforeRegisters: []int{103, 1, 99, 19}, inputs: []int{1234}, afterRegisters: []int{103, 1234, 99, 19}, outputs: []int{}},
		{beforeRegisters: []int{104, 583, 99, 19}, afterRegisters: []int{104, 583, 99, 19}, outputs: []int{583}},
		{beforeRegisters: []int{3, 0, 4, 0, 99}, afterRegisters: []int{3, 0, 4, -27, 99}, inputs: []int{-27}, outputs: []int{-27}},
	}

	for _, tc := range examples {
		afterRegisters := utils.CopyInts(tc.beforeRegisters)
		outputs := ExecuteProgram(afterRegisters, tc.inputs)
		if !(reflect.DeepEqual(afterRegisters, tc.afterRegisters)) || !(reflect.DeepEqual(outputs, tc.outputs)) {
			t.Errorf("Expected to get %v with outputs %v from %v with inputs %v but got %v with outputs %v",
				tc.afterRegisters, tc.outputs, tc.beforeRegisters, tc.inputs, afterRegisters, outputs)
		}
	}
}

func TestConditionalProgramming(t *testing.T) {
	examples := []testCase{
		// input[0] == 8

		// positional
		{beforeRegisters: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, inputs: []int{8}, outputs: []int{1}},
		{beforeRegisters: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, inputs: []int{7}, outputs: []int{0}},

		// immediate
		{beforeRegisters: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, inputs: []int{8}, outputs: []int{1}},
		{beforeRegisters: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, inputs: []int{7}, outputs: []int{0}},

		// input[0] < 8

		// positional
		{beforeRegisters: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, inputs: []int{7}, outputs: []int{1}},
		{beforeRegisters: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, inputs: []int{8}, outputs: []int{0}},
		{beforeRegisters: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, inputs: []int{9}, outputs: []int{0}},

		//immediate
		{beforeRegisters: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, inputs: []int{7}, outputs: []int{1}},
		{beforeRegisters: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, inputs: []int{8}, outputs: []int{0}},
		{beforeRegisters: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, inputs: []int{9}, outputs: []int{0}},
	}

	for _, tc := range examples {
		afterRegisters := utils.CopyInts(tc.beforeRegisters)
		outputs := ExecuteProgram(afterRegisters, tc.inputs)
		if !(reflect.DeepEqual(outputs, tc.outputs)) {
			t.Errorf("Expected to get outputs %v from %v with inputs %v but got %v with outputs %v",
				tc.afterRegisters, tc.beforeRegisters, tc.inputs, afterRegisters, outputs)
		}
	}
}

func MixedOpProgram(t *testing.T) {
	examples := []testCase{
		{beforeRegisters: []int{103, 1, 99, 19}, inputs: []int{1234}, afterRegisters: []int{103, 1234, 99, 19}, outputs: []int{}},
		{beforeRegisters: []int{104, 583, 99, 19}, afterRegisters: []int{104, 583, 99, 19}, outputs: []int{583}},
	}

	for _, tc := range examples {
		afterRegisters := utils.CopyInts(tc.beforeRegisters)
		outputs := ExecuteProgram(afterRegisters, tc.inputs)
		if !(reflect.DeepEqual(afterRegisters, tc.afterRegisters)) || !(reflect.DeepEqual(outputs, tc.outputs)) {
			t.Errorf("Expected to get %v with outputs %v from %v with inputs %v but got %v with outputs %v",
				tc.afterRegisters, tc.outputs, tc.beforeRegisters, tc.inputs, afterRegisters, outputs)
		}
	}
}
