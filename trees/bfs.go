package trees

import (
	"robtrizzo/data-and-algo-practice/queues"
)

func BFS(head BinaryNode, needle int) bool {
	q := queues.Queue[BinaryNode]{}
	q.Enqueue(head)

	for q.Length > 0 {
		curr, err := q.Dequeue()
		if err != nil {
			return false
		}
		if curr.Data == needle {
			return true
		}

		if curr.Left != nil {
			q.Enqueue(*curr.Left)
		}
		if curr.Right != nil {
			q.Enqueue(*curr.Right)
		}
	}

	return false
}
