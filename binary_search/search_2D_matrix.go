package binary_search

func SearchMatrix(matrix [][]int, target int) bool {
	for i := range matrix {
		if -1 != Search(matrix[i], target) {
			return true
		}
	}
	return false
}
