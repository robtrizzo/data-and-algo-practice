package trees

func compare(a *BinaryNode, b *BinaryNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Data != b.Data {
		return false
	}

	return compare(a.Left, b.Left) && compare(a.Right, b.Right)
}

func (bt *BinaryTree) Equals(otherTree *BinaryTree) bool {
	return compare(bt.Root, otherTree.Root)
}
