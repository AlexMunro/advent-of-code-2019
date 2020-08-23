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
	output, _ := intcode.ExecuteProgram(registers1, input, intcode.Channels{})
	fmt.Printf("The answer to part one is %d\n", output[len(output)-1])

	input = []int{5}
	output, _ = intcode.ExecuteProgram(registers, input, intcode.Channels{})
	fmt.Printf("The answer to part two is %d\n", output[len(output)-1])
}
