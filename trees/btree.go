package trees

type BinaryTree struct {
	Root *BinaryNode
}

type BinaryNode struct {
	Data  int
	Left  *BinaryNode
	Right *BinaryNode
}

func (bt *BinaryTree) Insert(data int) {
	if bt.Root == nil {
		bt.Root = &BinaryNode{Data: data}
	} else {
		bt.insertNode(bt.Root, data)
	}
}

func (bt *BinaryTree) insertNode(node *BinaryNode, data int) {
	if data < node.Data {
		if node.Left == nil {
			node.Left = &BinaryNode{Data: data}
		} else {
			bt.insertNode(node.Left, data)
		}
	} else {
		if node.Right == nil {
			node.Right = &BinaryNode{Data: data}
		} else {
			bt.insertNode(node.Right, data)
		}
	}
}

func walkPreOrder(curr *BinaryNode, path *[]int) []int {
	if curr == nil {
		return *path
	}
	*path = append(*path, curr.Data)
	walkPreOrder(curr.Left, path)
	walkPreOrder(curr.Right, path)
	return *path
}

func (bt *BinaryTree) PreOrderTraversal() []int {
	return walkPreOrder(bt.Root, &[]int{})
}

func walkInOrder(curr *BinaryNode, path *[]int) []int {
	if curr == nil {
		return *path
	}
	walkInOrder(curr.Left, path)
	*path = append(*path, curr.Data)
	walkInOrder(curr.Right, path)
	return *path
}

func (bt *BinaryTree) InOrderTraversal() []int {
	return walkInOrder(bt.Root, &[]int{})
}

func walkPostOrder(curr *BinaryNode, path *[]int) []int {
	if curr == nil {
		return *path
	}
	walkPostOrder(curr.Left, path)
	walkPostOrder(curr.Right, path)
	*path = append(*path, curr.Data)
	return *path
}

func (bt *BinaryTree) PostOrderTraversal() []int {
	return walkPostOrder(bt.Root, &[]int{})
}
