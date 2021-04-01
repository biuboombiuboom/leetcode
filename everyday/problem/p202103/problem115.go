package p202103

func numDistinct(s string, t string) int {
	m := len(s)
	n := len(t)
	dp := make([][]int, m+1)
	if m<n{
		return 0
	}
	for i, _ := range dp {
		dp[i]=make([]int,n+1)
		dp[i][n]=1
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i]==t[j]{
				dp[i][j]=dp[i+1][j+1]+dp[i+1][j]
			}else{
				dp[i][j]=dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}
