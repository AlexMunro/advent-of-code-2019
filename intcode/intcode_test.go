package intcode

import (
	"reflect"
	"strconv"
	"sync"
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
		ExecuteProgram(afterRegisters, tc.inputs, nil, nil, nil)
		if !(reflect.DeepEqual(afterRegisters, tc.afterRegisters)) {
			t.Errorf("Expected to get %v from %v but got %v", tc.afterRegisters, tc.beforeRegisters, afterRegisters)
		}
	}
}

func TestInputOutputProgram(t *testing.T) {
	examples := []testCase{
		{beforeRegisters: []int{103, 1, 99, 19}, inputs: []int{1234}, afterRegisters: []int{103, 1234, 99, 19}, outputs: []int{}},
		{beforeRegisters: []int{104, 583, 99, 19}, afterRegisters: []int{104, 583, 99, 19}, outputs: []int{583}},
		{beforeRegisters: []int{3, 0, 4, 0, 99}, afterRegisters: []int{-27, 0, 4, 0, 99}, inputs: []int{-27}, outputs: []int{-27}},
	}

	for _, tc := range examples {
		afterRegisters := utils.CopyInts(tc.beforeRegisters)
		outputs := ExecuteProgram(afterRegisters, tc.inputs, nil, nil, nil)
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
		outputs := ExecuteProgram(afterRegisters, tc.inputs, nil, nil, nil)
		if !(reflect.DeepEqual(outputs, tc.outputs)) {
			t.Errorf("Expected to get outputs %v from %v with inputs %v but got %v with outputs %v",
				tc.afterRegisters, tc.beforeRegisters, tc.inputs, afterRegisters, outputs)
		}
	}
}

func TestMixedOpProgram(t *testing.T) {
	examples := []testCase{
		{beforeRegisters: []int{103, 1, 99, 19}, inputs: []int{1234}, afterRegisters: []int{103, 1234, 99, 19}, outputs: []int{}},
		{beforeRegisters: []int{104, 583, 99, 19}, afterRegisters: []int{104, 583, 99, 19}, outputs: []int{583}},
	}

	for _, tc := range examples {
		afterRegisters := utils.CopyInts(tc.beforeRegisters)
		outputs := ExecuteProgram(afterRegisters, tc.inputs, nil, nil, nil)
		if !(reflect.DeepEqual(afterRegisters, tc.afterRegisters)) || !(reflect.DeepEqual(outputs, tc.outputs)) {
			t.Errorf("Expected to get %v with outputs %v from %v with inputs %v but got %v with outputs %v",
				tc.afterRegisters, tc.outputs, tc.beforeRegisters, tc.inputs, afterRegisters, outputs)
		}
	}
}

// Uses the examples provided in the day 7 instructions
type concurrentExample struct {
	registers []int
	inputs    [][]int
	result    int
}

func TestConcurrentProgram(t *testing.T) {
	examples := []concurrentExample{
		{
			registers: []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27,
				1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
			inputs: [][]int{{9, 0}, {8}, {7}, {6}, {5}},
			result: 139629729,
		},
		{
			registers: []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005,
				55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1,
				55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
			inputs: [][]int{{9, 0}, {7}, {8}, {5}, {6}},
			result: 18216,
		},
	}

	for _, tc := range examples {
		channels := []chan int{}

		for i := 0; i < len(tc.inputs); i++ {
			channels = append(channels, make(chan int))
		}

		var wg sync.WaitGroup
		wg.Add(len(channels) - 1)

		for i := 0; i < len(tc.inputs); i++ {
			r := utils.CopyInts(tc.registers)
			go func(i int, input, registers []int) {
				if i < len(tc.inputs)-1 { // The last process will write once after the others have finished
					defer wg.Done()
				}
				ExecuteProgram(r, input, channels[i], channels[(i+1)%len(channels)], nil)
			}(i, tc.inputs[i], r)
		}

		wg.Wait()
		var result int
		result = <-channels[0]

		if result != tc.result {
			t.Errorf("Expected to get %d but got %d", tc.result, result)
		}
	}
}

func TestRelativeModeProgram(t *testing.T) {
	// This program should output itself
	expected := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	registers := utils.CopyInts(expected)
	result := ExecuteProgram(registers, []int{}, nil, nil, nil)
	if !reflect.DeepEqual(registers, result) {
		t.Errorf("Expected %v to output itself but got %v", registers, result)
	}

	// This program should produce a 16 digit number
	registers = []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	result = ExecuteProgram(registers, []int{}, nil, nil, nil)
	if len(strconv.Itoa(result[0])) != 16 {
		t.Errorf("Expected %v to output a 16 digit number but got %d", registers, result[0])
	}

	// This program should output its middle register
	registers = []int{104, 1125899906842624, 99}
	result = ExecuteProgram(registers, []int{}, nil, nil, nil)
	if !reflect.DeepEqual([]int{1125899906842624}, result) {
		t.Errorf("Expected %v to output 1125899906842624 but got %d", registers, result)
	}
}

func TestExtraTests(t *testing.T) {
	examples := []testCase{
		// Test cases thanks to:
		// https://www.reddit.com/r/adventofcode/comments/e8aw9j/2019_day_9_part_1_how_to_fix_203_error/fac3294/
		{beforeRegisters: []int{109, -1, 4, 1, 99}, inputs: []int{1}, outputs: []int{-1}},
		{beforeRegisters: []int{109, -1, 104, 1, 99}, inputs: []int{1}, outputs: []int{1}},
		{beforeRegisters: []int{109, -1, 204, 1, 99}, inputs: []int{1}, outputs: []int{109}},
		{beforeRegisters: []int{109, 1, 9, 2, 204, -6, 99}, inputs: []int{1}, outputs: []int{204}},
		{beforeRegisters: []int{109, 1, 109, 9, 204, -6, 99}, inputs: []int{1}, outputs: []int{204}},
		{beforeRegisters: []int{109, 1, 209, -1, 204, -106, 99}, inputs: []int{1}, outputs: []int{204}},
		{beforeRegisters: []int{109, 1, 3, 3, 204, 2, 99}, inputs: []int{1234}, outputs: []int{1234}},
		{beforeRegisters: []int{109, 1, 203, 2, 204, 2, 99}, inputs: []int{4321}, outputs: []int{4321}},
	}

	for _, tc := range examples {
		outputs := ExecuteProgram(tc.beforeRegisters, tc.inputs, nil, nil, nil)
		if !(reflect.DeepEqual(outputs, tc.outputs)) {
			t.Errorf("Expected to get outputs %v but got outputs %v", tc.outputs, outputs)
		}
	}

}
