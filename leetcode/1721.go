package leetcode

// https://leetcode.cn/problems/volume-of-histogram-lcci/

func trap1(height []int) int {
	// 动态规划
	// 思路 ： 遍历两遍数组，最大左边界和最大右边界
	n := len(height)
	if n == 0 {
		return 0
	}
	leftBounds := make([]int, n)

	for i := range height {
		if i == 0 {
			leftBounds[i] = height[0]
			continue
		}
		leftBounds[i] = max99(leftBounds[i-1], height[i])
	}

	rightBounds := make([]int, n)
	rightBounds[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightBounds[i] = max99(rightBounds[i+1], height[i])
	}

	var ans int
	for i := range height {
		ans += min99(leftBounds[i], rightBounds[i]) - height[i]
	}

	return ans

}

func max99(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min99(a, b int) int {
	if a < b {
		return a
	}
	return b
}
