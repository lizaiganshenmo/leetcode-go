package leetcode

import "fmt"

// https://leetcode.cn/problems/house-robber-ii/?envType=study-plan-v2&envId=huawei-2023-spring-sprint

// 213. 打家劫舍 II
// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

// 示例 1：

// 输入：nums = [2,3,2]
// 输出：3
// 解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
// 示例 2：

// 输入：nums = [1,2,3,1]
// 输出：4
// 解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
//      偷窃到的最高金额 = 1 + 3 = 4 。
// 示例 3：

// 输入：nums = [1,2,3]
// 输出：3

// 提示：

// 1 <= nums.length <= 100
// 0 <= nums[i] <= 1000

// 边界条件为：

// {dp[start]=nums[start]只有一间房屋，则偷窃该房屋dp[start+1]=max⁡(nums[start],nums[start+1])只有两间房屋，偷窃其中金额较高的房屋 \begin{cases} \textit{dp}[\textit{start}] = \textit{nums}[\textit{start}] & 只有一间房屋，则偷窃该房屋 \\ \textit{dp}[\textit{start}+1] = \max(\textit{nums}[\textit{start}], \textit{nums}[\textit{start}+1]) & 只有两间房屋，偷窃其中金额较高的房屋 \end{cases}
// {
// dp[start]=nums[start]
// dp[start+1]=max(nums[start],nums[start+1])
// ​

// 只有一间房屋，则偷窃该房屋
// 只有两间房屋，偷窃其中金额较高的房屋
// ​

// 计算得到 dp[end]\textit{dp}[\textit{end}]dp[end] 即为下标范围 [start,end][\textit{start},\textit{end}][start,end] 内可以偷窃到的最高总金额。

// 分别取 (start,end)=(0,n−2)(\textit{start},\textit{end})=(0,n-2)(start,end)=(0,n−2) 和 (start,end)=(1,n−1)(\textit{start},\textit{end})=(1,n-1)(start,end)=(1,n−1) 进行计算，取两个 dp[end]\textit{dp}[\textit{end}]dp[end] 中的最大值，即可得到最终结果

func rob(nums []int) int {
	// 思路 : 动态规划
	// dp[i] 表示偷前i 家最大的偷到金额
	// dp[i] = max(dp[i-2]+dp[i-1](偷第i家) , dp[i-1](不偷第i家))

	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	// [0,n-2]区间
	dp := make([]int, n)
	dp[1] = nums[0]
	for i := range nums {
		if i == 0 {
			continue
		}
		if i == n-1 {
			continue
		}
		dp[i+1] = maxxx(dp[i-1]+nums[i], dp[i])
	}

	// [1,n-1]区间
	dp1 := make([]int, n)
	dp1[1] = nums[n-1]
	j := 1
	for i := n - 2; i >= 1; i-- {

		dp1[j+1] = maxxx(dp1[j-1]+nums[i], dp1[j])
		j++
	}
	fmt.Printf("ans dp1 is:%+v\n", dp1)

	return maxxx(dp[n-1], dp1[n-1])

}

func maxxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}
