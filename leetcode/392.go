package leetcode

// https://leetcode.cn/problems/is-subsequence/description/?envType=study-plan-v2&envId=top-interview-150

func isSubsequence(s string, t string) bool {
	if len(s) <= 0 {
		return true
	}
	left := 0
	n := len(t)
	for i := 0; i < n; i++ {
		if t[i] == s[left] {
			left++
		}
		if left == len(s) {
			return true
		}
	}
	if left == len(s) {
		return true
	}
	return false

}
