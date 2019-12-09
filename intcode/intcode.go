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
	case 4, 9:
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
	case 4, 5, 6, 9:
		return 0
	default:
		panic(fmt.Sprintf("Opcode %d not recognised", opcode))
	}
}

type registerMap struct {
	contents map[int]int
}

func new(registerArray []int) *registerMap {
	rm := registerMap{map[int]int{}}
	for i, n := range registerArray {
		rm.contents[i] = n
	}
	return &rm
}

func (rm *registerMap) get(pos int) int {
	if _, present := rm.contents[pos]; !present {
		rm.contents[pos] = 0
	}
	return rm.contents[pos]
}

func (rm *registerMap) set(pos, n int) {
	rm.contents[pos] = n
}

// The value that should be returned by a function parameter depends on the
// mode, which is derived from its position and the opcode.
func getParams(registers *registerMap, pos, relativeBase int) []int {
	params := []int{}
	opcode := registers.get(pos) % 100

	for i := 0; i < readParams(opcode); i++ {
		mode := registers.get(pos) / 100
		for j := 0; j < i; j++ {
			mode /= 10
		}
		mode = mode % 10

		switch mode {
		case 0: // position mode
			params = append(params, registers.get(registers.get(pos+1+i)))
		case 1: // immediate mode
			params = append(params, registers.get(pos+1+i))
		case 2: // relative mode
			params = append(params, registers.get(registers.get(pos+1+i)+relativeBase))
		default:
			panic(fmt.Sprintf("Mode %d not recognised for read parameters\n", mode))
		}
	}

	// The outputs register values should be calculated as though they are immediate, because
	// the instruction will access that register anyway
	for i := 0; i < writeParams(opcode); i++ {
		mode := registers.get(pos) / 100
		for j := 0; j < i+readParams(opcode); j++ {
			mode /= 10
		}
		mode = mode % 10

		switch mode {
		case 0, 1:
			params = append(params, registers.get(pos+readParams(opcode)+1+i))
		case 2:
			params = append(params, registers.get(pos+readParams(opcode)+1+i)+relativeBase)
		default:
			panic(fmt.Sprintf("Mode %d not recognised for write parameters\n", mode))
		}
	}
	return params
}

// ExecuteProgram and return its output. Will modify registers.
// inChan is used to find inputs where input has been depleted if it is not nil.
// outChan is used to write outputs if not nil. All outputs are also returned as a slice.
func ExecuteProgram(registerArray []int, input []int, inChan <-chan int, outChan chan<- int) []int {
	if outChan != nil {
		defer close(outChan)
	}

	registers := new(registerArray)

	relativeBase := 0
	i := 0
	outputs := []int{}

	for registers.get(i)%100 != 99 {
		opcode := registers.get(i) % 100
		params := getParams(registers, i, relativeBase)

		switch opcode {
		case 1:
			registers.set(params[2], params[1]+params[0])
		case 2:
			registers.set(params[2], params[1]*params[0])
		case 3:
			if len(input) > 0 {
				registers.set(params[0], input[0])
				input = input[1:]
			} else if inChan != nil {
				registers.set(params[0], <-inChan)
			} else {
				panic("No further input available")
			}
		case 4:
			outputs = append(outputs, params[0])
			if outChan != nil {
				outChan <- params[0]
			}
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
				registers.set(params[2], 1)
			} else {
				registers.set(params[2], 0)
			}
		case 8:
			if params[0] == params[1] {
				registers.set(params[2], 1)
			} else {
				registers.set(params[2], 0)
			}
		case 9:
			relativeBase += params[0]
		}

		i += (len(params) + 1)
	}

	// Need to modify the initial input slice to support legacy tests
	// Could do with a refactor, but somehow I don't see that happening
	for i := range registerArray {
		registerArray[i] = registers.get(i)
	}

	return outputs
}
