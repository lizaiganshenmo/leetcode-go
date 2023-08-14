package leetcode

// https://leetcode.cn/problems/symmetric-tree/description/?envType=study-plan-v2&envId=huawei-2023-spring-sprint

// 101. 对称二叉树

// 给你一个二叉树的根节点 root ， 检查它是否轴对称。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)

}

func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return check(a.Left, b.Right) && check(a.Right, b.Left)
}
