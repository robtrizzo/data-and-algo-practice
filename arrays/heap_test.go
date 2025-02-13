package arrays

import (
	"robtrizzo/data-and-algo-practice/internal/assert"
	"testing"
)

func TestHeapInsert(t *testing.T) {
	h := MinHeap{
		data:   []int{},
		Length: 0,
	}

	h.Insert(50)
	assert.Equal[int](t, h.data[0], 50)
	assert.Equal[int](t, h.Length, 1)
	h.Insert(71)
	h.Insert(100)
	h.Insert(101)
	h.Insert(80)
	h.Insert(200)
	h.Insert(101)
	assert.Equal[int](t, h.data[3], 101)
	assert.Equal(t, h.Length, 7)
}

func TestHeapDelete(t *testing.T) {
	h := MinHeap{
		data:   []int{},
		Length: 0,
	}
	h.Insert(50)
	h.Insert(71)
	h.Insert(100)
	h.Insert(101)
	h.Insert(80)
	h.Insert(200)
	h.Insert(101)
	actual, err := h.Delete()
	assert.NilError(t, err)
	assert.Equal(t, actual, 50)
	assert.Equal(t, h.Length, 6)
	assert.Equal(t, h.data[0], 71)
	assert.Equal(t, h.data[5], 200)
}
