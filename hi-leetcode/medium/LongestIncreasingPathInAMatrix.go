package medium

/**
 * Longest Increasing Path in a Matrix
 *
 * Given an integer matrix, find the length of the longest increasing path.
 * From each cell, you can either move to four directions: left, right, up or down. You may NOT move diagonally or move outside of the boundary (i.e. wrap-around is not allowed).
 *
 * Example 1:
 * nums = [
 *   [9,9,4],
 *   [6,6,8],
 *   [2,1,1]
 * ]
 * Return 4
 * The longest increasing path is [1, 2, 6, 9].
 *
 * Example 2:
 * nums = [
 *   [3,4,5],
 *   [3,2,6],
 *   [2,2,1]
 * ]
 * Return 4
 * The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.
 *
 *
 * 题意： 矩阵中的最长递增路径
 * 给我们一个二维数组，让我们求矩阵中最长的递增路径，规定我们只能上下左右行走，不能走斜线或者是超过了边界。
 *
 * 思路：DFS + dp
 * 1、dp[i][j]表示数组中以(i,j)为起点的最长递增路径的长度，初始将dp数组都赋为0，
 * 2、递归调用时，遇到某个位置(x, y), 如果dp[x][y]不为0的话，我们直接返回dp[x][y]即可，不需要重复计算。
 * 3、以数组中每个位置都为起点调用递归来做，比较找出最大值。在以一个位置为起点用DFS搜索时，对其四个相邻位置进行判断，如果相邻位置的值大于上一个位置，则对相邻位置继续调用递归，并更新一个最大值，搜素完成后返回即可
 *
 *
 * 这道题的解法要用递归和DP来解，用DP的原因是为了提高效率，避免重复运算。
 * 维护一个二维动态数组dp，其中dp[i][j]表示数组中以(i,j)为起点的最长递增路径的长度，初始将dp数组都赋为0，
 * 当我们用递归调用时，遇到某个位置(x, y), 如果dp[x][y]不为0的话，我们直接返回dp[x][y]即可，不需要重复计算。
 *
 * 我们需要以数组中每个位置都为起点调用递归来做，比较找出最大值。
 * 在以一个位置为起点用DFS搜索时，对其四个相邻位置进行判断，如果相邻位置的值大于上一个位置，则对相邻位置继续调用递归，并更新一个最大值，搜素完成后返回即可，
 */
func LongestIncreasingPath0(matrix [][]int) int {
	if nil == matrix {
		return -1
	}
	result, row, col := 1, len(matrix), len(matrix[0])

	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			tmp := helpLongestIncreasingPath0(i, j, dp, matrix)
			if tmp > result {
				result = tmp
			}
		}
	}
	return result
}

var dirs = [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} // 上/左/下/右
func helpLongestIncreasingPath0(i, j int, dp, matrix [][]int) int {
	if dp[i][j] > 0 {
		return dp[i][j]
	}
	result, row, col := 1, len(matrix), len(matrix[0])
	for _, dir := range dirs {
		x, y := i+dir[0], j+dir[1]
		if x < 0 || x >= row || y < 0 || y >= col || matrix[x][y] <= matrix[i][j] {
			continue
		}
		cnt := 1 + helpLongestIncreasingPath0(x, y, dp, matrix)
		if cnt > result {
			result = cnt
		}
	}
	dp[i][j] = result
	return result
}
