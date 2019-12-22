package main

import (
	"fmt"

	"../intcode"
	"../utils"
)

func main() {
	registers := utils.GetCommaSeparatedInts("input.txt")

	boost := intcode.ExecuteProgram(registers, []int{1}, nil, nil, nil)[0]
	fmt.Printf("The answer to part one is %d\n", boost)

	distressSignal := intcode.ExecuteProgram(registers, []int{2}, nil, nil, nil)[0]
	fmt.Printf("The answer to part two is %d\n", distressSignal)
}
