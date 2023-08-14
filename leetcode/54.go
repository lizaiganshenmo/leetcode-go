package leetcode

// https://leetcode.cn/problems/spiral-matrix/?envType=study-plan-v2&envId=huawei-2023-spring-sprint

// 54. 螺旋矩阵
// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

// 示例 1：

// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]
// 示例 2：

// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]

// 提示：

// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100

const (
	toRight = iota
	toDown
	toLeft
	toUp
)

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	num := m * n
	if num == 1 {
		return []int{matrix[0][0]}
	}
	// 边界
	leftBound, upBound, rightBound, downBound := 0, 0, n-1, m-1

	i, j := 0, 0
	ans := make([]int, 0, num)
	ans = append(ans, matrix[i][j])
	num--
	// 初始化怎么走
	var flag int
	if m > 1 && n == 1 {
		flag = toDown
	}
	for num > 0 {
		// 向右走
		switch flag {
		case toRight:
			// 向右走 j自增直到边界rightBound
			for j < rightBound {
				j++
				ans = append(ans, matrix[i][j])
				num--
			}
			flag = toDown
			upBound++
		case toDown:
			// 向下走 i自增到边界downBound
			for i < downBound {
				i++
				ans = append(ans, matrix[i][j])
				num--

			}
			flag = toLeft
			rightBound--
		case toLeft:
			// 向左走 j自减到边界leftBound
			for j > leftBound {
				j--
				ans = append(ans, matrix[i][j])
				num--
			}
			flag = toUp
			downBound--
		case toUp:
			// 向上走 i自减到边界upBound
			for i > upBound {
				i--
				ans = append(ans, matrix[i][j])
				num--
			}
			flag = toRight
			leftBound++

		}

	}

	return ans

}
