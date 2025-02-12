package trees

import (
	"robtrizzo/data-and-algo-practice/internal/assert"
	"testing"
)

func constructBinaryTree() BinaryTree {
	bt := BinaryTree{}
	bt.Insert(5)
	bt.Insert(3)
	bt.Insert(2)
	bt.Insert(4)
	bt.Insert(7)
	bt.Insert(6)
	bt.Insert(8)
	return bt
}

func TestInsert(t *testing.T) {
	bt := BinaryTree{}
	bt.Insert(5)
	assert.Equal(t, bt.Root.Data, 5)
	bt.Insert(4)
	assert.Equal(t, bt.Root.Left.Data, 4)
	bt.Insert(6)
	assert.Equal(t, bt.Root.Right.Data, 6)
}

func TestPreOrderTraversal(t *testing.T) {
	bt := constructBinaryTree()
	expected := []int{5, 3, 2, 4, 7, 6, 8}
	actual := bt.PreOrderTraversal()
	assert.ArraysEqual(t, actual, expected)
}

func TestInOrderTraversal(t *testing.T) {
	bt := constructBinaryTree()
	expected := []int{2, 3, 4, 5, 6, 7, 8}
	actual := bt.InOrderTraversal()
	assert.ArraysEqual(t, actual, expected)
}

func TestPostOrderTraversal(t *testing.T) {
	bt := constructBinaryTree()
	expected := []int{2, 4, 3, 6, 8, 7, 5}
	actual := bt.PostOrderTraversal()
	assert.ArraysEqual(t, actual, expected)
}
