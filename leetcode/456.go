package leetcode

import "math"

// https://leetcode.cn/problems/132-pattern/

func find132pattern(nums []int) bool {
	// 思路 :单调栈
	// https://leetcode.cn/problems/132-pattern/solutions/1260411/dan-diao-zhan-by-penghuima-syfz/
	var singleStack []int
	mid := math.MinInt32
	for i := len(nums) - 1; i >= 0; i-- {
		//单调栈模式下找到了次大者，是否比次大者还小的元素存在
		if nums[i] < mid {
			return true
		}
		//保持单调栈的单调性
		for len(singleStack) > 0 && nums[i] > singleStack[len(singleStack)-1] {
			//记录次大者
			mid = singleStack[len(singleStack)-1]
			singleStack = singleStack[:len(singleStack)-1]
		}
		//入栈
		singleStack = append(singleStack, nums[i])
	}
	return false
}
