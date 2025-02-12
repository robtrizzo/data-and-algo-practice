package arrays

// this naively assumes i and j are < len(*arr)
func swap(arr *[]int, i int, j int) {
	tmp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = tmp
}
