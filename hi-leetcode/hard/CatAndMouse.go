package hard

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
	a2 := make([][]int, 0)
	for z := 0; z < n; z++ {
		a2 = append(a2, a1)
	}
	a3 := make([][][]int, 0)
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
