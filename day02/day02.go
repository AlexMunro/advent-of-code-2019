package main

import (
	"fmt"

	"../intcode"
	"../utils"
)

func findNounAndVerb(registers []int, target int) (int, int) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			r := utils.CopyInts(registers)
			r[1] = noun
			r[2] = verb
			intcode.ExecuteProgram(r, []int{}, nil, nil)

			if r[0] == target {
				return noun, verb
			}
		}
	}
	panic("No solution found")
}

func main() {
	registers := utils.GetCommaSeparatedInts("input.txt")
	registers1 := utils.CopyInts(registers)
	registers1[1] = 12
	registers1[2] = 2
	intcode.ExecuteProgram(registers1, []int{}, nil, nil)
	fmt.Printf("The answer to part one is %d\n", registers1[0])

	noun, verb := findNounAndVerb(registers, 19690720)
	fmt.Printf("The answer to part two is %d\n", (100*noun)+verb)
}
