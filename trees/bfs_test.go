package trees

import (
	"robtrizzo/data-and-algo-practice/internal/assert"
	"testing"
)

func TestBFS(t *testing.T) {
	bt := constructBinaryTree()
	found := BFS(*bt.Root, 5)
	assert.Equal(t, found, true)
	found = BFS(*bt.Root, 1)
	assert.Equal(t, found, false)
}
