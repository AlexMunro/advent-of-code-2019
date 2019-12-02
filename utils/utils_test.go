package utils

import (
	"testing"
)

func TestSum(t *testing.T) {
	example := []int{ 1, 2, 3, 4, 5 }
	expected := 15

	result := Sum(example)

	if result != expected {
		t.Errorf("Expected the sum of %v to be %d but was %d", example, expected, result)
	}
}