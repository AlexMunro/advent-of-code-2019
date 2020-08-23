package main

import (
	"fmt"

	"../intcode"
	"../utils"
)

func main() {
	registers := utils.GetCommaSeparatedInts("input.txt")

	boost, _ := intcode.ExecuteProgram(registers, []int{1}, intcode.Channels{})
	fmt.Printf("The answer to part one is %d\n", boost[0])

	distressSignal, _ := intcode.ExecuteProgram(registers, []int{2}, intcode.Channels{})
	fmt.Printf("The answer to part two is %d\n", distressSignal[0])
}
