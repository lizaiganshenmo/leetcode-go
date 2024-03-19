package leetcode

import (
	"strconv"
)

// https://leetcode.cn/problems/maximum-swap/description/

func maximumSwap(num int) int {
	// 思路 : 贪心
	// 一个栈从后向前遍历, 记录当前最大元素下标
	if num == 0 {
		return 0

	}
	ans := num
	numStr := strconv.FormatInt(int64(num), 10)
	numByte := []byte(numStr)
	n := len(numStr)
	stack := make([]int, 0, n)
	for idx := n - 1; idx >= 0; idx-- {
		if len(stack) > 0 && numByte[stack[len(stack)-1]] >= numByte[idx] {
			stack = append(stack, stack[len(stack)-1])
		} else {
			stack = append(stack, idx)
		}

	}
	// fmt.Printf("stack is :%+v\n", stack)

	for i := 0; i < n; i++ {
		if numByte[i] < numByte[stack[len(stack)-i-1]] {
			numByte[i], numByte[stack[len(stack)-i-1]] = numByte[stack[len(stack)-i-1]], numByte[i]
			t, _ := strconv.Atoi(string(numByte))
			ans = max670(ans, t)
			break

		}
	}
	return ans

}

func max670(a, b int) int {
	if a > b {
		return a
	}

	return b

}
