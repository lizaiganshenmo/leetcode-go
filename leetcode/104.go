package leetcode

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/?envType=study-plan-v2&envId=top-interview-150

func maxDepth(root *TreeNode) int {
	return depth(root)

}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	var lDepth, rDepth int
	lDepth = depth(node.Left)
	rDepth = depth(node.Right)
	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}
