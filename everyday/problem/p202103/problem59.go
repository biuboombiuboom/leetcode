package p202103

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	num:=1
	left, right, top, bottom := 0, n-1, 0, n-1
	for left<=right&&top>=bottom{
		for column := left; column <= right; column++ {
			ans[top][column]=num
			num++
		}
		for row := top + 1; row <= bottom; row++ {
			ans[row][right] = num
			num++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				ans[bottom][column] = num
				num++
			}
			for row := bottom; row > top; row-- {
				ans[row][left] = num
				num++
			}
		}
		left++
		right--
		top++
		bottom--


	}
	return ans
}
