package graphs

func initSeen(len int) []bool {
	seen := make([]bool, len)
	for i := range seen {
		seen[i] = false
	}
	return seen
}

func initPrev(len int) []int {
	prev := make([]int, len)
	for i := range prev {
		prev[i] = -1
	}
	return prev
}
