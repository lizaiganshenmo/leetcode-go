package leetcode

// https://leetcode.cn/problems/diameter-of-binary-tree/description/

func diameterOfBinaryTree(root *TreeNode) int {
	// dfs
	// h(root) == max(h(left), h(right)) + 1

	// res(root) = max(res,h(left) + h(right))
	var ans int
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		lHeight := depth(root.Left)
		rHeight := depth(root.Right)

		ans = max543(ans, lHeight+rHeight)
		return max543(lHeight, rHeight) + 1
	}

	depth(root)
	return ans

}

func max543(a, b int) int {
	if a > b {
		return a
	}

	return b
}
