package hard

/**
 * Cat and Mouse
 * A game on an undirected graph is played by two players, Mouse and Cat, who alternate turns.
 * The graph is given as follows: graph[a] is a list of all nodes b such that ab is an edge of the graph.
 * Mouse starts at node 1 and goes first, Cat starts at node 2 and goes second, and there is a Hole at node 0.
 * During each player's turn, they must travel along one edge of the graph that meets where they are.  For example, if the Mouse is at node 1, it must travel to any node in graph[1].
 * Additionally, it is not allowed for the Cat to travel to the Hole (node 0.)
 *
 * Then, the game can end in 3 ways:
 * If ever the Cat occupies the same node as the Mouse, the Cat wins.
 * If ever the Mouse reaches the Hole, the Mouse wins.
 * If ever a position is repeated (ie. the players are in the same position as a previous turn, and it is the same player's turn to move), the game is a draw.
 *
 * Given a graph, and assuming both players play optimally, return 1 if the game is won by Mouse, 2 if the game is won by Cat, and 0 if the game is a draw.
 *
 * Example 1:
 * Input: [[2,5],[3],[0,4,5],[1,4,5],[2,3],[0,2,3]]
 * Output: 0
 * Explanation:
 * 4---3---1
 * |   |
 * 2---5
 *  \ /
 *   0
 *
 * Note:
 * 3 <= graph.length <= 50
 * It is guaranteed that graph[1] is non-empty.
 * It is guaranteed that graph[2] contains a non-zero element.
 */

/**
 * 题意：猫和老鼠
 * 思路：这道题是猫抓老鼠的问题，Tom and Jerry 都看过吧。这道题在无向图上模仿了猫抓老鼠的这一个过程，老鼠位于结点1，猫位于结点2，老鼠的目标是逃回老鼠洞结点0，猫的目标是在老鼠进洞之前抓住它。这里假设猫和老鼠都会选择最优的策略。若老鼠能成功逃回洞里，则返回1；若猫能成功抓到老鼠，则返回2；若谁也不能达到目标，则表示平局，返回0。其实这道题的本质还是一个无向图的遍历问题，只不过现在有两个物体在遍历，比一般的图遍历要复杂一些。假设图中有n个结点，不论是猫还是老鼠，当各自走完了n个结点时还没有分出胜负，则表示平局，若一人走一步，则最多有 2n 步。这样的话每一个状态实际上是由三个因素组成的：当前步数，老鼠所在结点，和猫所在结点。这里可以用动态规划 Dynamic Programming 来解，使用一个三维的 dp 数组，其中 dp[t][x][y] 表示当前步数为t，老鼠在结点x，猫在结点y时最终会返回的值，均初始化为 -1。要求的其实是起始状态 dp[0][1][2] 的返回值，但没法一下子求出，这个起始状态实际上是要通过其他状态转移过来，就比如说是求二叉树最大深度的递归函数，虽然对根结点调用递归函数的返回值就是最大深度，但在函数遇到叶结点之前都无法得知深度。先来看一些终止状态，首先当老鼠到达洞口的时候，此时老鼠赢，返回值是1，即所有 dp[?][0][?] 状态的返回值都是1。其次，当猫和老鼠处于同一个位置时，表示猫抓到老鼠了，此时猫赢，返回值是2，即所有 dp[?][y][y] 状态的返回值都是2。最后，当走完了 2n 步还没有分出胜负的话，则是平局，直接返回0即可。
 * 理清了上面的思路，其实代码就不难写了，这里使用递归的写法，在递归函中，首先判断步数是否到了 2n，是的话直接返回0；否则判断x和y是否相等，是的话当前状态赋值为2并返回；否则再判断x是否等于0，是的话当前状态赋值为1并返回。若当前状态的 dp 值不是 -1，则表示之前已经更新过了，不需要重复计算了，直接返回即可。否则就要来计算当前的 dp 值，先确定当前该谁走，只要判断t的奇偶即可，因为最开始步数0的时候是老鼠先走。若此时该老鼠走了，它能走的相邻结点可以在 graph 中找到，对于每一个可以到达的相邻结点，都调用递归函数，此时步数是 t+1，老鼠位置为相邻结点，猫的位置不变。若返回值是1，表示老鼠赢，则将当前状态赋值为1并返回；若返回状态是2，此时不能立马返回猫赢，因为老鼠可以不走这个结点；若返回值是0，表示老鼠走这个结点是有平局的机会，但老鼠还是要争取赢的机会，所以此时用一个 bool 变量标记下猫肯定赢不了，但此时也不能直接返回，因为 Jerry 一直要追寻赢的机会。直到遍历完了所有可能性，老鼠最终还是没有赢，则看下之前那个 bool 型变量 catWin，若为 true，则标记当前状态为2并返回，反之，则标记当前状态为0并返回。若此时该猫走了，基本跟老鼠的策略相同，它能走的相邻结点也可以在 graph 中找到，对于每一个可以到达的相邻结点，首先要判断是否为结点0（老鼠洞），因为猫是不能进洞的，所以要直接跳过这个结点。否则就调用递归函数，此时步数是 t+1，老鼠位置不变，猫的位置为相邻结点。若返回值是2，表示猫赢，则将当前状态赋值为2并返回；若返回状态是1，此时不能立马返回老鼠赢，因为猫可以不走这个结点；若返回值是0，表示猫走这个结点是有平局的机会，但猫还是要争取赢的机会，所以此时用一个 bool 变量标记下老鼠肯定赢不了，但此时也不能直接返回，因为 Tom 也一直要追寻赢的机会。直到遍历完了所有可能性，猫最终还是没有赢，则看下之前那个 bool 型变量 mouseWin，若为 true，则标记当前状态为1并返回，反之，则标记当前状态为0并返回
 */
