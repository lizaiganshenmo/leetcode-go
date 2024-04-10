package leetcode

// https://leetcode.cn/problems/maximum-absolute-sum-of-any-subarray/description/
func maxAbsoluteSum(nums []int) (ans int) {
	var fMax, fMin int
	for _, x := range nums {
		fMax = max1749(fMax, 0) + x
		fMin = min1749(fMin, 0) + x
		ans = max1749(ans, max1749(fMax, -fMin))
	}
	return
}

func min1749(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max1749(a, b int) int {
	if b > a {
		return b
	}
	return a
}
