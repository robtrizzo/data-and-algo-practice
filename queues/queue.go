package queues

import (
	"errors"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func (q *Queue[T]) Enqueue(val T) {
	node := Node[T]{value: val}
	if q.length == 0 {
		q.head = &node
		q.tail = &node
		q.length = 1
	} else {
		q.tail.next = &node
		q.tail = &node
		q.length++
	}
}

var ErrQueueEmpty = errors.New("cannot perform this operation on an empty queue")

func (q *Queue[T]) Dequeue() (T, error) {
	if q.head == nil {
		var zero T
		return zero, ErrQueueEmpty
	}

	q.length--

	head := q.head
	q.head = q.head.next

	head.next = nil

	return head.value, nil
}
func (q *Queue[T]) Peek() (T, error) {
	if q.head == nil {
		var zero T
		return zero, ErrQueueEmpty
	} else {
		return q.head.value, nil
	}
}
