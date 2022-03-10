package p202103

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	row := 0
	for ; row < m; row++ {
		if matrix[row][0] == target {
			return true
		}
		if matrix[row][0] > target {
			break
		}
	}
	row--
	if row < 0 {
		return false
	}
	for i := 0; i < len(matrix[row]); i++ {
		if matrix[row][i] == target {
			return true
		}
	}
	return false
}
