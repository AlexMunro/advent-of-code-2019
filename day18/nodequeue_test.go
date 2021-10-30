package main

import "testing"

func TestNodeQueue(t *testing.T) {
	queue := newNodeQueue(&pathNode{distance: 5})

	queue.add(&pathNode{distance: 2})
	queue.add(&pathNode{distance: 8})
	queue.add(&pathNode{distance: 5})

	two := queue.next().(*pathNode)
	if two.distance != 2 {
		t.Errorf("Expected 2, got %v", two.distance)
	}

	five := queue.next().(*pathNode)
	if five.distance != 5 {
		t.Errorf("Expected 5, got %v", five.distance)
	}

	fiveAgain := queue.next().(*pathNode)
	if fiveAgain.distance != 5 {
		t.Errorf("Expected 5, got %v", fiveAgain.distance)
	}

	eight := queue.next().(*pathNode)
	if eight.distance != 8 {
		t.Errorf("Expected 8, got %v", eight.distance)
	}
}
