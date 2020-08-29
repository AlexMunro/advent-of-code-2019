package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"../intcode"
	"../utils"
	"golang.org/x/image/bmp"

	// again, sod off go-lint
	. "../utils/location"
)

type status = int

const (
	hitWall status = 0
	moved   status = 1
	goal    status = 2
)

// Need to store channels separately from the computer to not get caught by direction types
type searchChannels struct {
	input     chan int
	output    chan int
	cloneReq  chan bool
	cloneResp chan *intcode.Computer
	kill      chan bool
}

func (c searchChannels) convert() intcode.Channels {
	return intcode.Channels{
		Input:     c.input,
		Output:    c.output,
		CloneReq:  c.cloneReq,
		CloneResp: c.cloneResp,
		Kill:      c.kill}
}

type searchNode = struct {
	loc      Location
	depth    int
	channels searchChannels
}

func visualise(maze map[Location]status, filename string) {
	locSlice := make([]Location, len(maze))
	i := 0
	for loc := range maze {
		locSlice[i] = loc
		i++
	}

	minX := MinX(locSlice)
	maxX := MaxX(locSlice)
	minY := MinY(locSlice)
	maxY := MaxY(locSlice)

	drawing := image.NewRGBA(image.Rect(minX, minY, maxX+1, maxY+1))

	for loc, stat := range maze {
		drawX := maxX - loc.X + minX
		drawY := maxY - loc.Y + minY

		switch stat {
		case hitWall:
			drawing.Set(drawX, drawY, color.RGBA{0, 255, 0, 1})
		case moved:
			drawing.Set(drawX, drawY, color.Black)
		case goal:
			drawing.Set(drawX, drawY, color.RGBA{0, 255, 255, 1})
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bmp.Encode(file, drawing)
}

func breadthFirstSearch(input []int) int {
	root := searchNode{
		loc:      Location{X: 0, Y: 0},
		depth:    0,
		channels: searchChannels{make(chan int), make(chan int), make(chan bool), make(chan *intcode.Computer), make(chan bool)}}

	go intcode.ExecuteProgram(input, []int{}, root.channels.convert())

	exploredNodes := map[Location]status{Location{X: 0, Y: 0}: moved}

	nodes := []searchNode{root}

	for len(nodes) > 0 {
		for _, node := range nodes {
			nodes = nodes[1:]
			for _, direction := range []Direction{North, South, West, East} {
				loc := node.loc.Head(direction)
				if _, present := exploredNodes[loc]; present {
					continue
				}

				node.channels.cloneReq <- true
				newComputer := <-node.channels.cloneResp

				// I have no idea why the input channel sometimes needs a buffer, but here we are
				newChannels := searchChannels{make(chan int, 1), make(chan int), make(chan bool), make(chan *intcode.Computer), make(chan bool)}
				go newComputer.ResumeExecution(newChannels.convert())

				newChannels.input <- direction
				nextStatus := <-newChannels.output

				exploredNodes[loc] = nextStatus
				visualise(exploredNodes, "bfs.bmp")

				switch nextStatus {
				case hitWall:
					newChannels.kill <- true
					continue
				case moved:
					nextNode := searchNode{
						loc:      loc,
						depth:    node.depth + 1,
						channels: newChannels,
					}
					nodes = append(nodes, nextNode)
				case goal:
					return node.depth + 1
				}
			}
			node.channels.kill <- true
		}
	}

	panic("Solution not found")
}

type encodedSearchNode = struct {
	loc   Location
	depth int
	path  []Direction
}

func main() {
	input := utils.GetCommaSeparatedInts("input.txt")
	fmt.Printf("The answer to part one is %d\n", breadthFirstSearch(input))
}
