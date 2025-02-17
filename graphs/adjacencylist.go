package graphs

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
