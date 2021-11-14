package priorityqueue

import "container/heap"

// Priority queue code mostly lifted from golang docs: https://golang.org/pkg/container/heap/

type PriorityElem interface {
	GetPriority() int
	SetIndex(int)
}

type PriorityQueue []PriorityElem

func New(initial PriorityElem) *PriorityQueue {
	nodes := make(PriorityQueue, 1)
	nodes[0] = initial
	heap.Init(&nodes)
	return &nodes
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].GetPriority() < pq[j].GetPriority()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].SetIndex(i)
	pq[j].SetIndex(j)
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(PriorityElem)
	node.SetIndex(n)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.SetIndex(-1)
	*pq = old[0 : n-1]
	return node
}

// The _actual_ add method, since the other one is used by container/heap
func (pq *PriorityQueue) Add(x interface{}) {
	heap.Push(pq, x)
}

// The _actual_ pop method, since the other one is used by container/heap
func (pq *PriorityQueue) Next() interface{} {
	return heap.Pop(pq)
}
