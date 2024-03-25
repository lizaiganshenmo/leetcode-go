package leetcode

import "slices"

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 思路 : 递归
	// 后续遍历数组postorder[n-1]即为该树的根节点
	// 则中序遍历该元素左侧即为该节点左子树, 右边则为该元素右子树
	n := len(postorder)
	if n == 0 {
		return nil
	}

	idx := slices.Index(inorder, postorder[n-1])
	root := &TreeNode{
		Val:   postorder[n-1],
		Left:  buildTree(inorder[:idx], postorder[:idx]),
		Right: buildTree(inorder[idx+1:], postorder[idx:n-1]),
	}

	return root

}
