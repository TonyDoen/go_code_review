package medium

/**
 * Alice plays the following game, loosely based on the card game "21".
 * Alice starts with 0 points, and draws numbers while she has less than K points.  During each draw, she gains an integer number of points randomly from the range [1, W], where W is an integer.  Each draw is independent and the outcomes have equal probabilities.
 * Alice stops drawing numbers when she gets K or more points.  What is the probability that she has N or less points?
 *
 * Example 1:
 * Input: N = 10, K = 1, W = 10
 * Output: 1.00000
 *
 * Explanation:  Alice gets a single card, then stops.
 *
 * Example 2:
 * Input: N = 6, K = 1, W = 10
 * Output: 0.60000
 *
 * Explanation:  Alice gets a single card, then stops.
 * In 6 out of W = 10 possibilities, she is at or below N = 6 points.
 *
 * Example 3:
 * Input: N = 21, K = 17, W = 10
 * Output: 0.73278
 *
 * Note:
 * 1. 0 <= K <= N <= 10000
 * 2. 1 <= W <= 10000
 * 3. Answers will be accepted as correct if they are within 10^-5 of the correct answer.
 * 4. The judging time limit has been reduced for this question.
 *
 *
 *  新二十一点游戏
 *
 */ /**
 * Alice plays the following game, loosely based on the card game "21".
 * Alice starts with 0 points, and draws numbers while she has less than K points.  During each draw, she gains an integer number of points randomly from the range [1, W], where W is an integer.  Each draw is independent and the outcomes have equal probabilities.
 * Alice stops drawing numbers when she gets K or more points.  What is the probability that she has N or less points?
 *
 * Example 1:
 * Input: N = 10, K = 1, W = 10
 * Output: 1.00000
 *
 * Explanation:  Alice gets a single card, then stops.
 *
 * Example 2:
 * Input: N = 6, K = 1, W = 10
 * Output: 0.60000
 *
 * Explanation:  Alice gets a single card, then stops.
 * In 6 out of W = 10 possibilities, she is at or below N = 6 points.
 *
 * Example 3:
 * Input: N = 21, K = 17, W = 10
 * Output: 0.73278
 *
 * Note:
 * 1. 0 <= K <= N <= 10000
 * 2. 1 <= W <= 10000
 * 3. Answers will be accepted as correct if they are within 10^-5 of the correct answer.
 * 4. The judging time limit has been reduced for this question.
 *
 *
 *  新二十一点游戏
 *
 */
func New21Game0(n, k, w int) float64 {
	if 0 == k || n >= k+w {
		return 1.0
	}
	sum := make([]float64, k+w)
	sum[0] = 1.0
	for i := 1; i < k+w; i++ {
		t := i - 1
		if t > k-1 {
			t = k - 1
		}
		if i <= w {
			sum[i] = sum[i-1] + sum[t]/float64(w)
		} else {
			sum[i] = sum[i-1] + (sum[t]-sum[i-w-1])/float64(w)
		}
	}
	return (sum[n] - sum[k-1]) / (sum[k+w-1] - sum[k-1])
}

func New21Game1(n, k, w int) float64 {
	if 0 == k || n >= k+w {
		return 1.0
	}
	dp := make([]float64, k+w)
	dp[0] = 1.0
	for i := 1; i < k+w; i++ {
		dp[i] = dp[i-1]
		if i <= w {
			dp[i] += dp[i-1] / float64(w)
		} else {
			dp[i] += (dp[i-1] - dp[i-w-1]) / float64(w)
		}
		if i > k {
			dp[i] -= (dp[i-1] - dp[k-1]) / float64(w)
		}
	}
	return dp[n] - dp[k-1]
}