func CatAndMouseGame1(graph [][]int) int {
	dp := d3(len(graph))
	res := catAndMouseGame1(dp, graph, 0, 1, 2)
	return res
}

func d3(n int) [][][]int {
	a1 := make([]int, n)
	for z := 0; z < n; z++ {
		a1[z] = -1
	}
	a2 := make([][]int, 0, n)
	for z := 0; z < n; z++ {
		a2 = append(a2, a1)
	}
	a3 := make([][][]int, 0, 2*n)
	for z := 0; z < 2*n; z++ {
		a3 = append(a3, a2)
	}
	return a3
}

func catAndMouseGame1(dp [][][]int, graph [][]int, t, x, y int) int {
	if t == len(graph)*2 {
		return 0
	}
	if x == y {
		dp[t][x][y] = 2
		return 2
	}
	if x == 0 {
		dp[t][x][y] = 1
		return 1
	}
	if dp[t][x][y] != -1 {
		return dp[t][x][y]
	}

	mouseTurn := t%2 == 0
	if mouseTurn {
		catWin := true
		for i := 0; i < len(graph[x]); i++ {
			next := catAndMouseGame1(dp, graph, t+1, graph[x][i], y)
			if next == 1 {
				dp[t][x][y] = 1
				return 1
			} else if next != 2 {
				catWin = false
			}
		}
		if catWin {
			dp[t][x][y] = 2
			return 2
		} else {
			dp[t][x][y] = 0
			return 0
		}
	} else {
		mouseWin := true
		for i := 0; i < len(graph[y]); i++ {
			if graph[y][i] == 0 {
				continue
			}
			next := catAndMouseGame1(dp, graph, t+1, x, graph[y][i])
			if next == 2 {
				dp[t][x][y] = 2
				return 2
			} else if next != 1 {
				mouseWin = false
			}
		}
		if mouseWin {
			dp[t][x][y] = 1
			return 1
		} else {
			dp[t][x][y] = 0
			return 0
		}
	}
}
