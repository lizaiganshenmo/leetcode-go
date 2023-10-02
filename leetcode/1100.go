package leetcode

// https://leetcode.cn/problems/find-k-length-substrings-with-no-repeated-characters/

func numKLenSubstrNoRepeats(s string, k int) int {
	// 思路 : 滑动窗口
	if k == 1 {
		return len(s)
	}
	if k > len(s) {
		return 0
	}
	n := len(s)
	left, right := 0, 0
	charNums := make([]int, 26)
	// charNums[s[left]-'a'] = 1
	var ans int
	for right < n {
		charNums[s[right]-'a']++
		for charNums[s[right]-'a'] > 1 {
			charNums[s[left]-'a']--
			left++

		}
		if right < n && right-left+1 == k {
			ans++
			charNums[s[left]-'a']--
			left++

		}
		right++

	}

	return ans

}
