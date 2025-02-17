package graphs

import (
	"math"
	"slices"
)

type WeightedAdjacencyList struct {
	edges [][]Edge
}

type Edge struct {
	to int
	w  int
}

func (wal *WeightedAdjacencyList) walk(curr int, needle int, seen *[]bool, path *[]int) bool {
	if (*seen)[curr] {
		return false
	}
	(*seen)[curr] = true

	*path = append(*path, curr)
	if curr == needle {
		return true
	}

	list := wal.edges[curr]
	for i := 0; i < len(list); i++ {
		edge := list[i]
		if wal.walk(edge.to, needle, seen, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1] // pop

	return false
}

func (wal *WeightedAdjacencyList) DFS(source int, needle int) []int {
	seen := initSeen(len(wal.edges))
	var path []int

	wal.walk(source, needle, &seen, &path)

	return path
}

func hasUnvisited(seen []bool, dists []int) bool {
	for i := 0; i < len(seen); i++ {
		if !seen[i] && dists[i] < math.MaxInt32 {
			return true
		}
	}
	return false
}

func getLowestUnvisited(seen []bool, dists []int) int {
	idx := -1
	lowestDistance := math.MaxInt32

	for i := 0; i < len(seen); i++ {
		if seen[i] {
			continue
		}

		if lowestDistance > dists[i] {
			lowestDistance = dists[i]
			idx = i
		}
	}

	return idx
}

func (wal *WeightedAdjacencyList) DSP(source int, sink int) []int {
	seen := initSeen(len(wal.edges))
	prev := initPrev(len(wal.edges))
	dists := initDists(len(wal.edges))
	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		adjs := wal.edges[curr]
		for i := 0; i < len(adjs); i++ {
			edge := adjs[i]
			if seen[edge.to] {
				continue
			}

			dist := dists[curr] + edge.w
			if dist < dists[edge.to] {
				dists[edge.to] = dist
				prev[edge.to] = curr
			}
		}
	}

	var out []int
	curr := sink

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	out = append(out, source)
	slices.Reverse(out)
	return out
}
