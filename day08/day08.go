package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"strings"

	"golang.org/x/image/bmp"

	"../utils"
)

func getImageLayers(input string, width, height int) [][]string {
	layers := make([][]string, len(input)/(width*height))
	for i := range layers {
		offset := i * width * height
		lines := make([]string, height)
		for j := 0; j < height; j++ {
			lines[j] = input[offset+(j*width) : offset+((j+1)*width)]
		}
		layers[i] = lines
	}
	return layers
}

func digitCounts(layer []string) map[rune]int {
	digitMap := map[rune]int{}

	for _, str := range layer {
		for _, char := range str {
			n, present := digitMap[char]
			if present {
				digitMap[char] = n + 1
			} else {
				digitMap[char] = 1
			}
		}
	}

	return digitMap
}

func layerWithLeastZeros(layers [][]string) ([]string, map[rune]int) {
	layerMaps := make([]map[rune]int, len(layers))
	for i, layer := range layers {
		layerMaps[i] = digitCounts(layer)
	}

	fewestZeroIdx := 0
	for i, layerMap := range layerMaps {
		if layerMap['0'] < layerMaps[fewestZeroIdx]['0'] {
			fewestZeroIdx = i
		}
	}
	return layers[fewestZeroIdx], layerMaps[fewestZeroIdx]
}

func resolvedLayer(layers [][]string) []string {
	height, width := len(layers[0]), len(layers[0][0])
	resolvedLayer := []string{}

	// Initialise it as a transparent layer
	for i := 0; i < height; i++ {
		resolvedLayer = append(resolvedLayer, strings.Repeat("2", width))
	}

	// Iterate through layers and show the forward-most pixel
	for _, layer := range layers {
		for y, line := range layer {
			for x, char := range line {
				if resolvedLayer[y][x] == '2' && char != '2' {
					resolvedLayer[y] = resolvedLayer[y][:x] + string(char) + resolvedLayer[y][x+1:]
				}
			}
		}
	}

	return resolvedLayer
}

func drawLayer(layer []string, width, height int, filename string) {
	drawing := image.NewRGBA(image.Rect(0, 0, width, height))
	for y, row := range layer {
		for x, pixel := range row {
			if pixel == '1' {
				drawing.Set(x, y, color.Black)
			} else if pixel == '0' {
				drawing.Set(x, y, color.White)
			} else {
				drawing.Set(x, y, color.RGBA{0, 255, 0, 255})
			}
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
	input := utils.GetInputSingleString("input.txt")
	layers := getImageLayers(input, 25, 6)
	_, testMap := layerWithLeastZeros(layers)
	testResult := testMap['1'] * testMap['2']
	fmt.Printf("The answer to part one is %d\n", testResult)
	drawLayer(resolvedLayer(layers), 25, 6, "answer2.bmp")
	fmt.Printf("The answer to part two can be found in answer2.bmp\n")
}
