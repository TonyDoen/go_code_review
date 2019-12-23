package hard

func SuperEggDrop1(k, n int) int {
	dp := make([][]int, k+1)
	for j := 0; j <= n; j++ {
		dp[1] = make([]int, n+1)
		dp[1][j] = j
	}
	for i := 2; i <= k; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = j
			left := 1
			right := j
			for left < right {
				mid := left + (right-left)/2
				if dp[i-1][mid-1] < dp[i][j-mid] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			dp[i][j] = min(dp[i][j], max(dp[i-1][right-1], dp[i][j-right])+1)
		}
	}
	return dp[k][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
