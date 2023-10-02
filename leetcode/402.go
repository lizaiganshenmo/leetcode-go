package leetcode

import "strings"

// https://leetcode.cn/problems/remove-k-digits/

func removeKdigits(num string, k int) string {
	// 思路: 贪心
	// 从左向右遍历, 遇到数字大于前面的,则后面数字删除， 前面数字大于后面的,前面删除
	//
	//
	// 从左至右扫描，当前扫描的数还不确定要不要删，入栈暂存。
	// 123531这样「高位递增」的数，肯定不会想删高位，会尽量删低位。
	// 432135这样「高位递减」的数，会想干掉高位，直接让高位变小，效果好。
	// 所以，如果当前遍历的数比栈顶大，符合递增，是满意的，让它入栈。
	// 如果当前遍历的数比栈顶小，栈顶立刻出栈，不管后面有没有更大的，
	n := len(num)
	stack := make([]byte, 0, n)
	for i := range num {
		tmp := num[i]
		for k > 0 && len(stack) > 0 && tmp < stack[len(stack)-1] {
			k--
			stack = stack[0 : len(stack)-1]

		}
		stack = append(stack, tmp)

	}

	stack = stack[0 : len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		return "0"
	}

	return ans

}
