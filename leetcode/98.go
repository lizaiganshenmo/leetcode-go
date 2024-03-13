package leetcode

// https://leetcode.cn/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	// 思路 :树中序遍历的结果一定是递增的
	nums := []int{}

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		nums = append(nums, root.Val)
		dfs(root.Right)
	}

	dfs(root)

	// fmt.Printf("nums is %+v\n", nums)

	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			return false
		}
	}

	return true

}
