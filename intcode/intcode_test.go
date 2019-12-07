package intcode

import (
	"reflect"
	"testing"
	"../utils"
)

func TestExecuteProgram(t *testing.T) {
	examples := map[*[]int][]int{
		&([]int{1, 0, 0, 0, 99}):              []int{2, 0, 0, 0, 99},
		&([]int{2, 3, 0, 3, 99}):              []int{2, 3, 0, 6, 99},
		&([]int{2, 4, 4, 5, 99, 0}):           []int{2, 4, 4, 5, 99, 9801},
		&([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}): []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
	}

	for k, v := range examples {
		result := utils.CopyInts(*k)
		ExecuteProgram(result, []int{})
		if !(reflect.DeepEqual(result, v)) {
			t.Errorf("Expected to get %v from %v but got %v", v, *k, result)
		}
	}
}
