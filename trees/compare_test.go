package trees

import (
	"robtrizzo/data-and-algo-practice/internal/assert"
	"testing"
)

func TestCompareDiffValues(t *testing.T) {
	bta := BinaryTree{}
	bta.Insert(5)
	bta.Insert(4)
	bta.Insert(6)
	btb := BinaryTree{}
	btb.Insert(5)
	btb.Insert(4)
	btb.Insert(7)
	assert.Equal(t, bta.Equals(&btb), false)
}

func TestCompareSameValues(t *testing.T) {
	bta := BinaryTree{}
	bta.Insert(5)
	bta.Insert(4)
	bta.Insert(6)
	btb := BinaryTree{}
	btb.Insert(5)
	btb.Insert(4)
	btb.Insert(6)
	assert.Equal(t, bta.Equals(&btb), true)
}

func TestCompareDiffStructure(t *testing.T) {
	bta := BinaryTree{}
	bta.Insert(5)
	bta.Insert(4)
	bta.Insert(6)
	btb := BinaryTree{}
	btb.Insert(5)
	btb.Insert(4)
	btb.Root.Left.Right = &BinaryNode{Data: 6}
	assert.Equal(t, bta.Equals(&btb), false)
}
