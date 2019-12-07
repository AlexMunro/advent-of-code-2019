package intcode

import (
	"fmt"
)

func readParams(opcode int) int {
	switch opcode {
	case 1, 2, 5, 6, 7, 8:
		return 2
	case 3:
		return 0
	case 4:
		return 1
	default:
		panic(fmt.Sprintf("Opcode %d not recognised", opcode))
	}
}

// is this YAGNI? I dunno, it feels neater
func writeParams(opcode int) int {
	switch opcode {
	case 1, 2, 3, 7, 8:
		return 1
	case 4, 5, 6:
		return 0
	default:
		panic(fmt.Sprintf("Opcode %d not recognised", opcode))
	}
}

// The value that should be returned by a function parameter depends on the
// mode, which is derived from its position and the opcode.
func getParams(registers []int, pos int) []int {
	params := []int{}
	opcode := registers[pos] % 100

	for i := 0; i < readParams(opcode); i++ {
		mode := registers[pos] / 100
		for j := 0; j < i; j++ {
			mode /= 10
		}
		mode = mode % 10
		value := registers[pos+1+i]

		switch mode {
		case 0: // position mode
			params = append(params, registers[value])
		case 1: // immediate mode
			if value > 0 {
				params = append(params, value)
			} else {
				params = append(params, value)
			}
		default:
			panic(fmt.Sprintf("Mode %d not recognised\n", mode))
		}
	}

	// The outputs register values are effectively always in immediate mode
	for i := 0; i < writeParams(opcode); i++ {
		params = append(params, registers[pos+readParams(opcode)+1+i])
	}

	// fmt.Printf(" with params %v\n", params)
	return params
}

// ExecuteProgram and return its output. Will modify registers.
func ExecuteProgram(registers []int, input []int) []int {
	i := 0
	outputs := []int{}

	for i < len(registers) && registers[i]%100 != 99 {

		opcode := registers[i] % 100
		params := getParams(registers, i)

		switch opcode {
		case 1:
			registers[params[2]] = params[1] + params[0]
		case 2:
			registers[params[2]] = params[1] * params[0]
		case 3:
			registers[params[0]] = input[0]
		case 4:
			outputs = append(outputs, params[0])
		case 5:
			if params[0] != 0 {
				i = params[1]
				continue
			}
		case 6:
			if params[0] == 0 {
				i = params[1]
				continue
			}
		case 7:
			if params[0] < params[1] {
				registers[params[2]] = 1
			} else {
				registers[params[2]] = 0
			}
		case 8:
			if params[0] == params[1] {
				registers[params[2]] = 1
			} else {
				registers[params[2]] = 0
			}
		}

		i += (readParams(opcode) + writeParams(opcode) + 1)
	}
	return outputs
}
