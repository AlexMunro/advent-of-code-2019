package utils

import (
	"io/ioutil"
	"math"
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

// GetInputSingleString retrieves input from a file and
// returns it as a single string
func GetInputSingleString(filename string) string {
	data, errors := ioutil.ReadFile(filename)
	if errors != nil {
		panic(errors)
	}
	return string(data)
}

// GetCommaSeparatedInts retrieves a slice of ints where
// these are all given on one line
func GetCommaSeparatedInts(filename string) []int {
	intStrings := strings.Split(GetInputSingleString(filename), ",")
	ints := make([]int, 0, len(intStrings))

	for _, s := range intStrings {
		n, errors := strconv.Atoi(s)
		if errors != nil {
			panic(errors)
		}
		ints = append(ints, n)
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

// Min of an int slice
func Min(xs []int) int {
	min := xs[0]
	for _, x := range xs[1:] {
		if min > x {
			min = x
		}
	}
	return min
}

// Max of an int slice
func Max(xs []int) int {
	max := xs[0]
	for _, x := range xs[1:] {
		if max < x {
			max = x
		}
	}
	return max
}

// CopyInts creates and returns a copy of a slice of ints
func CopyInts(xs []int) []int {
	xsCopy := make([]int, len(xs))
	copy(xsCopy, xs)
	return xsCopy
}

// Abs value of n (remove any negative sign)
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// CeilDiv float divides n by divisor and returns the ceiling
func CeilDiv(n, divisor int) int {
	return int(math.Ceil(float64(n) / float64(divisor)))
}
