package intcode

import (
	"fmt"
	"math"
)

func numInputs(opcode int) int {
	switch opcode {
	case 1, 2:
		return 2
	case 3:
		return 1
	case 4:
		return 0
	default:
		panic(fmt.Sprintf("Opcode %d not recognised", opcode))
	}	
}

// is this YAGNI? I dunno, it feels neater
func numOutputs(opcode int) int {
	switch opcode {
	case 1, 2, 4:
		return 1
	case 3:
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

	// fmt.Printf("Getting params starting from pos %d in %v\n", pos, registers)

	for i:= 0; i < numInputs(opcode); i++ {
		mode := opcode / int(math.Pow(10, float64(2 + i)))
		value := registers[pos + 1 + i]
		// fmt.Printf("Mode: %d with value %d", mode, value)
		switch mode {
		case 0: // position mode
			params = append(params, registers[value])
			// fmt.Printf(" is %d\n", registers[value])
		case 1: // immediate mode
			params = append(params, value)
			// fmt.Printf(" is %d\n", value)
		default:
			panic(fmt.Sprintf("Mode %d not recognised\n", mode))
		}
	}

	// The outputs register values are effectively always in immediate mode
	for i:= 0; i < numOutputs(opcode); i++{
		params = append(params, registers[pos + numInputs(opcode) + 1 + i])
	}

	// fmt.Printf("Executing instr %d with params %v\n", opcode, params)
	return params
}

// ExecuteProgram and return its output. Will modify registers.
func ExecuteProgram(registers []int, input []int) []int {
	i := 0
	outputs := []int{}

	for i < len(registers) && registers[i] % 100 != 99 {
		// fmt.Printf("Processing instruction %d\n", i)
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
		}
		// fmt.Printf("Advancing by %d registers\n", numInputs(opcode) + 1)
		i += (numInputs(opcode) + numOutputs(opcode) + 1)
	}
	return outputs
}