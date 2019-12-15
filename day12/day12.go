package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"../utils"
)

type object struct {
	position position
	velocity velocity
}

type position struct {
	x int
	y int
	z int
}

type velocity struct {
	x int
	y int
	z int
}

func potentialEnergy(o object) int {
	return int(math.Abs(float64(o.position.x)) + math.Abs(float64(o.position.y)) + math.Abs(float64(o.position.z)))
}

func kineticEnergy(o object) int {
	return int(math.Abs(float64(o.velocity.x)) + math.Abs(float64(o.velocity.y)) + math.Abs(float64(o.velocity.z)))
}

func totalEnergy(objects []object) int {
	energy := 0
	for _, o := range objects {
		energy += potentialEnergy(o) * kineticEnergy(o)
	}
	return energy
}

func nextVelocity(o object, objects []object) velocity {
	x, y, z := o.velocity.x, o.velocity.y, o.velocity.z
	for _, other := range objects {
		if other != o {
			if o.position.x > other.position.x {
				x--
			} else if o.position.x < other.position.x {
				x++
			}
			if o.position.y > other.position.y {
				y--
			} else if o.position.y < other.position.y {
				y++
			}
			if o.position.z > other.position.z {
				z--
			} else if o.position.z < other.position.z {
				z++
			}
		}
	}
	return velocity{x, y, z}
}

func simulateMotion(initialObjects []object, n int) []object {
	objects := initialObjects
	for i := 0; i < n; i++ {
		newObjects := make([]object, len(objects))
		for i, o := range objects {
			v := nextVelocity(o, objects)
			p := position{o.position.x + v.x, o.position.y + v.y, o.position.z + v.z}
			newObjects[i] = object{position: p, velocity: v}
		}
		objects = newObjects
	}
	return objects
}

func findRepeatedState(initialObjects []object) int {
	objects := initialObjects
	xSet, ySet, zSet := map[string]struct{}{}, map[string]struct{}{}, map[string]struct{}{}

	xFound, yFound, zFound := false, false, false
	var xAt, yAt, zAt int
	for iters := 0; !xFound || !yFound || !zFound; iters++ {
		if !xFound {
			xs := make([]int, 2*len(objects))
			for n, o := range objects {
				xs[2*n] = o.position.x
				xs[2*n+1] = o.velocity.x
			}
			rep := fmt.Sprintf("%v", xs)
			if _, present := xSet[rep]; present {
				xFound = true
				xAt = iters
			} else {
				xSet[rep] = struct{}{}
			}
		}

		if !yFound {
			ys := make([]int, 2*len(objects))
			for n, o := range objects {
				ys[2*n] = o.position.y
				ys[2*n+1] = o.velocity.y
			}
			rep := fmt.Sprintf("%v", ys)
			if _, present := ySet[rep]; present {
				yFound = true
				yAt = iters
			} else {
				ySet[rep] = struct{}{}
			}
		}

		if !zFound {
			zs := make([]int, 2*len(objects))
			for n, o := range objects {
				zs[2*n] = o.position.z
				zs[2*n+1] = o.velocity.z
			}
			rep := fmt.Sprintf("%v", zs)
			if _, present := zSet[rep]; present {
				zFound = true
				zAt = iters
			} else {
				zSet[rep] = struct{}{}
			}
		}

		newObjects := make([]object, len(objects))
		for i, o := range objects {
			v := nextVelocity(o, objects)
			p := position{o.position.x + v.x, o.position.y + v.y, o.position.z + v.z}
			newObjects[i] = object{position: p, velocity: v}
		}
		objects = newObjects
	}

	return leastCommonMultiple([]int{xAt, yAt, zAt})
}

// GCD and LCM algorithms found here: https://play.golang.org/p/SmzvkDjYlb
// Thank you anonymous hero!
func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func leastCommonMultiple(xs []int) int {
	result := xs[0] * xs[1] / greatestCommonDivisor(xs[0], xs[1])

	for i := 2; i < len(xs); i++ {
		result = leastCommonMultiple([]int{result, xs[i]})
	}

	return result
}

func main() {
	input := utils.GetInputLines("input.txt")
	moons := []object{}
	for _, line := range input {
		r := regexp.MustCompile(`<x=(-?\d+), y=(-?\d+), z=(-?\d+)>`)
		captureGroups := r.FindStringSubmatch(line)
		x, _ := strconv.Atoi(captureGroups[1])
		y, _ := strconv.Atoi(captureGroups[2])
		z, _ := strconv.Atoi(captureGroups[3])
		moons = append(moons, object{position{x, y, z}, velocity{0, 0, 0}})
	}

	finalState := simulateMotion(moons, 1000)
	fmt.Printf("The answer to part one is %d\n", totalEnergy(finalState))

	timeUntilRepeat := findRepeatedState(moons)
	fmt.Printf("The answer to part two is %d\n", timeUntilRepeat)
}
