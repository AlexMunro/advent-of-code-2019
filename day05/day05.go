package main

import (
	"fmt"

	"../intcode"
	"../utils"
)

func main() {
	registers := utils.GetCommaSeparatedInts("input.txt")
	registers1 := utils.CopyInts(registers)
	input := []int{1}
	output := intcode.ExecuteProgram(registers1, input, nil, nil, nil)
	fmt.Printf("The answer to part one is %d\n", output[len(output)-1])

	input = []int{5}
	output = intcode.ExecuteProgram(registers, input, nil, nil, nil)
	fmt.Printf("The answer to part two is %d\n", output[len(output)-1])
}
