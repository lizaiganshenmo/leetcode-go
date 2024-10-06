package leetcode

import "sort"

// https://leetcode.cn/problems/permutations-ii/description/

func permuteUnique(nums []int) [][]int {
	// 思路 : dfs + 剪枝
	// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
	// 题意有重复数字, 故需要剪枝,去重
	sort.Ints(nums)

	n := len(nums)
	var ans [][]int
	path := []int{}
	used := make([]bool, n) // 对应下标是否已经遍历
	var dfs func(num int)   // num:已添加个数
	dfs = func(num int) {
		// 递归终止条件
		if len(path) == n {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i, val := range nums {
			// 剪枝: 规定两个相同元素 , 如果上一个相同元素还未使用, 则直接跳过
			// 解释参考: https://leetcode.cn/problems/permutations-ii/solutions/417937/quan-pai-lie-ii-by-leetcode-solution/comments/800768
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}

			path = append(path, val)
			used[i] = true
			dfs(num + 1)

			// 复位
			path = path[0 : len(path)-1]
			used[i] = false
		}

	}

	dfs(0)
	return ans

}
