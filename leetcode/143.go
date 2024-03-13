package leetcode

// https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/description/

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	// 思路 : dfs
	if B == nil {
		return true
	}
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		// if root.Val != B.Val {
		// 	return false
		// }

		isSame := isSameTree1(root, B)
		if isSame {
			return true
		}

		return dfs(root.Left) || dfs(root.Right)

	}

	return dfs(A)

}

func isSameTree1(root1, root2 *TreeNode) bool {
	if root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return isSameTree1(root1.Left, root2.Left) && isSameTree1(root1.Right, root2.Right)

}
