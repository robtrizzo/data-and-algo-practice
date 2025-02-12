package arrays

import (
	"robtrizzo/data-and-algo-practice/internal/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	actual := []int{2, 1, 4, 5, 3}
	QuickSort(&actual)
	assert.ArraysEqual[int](t, actual, expected)
}
