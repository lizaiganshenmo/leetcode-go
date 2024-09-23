package leetcode

// https://leetcode.cn/problems/combination-sum/description/?envType=company&envId=bytedance&favoriteSlug=bytedance-thirty-days
func combinationSum1(candidates []int, target int) [][]int {
	// 思路 : dfs
	//
	var dfs func(idx int, target int) // idx数组下标 target目标数
	n := len(candidates)
	path := []int{}
	var ans [][]int
	dfs = func(idx, target int) {
		// 递归终止条件
		if n == idx {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}

		t := candidates[idx]
		if t <= target {
			path = append(path, t)
			dfs(idx, target-t)

			// 复位
			path = path[:len(path)-1]
		}
		// 不选当前下标
		dfs(idx+1, target)
	}

	dfs(0, target)
	return ans

}
