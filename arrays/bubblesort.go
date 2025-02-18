package arrays

func BubbleSort(arr *[]int) {
	l := len(*arr)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				Swap(arr, j, j+1)
			}
		}
	}
}
