package leetcode

import (
	"sort"
)

// https://leetcode.cn/problems/find-the-k-sum-of-an-array/description/?envType=study-plan-v2&envId=huawei-2023-spring-sprint
// // 2386. 找出数组的第 K 大和
// 提示
// 困难
// 给你一个整数数组 nums 和一个 正 整数 k 。你可以选择数组的任一 子序列 并且对其全部元素求和。

// 数组的 第 k 大和 定义为：可以获得的第 k 个 最大 子序列和（子序列和允许出现重复）

// 返回数组的 第 k 大和 。

// 子序列是一个可以由其他数组删除某些或不删除元素排生而来的数组，且派生过程不改变剩余元素的顺序。

// 注意：空子序列的和视作 0 。

// 示例 1：

// 输入：nums = [2,4,-2], k = 5
// 输出：2
// 解释：所有可能获得的子序列和列出如下，按递减顺序排列：
// - 6、4、4、2、2、0、0、-2
// 数组的第 5 大和是 2 。
// 示例 2：

// 输入：nums = [1,-2,3,4,-10,12], k = 16
// 输出：10
// 解释：数组的第 16 大和是 10 。

func kSum(nums []int, k int) int64 {
	// 思路: dfs获取所有可能的子序列和
	n := len(nums)
	// ziNums := []int{0}
	// todo 此题没有完全过， 34/111 内存爆了  ziNums 应当是个容量为k的优先队列
	ziNums := make([]int, 1, 2<<(n-1))
	var dfs func(int) //入参: 下标
	sum := 0
	dfs = func(i int) {
		// 定义终止条件
		if i == n {
			return
		}
		// 选择 nums[i]
		sum += nums[i]
		// fmt.Printf("sum is : %d\n", sum)
		ziNums = append(ziNums, sum)
		dfs(i + 1)

		// 不选
		sum -= nums[i] //复位
		dfs(i + 1)

	}

	dfs(0)

	// 排序
	sort.Slice(ziNums, func(i, j int) bool {
		return ziNums[i] > ziNums[j]
	})
	return int64(ziNums[k-1])

}
