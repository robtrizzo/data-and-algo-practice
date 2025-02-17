package graphs

import (
	"robtrizzo/data-and-algo-practice/queues"
	"slices"
)

type WeightedAdjacencyMatrix struct {
	vertices int
	matrix   [][]int
}

func (wam *WeightedAdjacencyMatrix) BFS(source int, needle int) ([]int, error) {
	seen := initSeen(wam.vertices)
	prev := initPrev(wam.vertices)

	seen[source] = true
	q := queues.Queue[int]{}
	q.Enqueue(source)

	for q.Length > 0 {
		curr, err := q.Dequeue()
		if err != nil {
			return []int{}, err
		}

		if curr == needle {
			break
		}

		adjs := wam.matrix[curr]
		for i := 0; i < len(adjs); i++ {
			if adjs[i] == 0 {
				continue
			}
			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			q.Enqueue(i)
		}
	}

	if prev[needle] == -1 {
		return []int{}, nil
	}

	curr := needle
	out := []int{}

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	out = append(out, source)
	slices.Reverse(out)
	return out, nil
}
