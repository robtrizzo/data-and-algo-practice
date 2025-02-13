package trees

type BinarySearchTree struct {
	Root *BinaryNode
}

func find(node *BinaryNode, needle int) bool {
	if node == nil {
		return false
	}
	if node.Data == needle {
		return true
	}
	if node.Data < needle {
		return find(node.Right, needle)
	} else {
		return find(node.Left, needle)
	}
}

func (bst *BinarySearchTree) Find(needle int) bool {
	return find(bst.Root, needle)
}

func insert(node *BinaryNode, val int) *BinaryNode {
	if node == nil {
		return &BinaryNode{Data: val}
	}

	if node.Data < val {
		node.Right = insert(node.Right, val)
	} else {
		node.Left = insert(node.Left, val)
	}

	return node
}

func (bst *BinarySearchTree) Insert(val int) {
	bst.Root = insert(bst.Root, val)
}

func findMin(node *BinaryNode) *BinaryNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func delete(node *BinaryNode, val int) *BinaryNode {
	if node == nil {
		return nil
	}

	if node.Data < val {
		node.Right = delete(node.Right, val)
	} else if node.Data > val {
		node.Left = delete(node.Left, val)
	} else {
		if node.Right == nil {
			return node.Left
		} else if node.Left == nil {
			return node.Right
		}

		min := findMin(node.Right)
		node.Data = min.Data
		node.Right = delete(node.Right, min.Data)
	}
	return node
}

func (bst *BinarySearchTree) Delete(val int) {
	bst.Root = delete(bst.Root, val)
}
