package leetcode

// https://leetcode.cn/problems/permutations/description/

func permute(nums []int) [][]int {
	// 思路 : dfs
	n := len(nums)
	set := make(map[int]bool, n)
	for _, val := range nums {
		set[val] = true // 默认初始元素都没有遍历到, 初始化为true
	}

	var ans [][]int
	path := make([]int, 0, n)
	var dfs func(i int) // 已经遍历i个数
	dfs = func(i int) {
		// 递归终止条件
		if i == n {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for num, canAdd := range set {
			if !canAdd {
				continue
			}

			path = append(path, num)
			set[num] = false
			dfs(i + 1)

			// 复位
			path = path[:len(path)-1]
			set[num] = true

		}

	}

	dfs(0)
	return ans

}
