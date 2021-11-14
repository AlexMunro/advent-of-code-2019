package main

import (
	"fmt"
	"unicode"
)

type pathNode struct {
	position rune
	distance int
	visited  map[rune]struct{}
	keyCount int
	index    int // used by container/heap
}

func (p pathNode) hash() string {
	encodedKeys := encodeKeys(p.visited)

	return fmt.Sprint(p.position, '_', encodedKeys)
}

func (p *pathNode) GetPriority() int {
	return p.distance
}

func (p *pathNode) SetIndex(i int) {
	p.index = i
}

type quadPathNode struct {
	positions []rune
	distance  int
	visited   map[rune]struct{}
	keyCount  int
	index     int // used by container/heap
}

func (q quadPathNode) hash() string {
	encodedKeys := encodeKeys(q.visited)

	return fmt.Sprint(q.positions, '_', encodedKeys)
}

func encodeKeys(keys map[rune]struct{}) int {
	encodedKeys := 0

	for key := range keys {
		if unicode.IsLower(key) {
			position := int(key - 'a')
			shiftedPosition := 1 << position
			encodedKeys = encodedKeys | shiftedPosition
		}
	}
	return encodedKeys
}

func (q *quadPathNode) GetPriority() int {
	return q.distance
}

func (q *quadPathNode) SetIndex(i int) {
	q.index = i
}
