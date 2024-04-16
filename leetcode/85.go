package leetcode

// https://leetcode.cn/problems/maximal-rectangle/description/
func maximalRectangle(matrix [][]byte) int {
	// 思路: 单调栈
	// 参考 的方法二https://leetcode.cn/problems/maximal-rectangle/solutions/9535/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-1-8/
	m, n := len(matrix), len(matrix[0])
	heights := make([]int, n)
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}

		}
		ans = max85(ans, largestRectangleArea1(heights))

	}

	return ans

}

func max85(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func largestRectangleArea1(heights []int) int {
	// 思路 : 单调栈
	// 要求出最大矩形，则要直到每个i下标对应的height[i] ,其左右边界如何找到，保证左右边界范围内的数，都大于等于height[i] 使用单调栈实现
	n := len(heights)
	if n == 0 {
		return 0
	}

	lBounds, rBounds := make([]int, n), make([]int, n)

	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			lBounds[i] = -1
		} else {
			lBounds[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			rBounds[i] = n
		} else {
			rBounds[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = max85(ans, (rBounds[i]-lBounds[i]-1)*heights[i])
	}
	return ans

}
