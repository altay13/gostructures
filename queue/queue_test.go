package queue

import "testing"

func TestEnqueue(t *testing.T) {
	q := NewQueue()

	q.Enqueue("Queue_1")
	q.Enqueue("Queue_2")
	q.Enqueue("Queue_3")

	if q.size != 3 {
		t.Errorf("Failed to Enqueue item. Expected 3, instead %d", q.size)
	}

	le := q.list.GetLastNode()
	if le.Item != "Queue_3" {
		t.Errorf("Failed to Enqueue item. Expected 'Queue_3', instead %s", le)
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue()

	q.Enqueue("Queue_1")
	q.Enqueue("Queue_2")
	q.Enqueue("Queue_3")

	n := q.Dequeue()

	if q.size != 2 {
		t.Errorf("Failed to Dequeue item. Expected 2, instead %d", q.size)
	}

	if n != "Queue_3" {
		t.Errorf("Failed to Dequeue item. Expected 'Queue_3', instead %s", n)
	}
}
