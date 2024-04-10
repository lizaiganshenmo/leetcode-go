package leetcode

// https://leetcode.cn/problems/Ygoe9J/description/
// func combinationSum(candidates []int, target int) [][]int {
// 	// 此方法未去重
// 	// 思路 :动态规划
// 	// dp[i] 表示和为i的所有组合
// 	// dp[i] = append(dp[i-j],j) j->candidates && i>j
// 	n := len(candidates)
// 	if n == 0 {
// 		return [][]int{}
// 	}

// 	sort.Ints(candidates)
// 	dp := make([][][]int, target+1)
// 	dp[0] = make([][]int, 1)
// 	for i := 1; i <= target; i++ {
// 		dp[i] = [][]int{}
// 		for _, val := range candidates {
// 			if val > i {
// 				break
// 			}

// 			for _, v := range dp[i-val] {
// 				t := append([]int{}, v...)
// 				t = append(t, val)
// 				dp[i] = append(dp[i], t)

// 			}

// 		}
// 	}

// 	return dp[target]

// }

func combinationSum(candidates []int, target int) [][]int {
	// 思路 :dfs
	n := len(candidates)
	if n == 0 {
		return [][]int{}
	}

	var path = []int{}
	var ans [][]int
	var dfs func(candidates []int, target, idx int)
	dfs = func(candidates []int, target int, idx int) {
		// 递归终止条件
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}

		// 选择加上当前下标的数
		if target >= candidates[idx] {
			old := path
			path = append(path, candidates[idx])
			dfs(candidates, target-candidates[idx], idx)
			// 复位
			path = old
		}

		// 不选当前下标元素
		dfs(candidates, target, idx+1)
	}

	dfs(candidates, target, 0)
	return ans

}
