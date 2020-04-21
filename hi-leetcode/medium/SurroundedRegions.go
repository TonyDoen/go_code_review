package medium

/**
 * Surrounded Regions
 *
 * Given a 2D board containing'X'and'O', capture all regions surrounded by'X'.
 * A region is captured by flipping all'O's into'X's in that surrounded region .
 * For example,
 *     X X X X
 *     X O O X
 *     X X O X
 *     X O X X
 *
 * After running your function, the board should be:
 *     X X X X
 *     X X X X
 *     X X X X
 *     X O X X
 *
 *
 * 题意：所有与四条边相连的O都保留，其他O都变为X
 * 1、被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。
 * 2、任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。
 * 3、如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
 *
 * 思路：
 * 1、遍历四条边上的O，并深度遍历与其相连的O，将这些O都转为*
 * 2、将剩余的O变为X
 * 3、将剩余的*变为O
 */
const (
	O    = 'O'
	X    = 'X'
	STAR = '*'
)

func SurroundedRegions(matrix [][]rune) {
	if nil == matrix {
		return
	}
	row, col := len(matrix), len(matrix[0])
	if row < 1 || col < 1 {
		return
	}

	// 1、四条边上开始深度遍历，边上O及与边上O相连的全变为*
	for i := 0; i < row; i++ { // 行
		helpSurroundedRegions(i, 0, matrix)     // 第一列元素开始深度遍历
		helpSurroundedRegions(i, col-1, matrix) // 最后一列元素开始深度遍历
	}
	for j := 0; j < col; j++ {
		helpSurroundedRegions(0, j, matrix)
		helpSurroundedRegions(row-1, j, matrix)
	}

	// 2、将剩余的O变为X
	// 3、将剩余的*变为O
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			cij := matrix[i][j]
			if O == cij { // 将剩余的O变为X
				matrix[i][j] = X
			} else if STAR == cij { // 将剩余的*变为O
				matrix[i][j] = O
			}
		}
	}
}

//var dirs = [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
func helpSurroundedRegions(row, col int, matrix [][]rune) {
	// 非法判断
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
		return
	}
	if O == matrix[row][col] {
		matrix[row][col] = STAR
		for _, dir := range dirs {
			helpSurroundedRegions(row+dir[0], col+dir[1], matrix)
		}
	}
}
