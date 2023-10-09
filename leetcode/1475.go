package leetcode

// https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/

func finalPrices(prices []int) []int {
	// 思路 : 单调栈
	n := len(prices)
	stack := make([]int, 0, n)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && prices[i] < prices[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = prices[i] - prices[stack[len(stack)-1]]
		} else {
			ans[i] = prices[i]
		}
		stack = append(stack, i)
	}

	return ans

}
