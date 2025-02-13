package trees

import (
	"robtrizzo/data-and-algo-practice/internal/assert"
	"testing"
)

func TestBSTInsert(t *testing.T) {
	bst := BinarySearchTree{}
	bst.Insert(15)
	assert.Equal[int](t, bst.Root.Data, 15)
	bst.Insert(7)
	assert.Equal(t, bst.Root.Left.Data, 7)
	bst.Insert(9)
	assert.Equal(t, bst.Root.Left.Right.Data, 9)
}

func TestBSTFind(t *testing.T) {
	bst := BinarySearchTree{}
	bst.Insert(15)
	bst.Insert(7)
	bst.Insert(9)
	assert.Equal(t, bst.Find(5), false)
	assert.Equal(t, bst.Find(7), true)
}

func TestBSTDelete(t *testing.T) {
	bst := BinarySearchTree{}
	bst.Insert(15)
	bst.Delete(15)
	assert.Equal(t, bst.Root, nil)
	bst.Insert(15)
	bst.Insert(7)
	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(51)
	bst.Insert(25)
	bst.Insert(37)
	bst.Insert(100)
	bst.Insert(80)
	bst.Delete(1)
	assert.Equal(t, bst.Root.Left.Left.Left, nil)
	bst.Delete(7)
	assert.Equal(t, bst.Root.Left.Data, 3)
	bst.Delete(51)
	assert.Equal(t, bst.Root.Right.Data, 80)
}
