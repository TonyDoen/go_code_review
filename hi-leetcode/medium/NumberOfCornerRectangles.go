package medium

/**
 * Given a grid where each entry is only 0 or 1, find the number of corner rectangles.
 * A corner rectangle is 4 distinct 1s on the grid that form an axis-aligned rectangle.
 * Note that only the corners need to have the value 1. Also, all four 1s used must be distinct.
 *
 * Example 1:
 * Input: grid =
 * [[1, 0, 0, 1, 0],
 *  [0, 0, 1, 0, 1],
 *  [0, 0, 0, 1, 0],
 *  [1, 0, 1, 0, 1]]
 * Output: 1
 * Explanation: There is only one corner rectangle, with corners grid[1][2], grid[1][4], grid[3][2], grid[3][4].
 *
 * Example 2:
 * Input: grid =
 * [[1, 1, 1],
 *  [1, 1, 1],
 *  [1, 1, 1]]
 * Output: 9
 * Explanation: There are four 2x2 rectangles, four 2x3 and 3x2 rectangles, and one 3x3 rectangle.
 *
 * Example 3:
 * Input: grid =
 * [[1, 1, 1, 1]]
 * Output: 0
 * Explanation: Rectangles must have four distinct corners.
 *
 * Note:
 * The number of rows and columns of grid will each be in the range [1, 200].
 * Each grid[i][j] will be either 0 or 1.
 * The number of 1s in the grid will be at most 6000.
 */
func CountCornerRectangles(grid [][]int) int {
	if nil == grid {
		return -1
	}

	result, row, col := 0, len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := i + 1; j < row; j++ {
			cnt := 0
			for k := 0; k < col; k++ {
				if 1 == grid[i][k] && 1 == grid[j][k] {
					cnt++
				}
			}
			result += cnt * (cnt - 1) / 2
		}
	}
	return result
}

func CountCornerRectangles1(grid [][]int) int {
	if nil == grid {
		return -1
	}
	result, row, col := 0, len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if 0 == grid[i][j] {
				continue
			}
			for h := 1; h < row-i; h++ {
				if 0 == grid[i+h][j] {
					continue
				}
				for w := 1; w < col-j; w++ {
					if 1 == grid[i][j+w] && 1 == grid[i+h][j+w] {
						result++
					}
				}
			}
		}
	}
	return result
}
