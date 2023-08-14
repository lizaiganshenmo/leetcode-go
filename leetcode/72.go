package leetcode

// https://leetcode.cn/problems/edit-distance/description/?envType=study-plan-v2&envId=huawei-2023-spring-sprint
// 72. 编辑距离
// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

// 你可以对一个单词进行如下三种操作：

// 插入一个字符
// 删除一个字符
// 替换一个字符

// 示例 1：

// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
// 示例 2：

// 输入：word1 = "intention", word2 = "execution"
// 输出：5
// 解释：
// intention -> inention (删除 't')
// inention -> enention (将 'i' 替换为 'e')
// enention -> exention (将 'n' 替换为 'x')
// exention -> exection (将 'n' 替换为 'c')
// exection -> execution (插入 'u')

func minDistance(word1 string, word2 string) int {
	// 思路 : 动态规划
	// dp[i][j] 表示word1 前i个字转换成word2前 j 个字符要几步
	// dp[0][0] = 0
	// if word1[i] == word2[j]  dp[i][j] = dp[i-1][j-1]
	// else  dp[i][j] = min(dp[i-1][j-1],dp[i-1][j],dp[i][j-1]) + 1

	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i // 初始化
	}

	for i := 0; i <= m; i++ {

		for j := 1; j <= n; j++ {
			if i == 0 {
				dp[0][j] = j
				continue

			}
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}

		}

	}

	return dp[m][n]

}
func min(nums ...int) int {
	ans := nums[0]
	for i := range nums {
		if ans > nums[i] {
			ans = nums[i]
		}

	}
	return ans

}
