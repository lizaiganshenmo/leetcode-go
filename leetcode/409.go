package leetcode

// https://leetcode.cn/problems/longest-palindrome/description/

func longestPalindrome12(s string) int {
	// 思路 : map存储字符数目
	n := len(s)
	if n <= 1 {
		return n
	}

	mp := make(map[byte]int)
	for i := range s {
		mp[s[i]]++
	}

	var canOdd bool
	var ans int
	for _, v := range mp {
		t := v / 2
		ans += 2 * t
		if v%2 != 0 {
			canOdd = true
		}
	}

	if canOdd {
		ans++
	}

	return ans

}
