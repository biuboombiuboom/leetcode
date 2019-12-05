package algorithm

func ClimbStairs(n int) int {
	dp := make([]int, 2)
	count := 0
	for i := 0; i < n; i++ {
		if i > 1 {
			count = dp[0] + dp[1]
			dp[0] = dp[1]
			dp[1] = count
		} else {
			count = i + 1
			dp[i] = count
		}
	}
	return count
}
