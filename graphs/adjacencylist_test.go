package graphs

import (
	"robtrizzo/data-and-algo-practice/internal/assert"
	"testing"
)

func TestDFS(t *testing.T) {
	var edges [][]Edge
	edges = append(edges, []Edge{{to: 1, w: 1}, {to: 2, w: 4}, {to: 3, w: 5}})
	edges = append(edges, []Edge{{to: 0, w: 1}})
	edges = append(edges, []Edge{{to: 3, w: 2}})
	edges = append(edges, []Edge{{to: 4, w: 5}})
	edges = append(edges, []Edge{})
	wal := WeightedAdjacencyList{
		edges: edges,
	}
	path := wal.DFS(0, 4)
	assert.ArraysEqual(t, path, []int{0, 2, 3, 4})
	path = wal.DFS(0, 1)
	assert.ArraysEqual(t, path, []int{0, 1})
	path = wal.DFS(1, 4)
	assert.ArraysEqual(t, path, []int{1, 0, 2, 3, 4})
}
