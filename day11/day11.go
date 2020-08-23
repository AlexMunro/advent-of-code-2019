package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"sync"

	"../intcode"
	"../utils"

	// I know what I'm doing, go-lint. Dot imports are fine.
	. "../utils/location"
	"golang.org/x/image/bmp"
)

type colour int

const (
	black colour = 0
	white colour = 1
)

type direction int

const (
	up    direction = 0
	right direction = 1
	down  direction = 2
	left  direction = 3
)

func paintRobot(paintMap map[Location]colour, inChan <-chan int, outChan chan<- int) {
	defer close(outChan)

	// Alternate betwen even paint colours and odd directions
	evenInstruction := true
	position := Location{X: 0, Y: 0}
	orientation := up

	if cameraColour, present := paintMap[position]; present {
		outChan <- int(cameraColour)
	} else {
		outChan <- int(black)
	}

	for instruction := range inChan {
		if evenInstruction {
			paintMap[position] = colour(instruction)
		} else {
			if instruction == 0 {
				orientation = direction((orientation + 3) % 4) // saves worrying about negative modulo
			} else {
				orientation = direction((orientation + 1) % 4)
			}
			switch orientation {
			case up:
				position = Location{X: position.X, Y: position.Y + 1}
			case right:
				position = Location{X: position.X + 1, Y: position.Y}
			case down:
				position = Location{X: position.X, Y: position.Y - 1}
			case left:
				position = Location{X: position.X - 1, Y: position.Y}
			}

			if cameraColour, present := paintMap[position]; present {
				outChan <- int(cameraColour)
			} else {
				outChan <- int(black)
			}
		}
		evenInstruction = !evenInstruction
	}
}

func paintTiles(program []int, startOnWhiteTile bool) map[Location]colour {
	var wg sync.WaitGroup
	wg.Add(2)
	programToRobot := make(chan int)
	robotToProgram := make(chan int, 2)

	paintMap := map[Location]colour{}
	if startOnWhiteTile {
		paintMap[Location{X: 0, Y: 0}] = white
	}

	// program
	go func() {
		defer wg.Done()
		registers := utils.CopyInts(program)
		intcode.ExecuteProgram(registers, []int{}, robotToProgram, programToRobot, nil)
	}()

	// robot
	go func() {
		defer wg.Done()
		paintRobot(paintMap, programToRobot, robotToProgram)
	}()

	wg.Wait()
	return paintMap
}

func drawIdentifier(paintMap map[Location]colour, filename string) {
	locSlice := []Location{}
	for loc := range paintMap {
		locSlice = append(locSlice, loc)
	}

	minX := MinX(locSlice)
	maxX := MaxX(locSlice)
	minY := MinY(locSlice)
	maxY := MaxY(locSlice)

	drawing := image.NewRGBA(image.Rect(minX, minY, maxX+1, maxY+1))

	for loc, colour := range paintMap {
		if colour == white {
			drawing.Set(loc.X, maxY-loc.Y+minY, color.White)
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bmp.Encode(file, drawing)
}

func main() {
	registers := utils.GetCommaSeparatedInts("input.txt")
	locations := paintTiles(registers, false)

	fmt.Printf("The answer to part one is %d\n", len(locations))

	identifier := paintTiles(registers, true)
	drawIdentifier(identifier, "answer2.bmp")

	fmt.Println("The answer to part two can be found in answer2.bmp")
}
