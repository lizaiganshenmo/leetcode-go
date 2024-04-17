package leetcode

// https://leetcode.cn/problems/find-the-longest-semi-repetitive-substring/description/
func longestSemiRepetitiveSubstring(s string) int {
	// 思路: 滑动窗口
	var (
		dupByteIdx  = -1
		n           = len(s)
		left, right = 0, 1
		ans         = 0
	)
	if n <= 1 {
		return n
	}

	for ; right < n; right++ {
		if s[right] == s[right-1] {
			if dupByteIdx >= 0 {
				left = dupByteIdx
			}
			dupByteIdx = right

		}
		ans = max2730(ans, right-left+1)

	}
	return ans

}

func max2730(a, b int) int {
	if a > b {
		return a
	}
	return b
}
