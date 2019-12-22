package main

import (
	"../intcode"
	"../utils"
	"../utils/intset"

	"fmt"
	"sync"
)

func orderingRec(previous []int, remaining *intset.IntSet) [][]int {
	orderings := [][]int{}
	for _, i := range remaining.ToSlice() {
		nextOrdering := append(previous, i)

		nextRemaining := remaining.Clone()
		nextRemaining.Remove(i)
		if nextRemaining.IsEmpty() {
			orderings = append(orderings, nextOrdering)
		} else {
			for _, o := range orderingRec(nextOrdering, nextRemaining) {
				orderings = append(orderings, o)
			}
		}
	}
	return orderings
}

func allPhaseOrderings(phases []int) [][]int {
	phaseSet := intset.FromSlice(phases)
	return orderingRec([]int{}, phaseSet)
}

func execute(ordering []int, registers []int) int {
	signal := 0
	for _, phase := range ordering {
		r := utils.CopyInts(registers)
		output := intcode.ExecuteProgram(r, []int{phase, signal}, nil, nil, nil)
		signal = output[len(output)-1]
	}
	return signal
}

func executeLooping(phases []int, registers []int) int {
	channels := []chan int{}
	for i := 0; i < len(phases); i++ {
		channels = append(channels, make(chan int))
	}

	var wg sync.WaitGroup
	wg.Add(len(channels) - 1)

	for i, phase := range phases {
		r := utils.CopyInts(registers)
		go func(i, phase int, r []int) {
			initialInput := []int{phase}
			if i == 0 {
				initialInput = append(initialInput, 0)
			}
			if i < len(phases)-1 { // The last process will write once after the others have finished
				defer wg.Done()
			}
			intcode.ExecuteProgram(r, initialInput, channels[i], channels[(i+1)%len(channels)], nil)
		}(i, phase, r)
	}
	wg.Wait()
	return <-channels[0]
}

func maxSignal(phases []int, registers []int, exec func([]int, []int) int) int {
	orderings := allPhaseOrderings(phases)
	maxOrdering := 0

	for _, o := range orderings {
		resultSignal := exec(o, registers)
		if resultSignal > maxOrdering {
			maxOrdering = resultSignal
		}
	}

	return maxOrdering
}

func main() {
	registers := utils.GetCommaSeparatedInts("input.txt")
	maxSequentialSignal := maxSignal([]int{0, 1, 2, 3, 4}, registers, execute)

	fmt.Printf("The answer to part one is %d\n", maxSequentialSignal)

	maxLoopingSignal := maxSignal([]int{5, 6, 7, 8, 9}, registers, executeLooping)
	fmt.Printf("The answer to part two is %d\n", maxLoopingSignal)
}
