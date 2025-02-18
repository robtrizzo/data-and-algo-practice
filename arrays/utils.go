package arrays

// this naively assumes i and j are < len(*arr)
func Swap[T any](arr *[]T, i int, j int) {
	tmp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = tmp
}

func Parent(idx int) int {
	return (idx - 1) / 2
}

func LeftChild(idx int) int {
	return idx*2 + 1
}

func RightChild(idx int) int {
	return idx*2 + 2
}
