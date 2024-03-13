package leetcode

// https://leetcode.cn/problems/path-sum/description/

func hasPathSum(root *TreeNode, targetSum int) bool {
	// dfs
	var dfs func(root *TreeNode, targetSum int) bool
	dfs = func(root *TreeNode, targetSum int) bool {
		if root == nil {
			return false
		}
		if root.Left == nil && root.Right == nil {
			if root.Val == targetSum {
				return true
			} else {
				return false
			}
		}

		return dfs(root.Left, targetSum-root.Val) || dfs(root.Right, targetSum-root.Val)

	}

	return dfs(root, targetSum)

}
