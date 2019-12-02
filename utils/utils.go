package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// GetInputLines retrieves each line from the file at
// filename and returns each as a separate string
func GetInputLines(filename string) []string {
	data, errors := ioutil.ReadFile(filename)
	if errors != nil {
		panic(errors)
	}
	return strings.Split(string(data), "\n")
}

// GetInputInts retrieves each line from the file at
// filename and returns each as a separate int
func GetInputInts(filename string) []int {
	input := GetInputLines(filename)
	ints := make([]int, len(input))
	for i := range input {
		nextInt, errors := strconv.Atoi(input[i])
		if errors != nil {
			panic(errors)
		}
		ints[i] = nextInt
	}
	return ints
}

// Sum adds all the values in xs together
func Sum(xs []int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
}
