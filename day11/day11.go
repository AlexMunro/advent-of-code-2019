package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"sync"

	"../intcode"
	"../utils"
	"../utils/location"
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

func paintRobot(paintMap map[location.Location]colour, inChan <-chan int, outChan chan<- int) {
	defer close(outChan)

	// Alternate betwen even paint colours and odd directions
	evenInstruction := true
	position := location.Location{0, 0}
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
				position = location.Location{position.X, position.Y + 1}
			case right:
				position = location.Location{position.X + 1, position.Y}
			case down:
				position = location.Location{position.X, position.Y - 1}
			case left:
				position = location.Location{position.X - 1, position.Y}
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

func paintTiles(program []int, startOnWhiteTile bool) map[location.Location]colour {
	var wg sync.WaitGroup
	wg.Add(2)
	programToRobot := make(chan int)
	robotToProgram := make(chan int, 2)

	paintMap := map[location.Location]colour{}
	if startOnWhiteTile {
		paintMap[location.Location{0, 0}] = white
	}

	// program
	go func() {
		defer wg.Done()
		registers := utils.CopyInts(program)
		intcode.ExecuteProgram(registers, []int{}, robotToProgram, programToRobot)
	}()

	// robot
	go func() {
		defer wg.Done()
		paintRobot(paintMap, programToRobot, robotToProgram)
	}()

	wg.Wait()
	return paintMap
}

func drawIdentifier(paintMap map[location.Location]colour, filename string) {
	locSlice := []location.Location{}
	for loc := range paintMap {
		locSlice = append(locSlice, loc)
	}

	minX := location.MinX(locSlice)
	maxX := location.MaxX(locSlice)
	minY := location.MinY(locSlice)
	maxY := location.MaxY(locSlice)

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
