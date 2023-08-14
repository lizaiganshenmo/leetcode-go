package leetcode

// https://leetcode.cn/problems/two-sum-iv-input-is-a-bst/

// 653. 两数之和 IV - 输入二叉搜索树

// 给定一个二叉搜索树 root 和一个目标结果 k，如果二叉搜索树中存在两个元素且它们的和等于给定的目标结果，则返回 true。

// 示例 1：

// 输入: root = [5,3,6,2,4,null,7], k = 9
// 输出: true
// 示例 2：

// 输入: root = [5,3,6,2,4,null,7], k = 28
// 输出: false

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	// flag := false
	var dfs func(*TreeNode) bool
	dfs = func(tn *TreeNode) bool {
		if tn == nil {
			return false
		}
		if _, ok := m[k-tn.Val]; ok {
			// flag = true
			return true
		}
		m[tn.Val] = struct{}{}

		return dfs(tn.Left) || dfs(tn.Right)

	}

	return dfs(root)

}
