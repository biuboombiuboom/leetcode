package problem

func maxProfit(prices []int, fee int) int {
	dp := make([][2]int, len(prices))
	//dp[0][0] 不持有股票时候的最大收益
	//dp[0][1] 持有股票时候的最大收益
	dp[0] =  [2]int{0, 0 - prices[0]}
	for i := 1; i < len(prices); i++ {
		//如果当天有卖出操作的话 则当天收益为dp[i-1][1]+prices[i]-fee
		//如果当天有买入的话 则当天收益为dp[i-1][0]-prices[i]
		dp[i][0] = max(dp[i-1][0], prices[i]+dp[i-1][1]-fee)
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}
