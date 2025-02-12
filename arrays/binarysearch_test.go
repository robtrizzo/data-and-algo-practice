package arrays

import (
	"robtrizzo/data-and-algo-practice/internal/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5}
	assert.Equal[bool](t, BinarySearch(haystack, 2), true)
	assert.Equal[bool](t, BinarySearch(haystack, 8), false)
}
