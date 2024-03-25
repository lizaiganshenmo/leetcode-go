package leetcode

import "slices"

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree1(preorder []int, inorder []int) *TreeNode {
	// 思路 : 递归
	// 前序遍历数组preorder[0]即为该树的根节点
	// 则中序遍历该元素左侧即为该节点左子树, 右边则为该元素右子树
	n := len(preorder)
	if n == 0 {
		return nil
	}

	idx := slices.Index(inorder, preorder[0])
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree1(preorder[1:idx+1], inorder[:idx]),
		Right: buildTree1(preorder[idx+1:], inorder[idx+1:]),
	}

}
