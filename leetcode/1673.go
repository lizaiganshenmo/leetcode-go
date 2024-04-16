package leetcode

// https://leetcode.cn/problems/find-the-most-competitive-subsequence/description/
func mostCompetitive(nums []int, k int) []int {
	// 思路 : 单调栈
	n := len(nums)
	if n == k {
		return nums
	}

	ans := make([]int, 0, k)
	delNum := n - k
	for _, v := range nums {
		for delNum > 0 && len(ans) > 0 && ans[len(ans)-1] > v {
			delNum--
			ans = ans[:len(ans)-1]
		}
		ans = append(ans, v)
	}

	for delNum > 0 {
		ans = ans[:len(ans)-1]
		delNum--
	}

	return ans

}
