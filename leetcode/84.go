package leetcode

// https://leetcode.cn/problems/largest-rectangle-in-histogram/description/
func largestRectangleArea(heights []int) int {
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
		ans = max84(ans, (rBounds[i]-lBounds[i]-1)*heights[i])
	}
	return ans

}

func max84(x, y int) int {
	if x > y {
		return x
	}
	return y
}
