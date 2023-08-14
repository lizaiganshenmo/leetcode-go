package leetcode

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

// 3. 无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

// 示例 1:

// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:

// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:

// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

func lengthOfLongestSubstring(s string) int {
	// 思路 : 滑动窗口 + map
	// n := len(s)
	m := make(map[byte]int) // 字符出现的下标
	left := 0
	ans := 0

	for i := range s {

		if idx, ok := m[s[i]]; ok {
			left = max(idx+1, left)

		}

		m[s[i]] = i
		ans = max(ans, i-left+1)

	}
	return ans

}

func max1(a, b int) int {
	if a < b {
		return b
	}
	return a
}
