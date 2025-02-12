package arrays

import (
	"robtrizzo/data-and-algo-practice/internal/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	actual := []int{2, 1, 4, 5, 3}
	BubbleSort(&actual)
	assert.ArraysEqual[int](t, actual, expected)
}
