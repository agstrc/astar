package pqueue

import "testing"

func TestPriorityQueue(t *testing.T) {
	var queue PriorityQueue[string]

	queue.Push("B value", 10)
	queue.Push("A value", 15)
	queue.Push("C value", 5)

	for _, str := range [...]string{"A value", "B value", "C value"} {
		popped := queue.Pop()
		if popped != str {
			t.Fatalf("Expected '%s', got '%s'", str, popped)
		}
	}

	if !queue.Empty() {
		t.Fatal("Expected queue to be empty")
	}
}
