package hard

import "math"

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
			dp[i][j] = int(math.Min(float64(dp[i][j]), math.Max(float64(dp[i-1][right-1]), float64(dp[i][j-right]))+1.0))
		}
	}
	return dp[k][n]
}
