package graphs

import (
	"robtrizzo/data-and-algo-practice/internal/assert"
	"testing"
)

func TestDFS(t *testing.T) {
	var edges [][]Edge
	edges = append(edges, []Edge{{To: 1, W: 1}, {To: 2, W: 4}, {To: 3, W: 5}})
	edges = append(edges, []Edge{{To: 0, W: 1}})
	edges = append(edges, []Edge{{To: 3, W: 2}})
	edges = append(edges, []Edge{{To: 4, W: 5}})
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

func TestDSP(t *testing.T) {
	var edges [][]Edge
	edges = append(edges, []Edge{{To: 1, W: 1}, {To: 2, W: 5}})
	edges = append(edges, []Edge{{To: 2, W: 6}, {To: 3, W: 7}})
	edges = append(edges, []Edge{{To: 4, W: 1}})
	edges = append(edges, []Edge{{To: 2, W: 1}})
	edges = append(edges, []Edge{})
	wal := WeightedAdjacencyList{
		edges: edges,
	}
	path := wal.DSP(0, 4)
	assert.ArraysEqual(t, path, []int{0, 2, 4})
}

func TestDSPHeap(t *testing.T) {
	var edges [][]Edge
	edges = append(edges, []Edge{{To: 1, W: 1}, {To: 2, W: 5}})
	edges = append(edges, []Edge{{To: 2, W: 6}, {To: 3, W: 7}})
	edges = append(edges, []Edge{{To: 4, W: 1}})
	edges = append(edges, []Edge{{To: 2, W: 1}})
	edges = append(edges, []Edge{})
	wal := WeightedAdjacencyList{
		edges: edges,
	}
	path, err := wal.DSP_Heap(0, 4)
	assert.NilError(t, err)
	assert.ArraysEqual(t, path, []int{0, 2, 4})
}
