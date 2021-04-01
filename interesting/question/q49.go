package question

func longest(x, y int) int {
	dp := make([][]int, x)
	xCount := make(map[int]int)
	yCount := make(map[int]int)
	for i := 0; i < x; i++ {
		xCount[i] = 0
		dp[i][0] = 1
	}
	for i := 0; i < y; i++ {
		yCount[i] = 0
		dp[0][i] = 1
	}

	for i := 0; i < x; i++ {

	}
}
