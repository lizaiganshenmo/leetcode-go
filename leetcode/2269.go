package leetcode

import "strconv"

// https://leetcode.cn/problems/find-the-k-beauty-of-a-number/description/
func divisorSubstrings(num int, k int) int {
	// 思路: 滑动窗口
	numStr := strconv.Itoa(num)
	n := len(numStr)
	if n < k {
		return 0
	}

	left, right := 0, 0
	var ans int
	for ; right < n; right++ {
		if right-left+1 == k {
			tmp, _ := strconv.Atoi(numStr[left : right+1])
			if tmp != 0 && num%tmp == 0 {
				ans++
			}
			left++
		}
	}

	return ans

}
