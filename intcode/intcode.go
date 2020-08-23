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

type registerMap map[int]int

func arrayToMap(registerArray []int) registerMap {
	rm := registerMap{}
	for i, n := range registerArray {
		rm[i] = n
	}
	return rm
}

func (rm registerMap) get(pos int) int {
	if _, present := rm[pos]; !present {
		rm[pos] = 0
	}
	return rm[pos]
}

// The value that should be returned by a function parameter depends on the
// mode, which is derived from its position and the opcode.
func getParams(registers registerMap, pos, relativeBase int) []int {
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

// Used to respond to clone requests with a copy of the current state of this computer
func clone(this *Computer) *Computer {
	newRegisters := make(map[int]int)
	for k, v := range this.Registers {
		newRegisters[k] = v
	}

	newInput := make([]int, len(this.Input))
	copy(newInput, this.Input)

	newOutput := make([]int, len(this.output))
	copy(newOutput, this.output)

	new := Computer{
		Registers: newRegisters,
		Input:     newInput,
		Channels:  Channels{},

		instrPtr:     this.instrPtr,
		relativeBase: this.relativeBase,
		output:       newOutput}

	return &new
}

// Channels required for a computer, grouped together to tidy up args
// All channels nillable, though the two clone channels depend on each other
type Channels struct {
	// Basic IO
	Input  <-chan int
	Output chan<- int
	// Reports on
	Status    chan<- bool
	CloneReq  <-chan bool
	CloneResp chan<- *Computer
	Kill      <-chan bool
}

// Computer encapsulates the state of an intcode computer
type Computer struct {
	Registers registerMap
	Input     []int
	Channels  Channels

	instrPtr     int
	relativeBase int
	output       []int
}

func closeOutChannels(comp *Computer) {
	if comp.Channels.Output != nil {
		close(comp.Channels.Output)
	}
	if comp.Channels.Status != nil {
		close(comp.Channels.Status)
	}
}

// Extra goroutine to handle cloning
func cloneRoutine(comp *Computer, kill <-chan bool) {
	for {
		select {
		case <-kill:
			return
		case <-comp.Channels.CloneReq:
			comp.Channels.CloneResp <- clone(comp)
		}
	}
}

func cleanupCloner(cloneResp chan<- *Computer, killChan chan<- bool) {
	killChan <- true
	close(cloneResp)
	close(killChan)
}

// Core computation loop
func execute(comp *Computer) []int {
	defer closeOutChannels(comp)

	if comp.Channels.CloneReq != nil {
		if comp.Channels.CloneResp == nil {
			panic("Clone requests require clone responses!")
		}

		killChan := make(chan bool)
		go cloneRoutine(comp, killChan)
		defer cleanupCloner(comp.Channels.CloneResp, killChan)
	}

	for comp.Registers.get(comp.instrPtr)%100 != 99 {
		opcode := comp.Registers.get(comp.instrPtr) % 100
		params := getParams(comp.Registers, comp.instrPtr, comp.relativeBase)

		switch opcode {
		case 1:
			comp.Registers[params[2]] = params[1] + params[0]
		case 2:
			comp.Registers[params[2]] = params[1] * params[0]
		case 3:
			if len(comp.Input) > 0 {
				comp.Registers[params[0]] = comp.Input[0]
				comp.Input = comp.Input[1:]
			} else if comp.Channels.Input != nil {
				if comp.Channels.Status != nil {
					comp.Channels.Status <- true
				}
				comp.Registers[params[0]] = <-comp.Channels.Input
			} else {
				panic("No further input available")
			}
		case 4:
			comp.output = append(comp.output, params[0])
			if comp.Channels.Output != nil {
				comp.Channels.Output <- params[0]
			}
		case 5:
			if params[0] != 0 {
				comp.instrPtr = params[1]
				continue
			}
		case 6:
			if params[0] == 0 {
				comp.instrPtr = params[1]
				continue
			}
		case 7:
			if params[0] < params[1] {
				comp.Registers[params[2]] = 1
			} else {
				comp.Registers[params[2]] = 0
			}
		case 8:
			if params[0] == params[1] {
				comp.Registers[params[2]] = 1
			} else {
				comp.Registers[params[2]] = 0
			}
		case 9:
			comp.relativeBase += params[0]
		}
		comp.instrPtr += (len(params) + 1)
	}

	return comp.output
}

// ExecuteProgram from the beginning and return its output with the computer. Will modify registers.
func ExecuteProgram(registers []int, input []int, channels Channels) ([]int, *Computer) {
	instrPtr := 0
	relativeBase := 0
	outputs := []int{}

	computer := Computer{arrayToMap(registers), input, channels, instrPtr, relativeBase, outputs}

	return execute(&computer), &computer
}

// ResumeExecution for an existing intcode computer, attaching it to new channels
// Preserves the current state of execution
// This is used for testing multiple inputs from a given point in an intcode program
func (comp *Computer) ResumeExecution(channels Channels) {
	comp.Channels = channels

	execute(comp)
}
