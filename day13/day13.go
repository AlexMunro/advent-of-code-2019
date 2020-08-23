package main

import (
	"fmt"

	"../intcode"
	"../utils"
)

type tile int

const (
	block  tile = 2
	paddle tile = 3
	ball   tile = 4
)

func countBlocks(gameState []int) int {
	count := 0
	for i := 2; i < len(gameState); i += 3 {
		if tile(gameState[i]) == block {
			count++
		}
	}
	return count
}

type joystick int

const (
	left    joystick = -1
	neutral joystick = 0
	right   joystick = 1
)

func playGame(registers, program []int) int {
	gameInput := make(chan int)
	gameOutput := make(chan int)
	gameStatus := make(chan bool)

	go func() {
		intcode.ExecuteProgram(registers, nil, intcode.Channels{gameInput, gameOutput, gameStatus, nil, nil, nil})
	}()

	var score, paddleX, ballX int

	for true {
		select {
		case x := <-gameOutput:
			y := <-gameOutput
			id := <-gameOutput

			if x == -1 && y == 0 {
				score = id
				continue
			}

			switch tile(id) {
			case paddle:
				paddleX = x
			case ball:
				ballX = x
			}
		case waiting := <-gameStatus:
			if waiting {
				if ballX > paddleX {
					gameInput <- int(right)
				} else if ballX < paddleX {
					gameInput <- int(left)
				} else {
					gameInput <- int(neutral)
				}
			} else {
				return score
			}
		}
	}
	panic("Again, the compiler forces me to add this silly line")
}

func main() {
	testRegisters := utils.GetCommaSeparatedInts("input.txt")
	program := utils.CopyInts(testRegisters)

	// probably not necessary, but a headache I'd rather avoid
	playRegisters := utils.CopyInts(testRegisters)

	output, _ := intcode.ExecuteProgram(testRegisters, program, intcode.Channels{nil, nil, nil, nil, nil, nil})
	blocks := countBlocks(output)

	fmt.Printf("The answer to part one is %d\n", blocks)

	playRegisters[0] = 2
	playProgram := utils.CopyInts(playRegisters)

	finalScore := playGame(playRegisters, playProgram)
	fmt.Printf("The answer to part two is %d\n", finalScore)
}
