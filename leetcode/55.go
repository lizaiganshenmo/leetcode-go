package leetcode

import "fmt"

// https://leetcode.cn/problems/jump-game/?envType=study-plan-v2&envId=top-interview-150

// 55. 跳跃游戏
// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。

// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

// 示例 1：

// 输入：nums = [2,3,1,1,4]
// 输出：true
// 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
// 示例 2：

// 输入：nums = [3,2,1,0,4]
// 输出：false
// 解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。

// 提示：

// 1 <= nums.length <= 104
// 0 <= nums[i] <= 105

func canJump(nums []int) bool {
	// 思路 : 贪心算法 每次走该区间内能走最大的
	n := len(nums)
	left, right := 0, nums[0] // 初始在下标为0点
	for right < n-1 {
		r := right
		for i := left; i <= right; i++ {
			r = max11(r, nums[i]+i)
		}
		if r == right && right < n-1 {
			// 表明停滞不前了
			return false
		}
		left = right + 1
		right = r
		fmt.Printf("left now : %+v right now :%+v\n", left, right)

	}

	return true

}

func max11(a, b int) int {
	if a > b {
		return a
	}

	return b
}
