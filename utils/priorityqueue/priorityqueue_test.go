package priorityqueue

import "testing"

type testQueueElem struct {
	priority int
	index    int
}

func (t *testQueueElem) GetPriority() int {
	return t.priority
}

func (t *testQueueElem) SetIndex(index int) {
	t.index = index
}

func TestPriorityQueue(t *testing.T) {
	queue := New(&testQueueElem{priority: 5})

	queue.Add(&testQueueElem{priority: 2})
	queue.Add(&testQueueElem{priority: 8})
	queue.Add(&testQueueElem{priority: 5})

	two := queue.Next().(*testQueueElem)
	if two.priority != 2 {
		t.Errorf("Expected 2, got %v", two.priority)
	}

	five := queue.Next().(*testQueueElem)
	if five.priority != 5 {
		t.Errorf("Expected 5, got %v", five.priority)
	}

	fiveAgain := queue.Next().(*testQueueElem)
	if fiveAgain.priority != 5 {
		t.Errorf("Expected 5, got %v", fiveAgain.priority)
	}

	eight := queue.Next().(*testQueueElem)
	if eight.priority != 8 {
		t.Errorf("Expected 8, got %v", eight.priority)
	}
}
