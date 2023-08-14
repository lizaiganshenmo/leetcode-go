package leetcode

import "fmt"

// https://leetcode.cn/problems/jump-game-ii/?envType=study-plan-v2&envId=huawei-2023-spring-sprint

// 45. 跳跃游戏 II
// 中等
// 2.2K
// company
// 亚马逊
// company
// 谷歌 Google
// company
// 奥多比 Adobe
// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

// 0 <= j <= nums[i]
// i + j < n
// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

// 示例 1:

// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//      从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
// 示例 2:

// 输入: nums = [2,3,0,1,4]
// 输出: 2

// 提示:

// 1 <= nums.length <= 104
// 0 <= nums[i] <= 1000
// 题目保证可以到达 nums[n-1]

func jump(nums []int) int {
	// 思路 : 贪心算法  题意  每个元素 nums[i] 表示从索引 i 向前跳转的最大长度
	//       那么维护一个 变量  , 在遍历过程中记录从当前下标走的话，下个最大下标
	//       局部最优解扩展得到最终最优解
	// [start, end] 窗口走
	n := len(nums)
	end := 0   // 初始位置为下标0
	start := 0 // 本次滑动开始位置
	step := 0  // 记录最终步数
	for end < n-1 {
		tmp := 0 // 记录从当前位置可走的最大长度
		for i := start; i <= end; i++ {
			// 从i这个位置, 找寻最大的下个节点
			tmp = max(tmp, nums[i]+i)
		}
		fmt.Printf("end is: %+v\n", end)
		start = end + 1
		end = tmp // 当前这个区间可以走的最远位置

		step++

	}
	return step

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
