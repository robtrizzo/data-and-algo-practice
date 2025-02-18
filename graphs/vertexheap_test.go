package graphs

import (
	"robtrizzo/data-and-algo-practice/internal/assert"
	"testing"
)

func TestVertexHeapInsert(t *testing.T) {
	h := NewVertexMinHeap()

	h.Insert(vertex{val: 1, distance: 50})
	assert.Equal[int](t, h.Data[0].distance, 50)
	assert.Equal[int](t, h.Length, 1)
	h.Insert(vertex{val: 1, distance: 71})
	h.Insert(vertex{val: 1, distance: 100})
	h.Insert(vertex{val: 1, distance: 101})
	h.Insert(vertex{val: 1, distance: 80})
	h.Insert(vertex{val: 1, distance: 200})
	h.Insert(vertex{val: 1, distance: 101})
	assert.Equal[int](t, h.Data[3].distance, 101)
	assert.Equal(t, h.Length, 7)
}

func TestVertexHeapDelete(t *testing.T) {
	h := NewVertexMinHeap()

	h.Insert(vertex{val: 1, distance: 50})
	h.Insert(vertex{val: 1, distance: 71})
	h.Insert(vertex{val: 1, distance: 100})
	h.Insert(vertex{val: 1, distance: 101})
	h.Insert(vertex{val: 1, distance: 80})
	h.Insert(vertex{val: 1, distance: 200})
	h.Insert(vertex{val: 1, distance: 101})
	actual, err := h.Delete()
	assert.NilError(t, err)
	assert.Equal(t, actual.distance, 50)
	assert.Equal(t, h.Length, 6)
	assert.Equal(t, h.Data[0].distance, 71)
	assert.Equal(t, h.Data[5].distance, 200)
}
