package problem

func minCostClimbingStairs(cost []int) int {
	l:=len(cost)
	dp:=make([]int,l+1)
	dp[0]=0
	dp[1]=0
	for i := 2; i <=l ; i++ {
		dp[i]=min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])
	}
	return dp[l]
}
