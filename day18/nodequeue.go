package main

import "container/heap"

// Priority queue code mostly lifted from golang docs: https://golang.org/pkg/container/heap/
type nodeQueue []*pathNode

func newNodeQueue(initial *pathNode) *nodeQueue {
	nodes := make(nodeQueue, 1)
	nodes[0] = initial
	heap.Init(&nodes)
	return &nodes
}

func (nq nodeQueue) Len() int {
	return len(nq)
}

func (nq nodeQueue) Less(i, j int) bool {
	return nq[i].distance < nq[j].distance
}

func (nq nodeQueue) Swap(i, j int) {
	nq[i], nq[j] = nq[j], nq[i]
	nq[i].index = i
	nq[j].index = j
}

func (nq *nodeQueue) Push(x interface{}) {
	n := len(*nq)
	node := x.(*pathNode)
	node.index = n
	*nq = append(*nq, node)
}

func (nq *nodeQueue) Pop() interface{} {
	old := *nq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
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
