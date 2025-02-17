package graphs

import (
	"robtrizzo/data-and-algo-practice/internal/assert"
	"testing"
)

func TestBFS(t *testing.T) {
	var matrix [][]int
	matrix = append(matrix, []int{0, 1, 4, 5, 0})
	matrix = append(matrix, []int{1, 0, 0, 0, 0})
	matrix = append(matrix, []int{0, 0, 0, 2, 0})
	matrix = append(matrix, []int{0, 0, 0, 0, 5})
	matrix = append(matrix, []int{0, 0, 0, 0, 0})
	wam := WeightedAdjacencyMatrix{
		vertices: 5,
		matrix:   matrix,
	}
	path, err := wam.BFS(0, 4)
	assert.NilError(t, err)
	assert.ArraysEqual(t, path, []int{0, 3, 4})
	path, err = wam.BFS(0, 1)
	assert.NilError(t, err)
	assert.ArraysEqual(t, path, []int{0, 1})
	path, err = wam.BFS(1, 4)
	assert.NilError(t, err)
	assert.ArraysEqual(t, path, []int{1, 0, 3, 4})
}
