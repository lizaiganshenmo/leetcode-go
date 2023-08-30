package leetcode

// 36. 有效的数独
// 请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）

// 注意：

// 一个有效的数独（部分已被填充）不一定是可解的。
// 只需要根据以上规则，验证已经填入的数字是否有效即可。
// 空白格用 '.' 表示。

func isValidSudoku(board [][]byte) bool {
	// 思路：哈希表 + 一次遍历
	rows := [10][10]bool{}
	cols := [10][10]bool{}
	boxes := [10][10]bool{}

	// 9*9 网格
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '0'
			// 判断该坐标在第几个box
			t := i/3*3 + j/3
			if rows[i][num] || cols[j][num] || boxes[t][num] {
				return false
			}

			rows[i][num] = true
			cols[j][num] = true
			boxes[t][num] = true
		}
	}

	return true
}
