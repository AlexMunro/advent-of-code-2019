package main

import "container/heap"

// Priority queue code mostly lifted from golang docs: https://golang.org/pkg/container/heap/

type genericPathNode interface {
	getDistance() int
	setIndex(int)
}

type nodeQueue []genericPathNode

func newNodeQueue(initial genericPathNode) *nodeQueue {
	nodes := make(nodeQueue, 1)
	nodes[0] = initial
	heap.Init(&nodes)
	return &nodes
}

func (nq nodeQueue) Len() int {
	return len(nq)
}

func (nq nodeQueue) Less(i, j int) bool {
	return nq[i].getDistance() < nq[j].getDistance()
}

func (nq nodeQueue) Swap(i, j int) {
	nq[i], nq[j] = nq[j], nq[i]
	nq[i].setIndex(i)
	nq[j].setIndex(j)
}

func (nq *nodeQueue) Push(x interface{}) {
	n := len(*nq)
	node := x.(genericPathNode)
	node.setIndex(n)
	*nq = append(*nq, node)
}

func (nq *nodeQueue) Pop() interface{} {
	old := *nq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.setIndex(-1)
	*nq = old[0 : n-1]
	return node
}

// The _actual_ add method, since the other one is used by container/heap
func (nq *nodeQueue) add(x interface{}) {
	heap.Push(nq, x)
}

// The _actual_ pop method, since the other one is used by container/heap
func (nq *nodeQueue) next() interface{} {
	return heap.Pop(nq)
}
