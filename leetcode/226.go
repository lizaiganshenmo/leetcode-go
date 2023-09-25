package leetcode

// https://leetcode.cn/problems/invert-binary-tree/?envType=study-plan-v2&envId=top-interview-150

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	_ = invertTree(root.Left)
	_ = invertTree(root.Right)

	return root

}
