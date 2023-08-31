package leetcode

import "fmt"

// https://leetcode.cn/problems/game-of-life/?envType=study-plan-v2&envId=top-interview-150

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	willLive := make([][]bool, m)
	for i := range willLive {
		willLive[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if canLive(board, i, j) {
				willLive[i][j] = true
			}

		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if willLive[i][j] {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}

		}
	}

}

func canLive(board [][]int, i, j int) bool {
	// 判断 (i,j)位置的元素下阶段存活
	aroundLiveNum := 0 // 周围活细胞数目
	m, n := len(board), len(board[0])
	// 左上
	if i-1 >= 0 && j-1 >= 0 && board[i-1][j-1] == 1 {
		aroundLiveNum++
	}
	// 正上
	if i-1 >= 0 && board[i-1][j] == 1 {
		aroundLiveNum++
	}
	// 右上
	if i-1 >= 0 && j+1 < n && board[i-1][j+1] == 1 {
		aroundLiveNum++
	}
	// 右
	if j+1 < n && board[i][j+1] == 1 {
		aroundLiveNum++
	}
	// 右下
	if i+1 < m && j+1 < n && board[i+1][j+1] == 1 {
		aroundLiveNum++
	}
	// 下
	if i+1 < m && board[i+1][j] == 1 {
		aroundLiveNum++
	}
	// 左下
	if i+1 < m && j-1 >= 0 && board[i+1][j-1] == 1 {
		aroundLiveNum++
	}
	// 左
	if j-1 >= 0 && board[i][j-1] == 1 {
		aroundLiveNum++
	}

	// 存活规则
	// 	如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
	// 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
	// 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
	// 如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
	// 死细胞
	if i == 2 && j == 2 {
		fmt.Println("m  n : ", m, n)
		fmt.Println("aroundNum : ", aroundLiveNum)
		fmt.Println("aroundNum board ij: ", board[i][j])
	}
	if board[i][j] == 0 && aroundLiveNum == 3 {
		return true
	}
	if board[i][j] == 1 {
		// fmt.Println("aaaaaaaaaa", board[i][j])
		if aroundLiveNum < 2 {
			return false
		} else if aroundLiveNum == 2 || aroundLiveNum == 3 {
			// fmt.Println("bbbbbbbbbbbj: ", board[i][j])
			return true
		}

	}

	return false

}
