package leetcode

// https://leetcode.cn/problems/plus-one/description/

func plusOne(digits []int) []int {
	n := len(digits)

	ans := make([]int, n+1)
	t := digits[n-1] + 1
	extra := t / 10
	ans[n] = t % 10
	for i := n - 2; i >= 0; i-- {
		t = digits[i] + extra
		extra = t / 10
		ans[i+1] = t % 10
	}

	ans[0] = extra
	if ans[0] == 0 {
		return ans[1:]
	}

	return ans

}
