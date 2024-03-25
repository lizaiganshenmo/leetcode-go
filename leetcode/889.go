package leetcode

import "slices"

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(preorder, postorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 { // 空节点
		return nil
	}
	if n == 1 { // 叶子节点
		return &TreeNode{Val: preorder[0]}
	}
	leftSize := slices.Index(postorder, preorder[1]) + 1 // 左子树的大小
	left := constructFromPrePost(preorder[1:1+leftSize], postorder[:leftSize])
	right := constructFromPrePost(preorder[1+leftSize:], postorder[leftSize:n-1])
	return &TreeNode{preorder[0], left, right}
}
