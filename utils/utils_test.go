package utils

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	example := []int{1, 2, 3, 4, 5}
	expected := 15

	result := Sum(example)

	if result != expected {
		t.Errorf("Expected the sum of %v to be %d but was %d", example, expected, result)
	}
}

func TestRepeatInts(t *testing.T) {
	example := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}

	result := RepeatInts(example, 2)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v repeated %d times to be %v but was %v", example, 2, expected, result)
	}
}
