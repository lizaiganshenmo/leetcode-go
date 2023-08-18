package leetcode

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/last-stone-weight-ii/?company_slug=bytedance

// 1049. 最后一块石头的重量 II
// 提示
// 中等
// 717
// company
// 字节跳动
// company
// 谷歌 Google
// 有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。

// 每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

// 如果 x == y，那么两块石头都会被完全粉碎；
// 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
// 最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。

// 示例 1：

// 输入：stones = [2,7,4,1,8,1]
// 输出：1
// 解释：
// 组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
// 组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
// 组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
// 组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。
// 示例 2：

// 输入：stones = [31,26,33,21,40]
// 输出：5

func lastStoneWeightII(stones []int) int {
	// 思路 : 动态规划  01背包
	// 石头总重sum  选一些石头，让其尽可能装sum/2 (背包重量sum/2)
	sort.Ints(stones)

	sum := 0
	n := len(stones)
	for i := range stones {
		sum += stones[i]

	}
	fmt.Println("sum is: ", sum)
	weight := sum / 2

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, weight+1)
	}
	dp[0][0] = true

	// dp[n][weight] 表示前n个石头能否装到weight重量
	// dp[i][w] = dp[i-1][w] || dp[i][w-stones[i]
	for i := range stones {
		for j := 0; j <= weight; j++ {
			if j >= stones[i] {
				dp[i+1][j] = dp[i][j] || dp[i][j-stones[i]]
			} else {
				dp[i+1][j] = dp[i][j]
			}

		}
	}
	fmt.Printf("dp res is: %+v\n", dp)

	// 倒序遍历
	max := 0
	for j := weight; j >= 0; j-- {
		if dp[n][j] {
			max = j
			break
		}
	}
	return sum - max*2

}