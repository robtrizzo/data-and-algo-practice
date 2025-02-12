package queues

import (
	"robtrizzo/data-and-algo-practice/internal/assert"
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := Queue[int]{}
	assert.Equal(t, 0, q.length)
	q.Enqueue(1)
	assert.Equal(t, 1, q.length)
	assert.Equal(t, 1, q.head.value)
	assert.Equal(t, 1, q.tail.value)
}

func TestPeek(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	val, err := q.Peek()
	assert.NilError(t, err)
	assert.Equal(t, val, 1)
}

func TestDequeue(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)

	val, err := q.Dequeue()
	assert.NilError(t, err)
	assert.Equal(t, val, 1)
	assert.Equal(t, q.length, 1)

	val, err = q.Dequeue()
	assert.NilError(t, err)
	assert.Equal(t, val, 2)
	assert.Equal(t, q.length, 0)
}
