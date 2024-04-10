package leetcode

// https://leetcode.cn/problems/wtcaE1/description/?company_slug=bytedance
func lengthOfLongestSubstring1(s string) int {
	// 思路 : 滑动窗口 + map
	n := len(s)
	if n <= 1 {
		return n
	}

	left, right := 0, 1
	mp := make(map[byte]int) // key:val -> byte: idx
	mp[s[left]] = 0
	var ans int
	for right < n {
		if idx, ok := mp[s[right]]; ok {
			left = max16(left, idx+1)
		}
		mp[s[right]] = right

		ans = max16(ans, right-left+1)
		right++

	}

	return ans

}

func max16(a, b int) int {
	if a > b {
		return a
	}
	return b
}
