package leetcode

// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func isCompleteTree(root *TreeNode) bool {
// 	// 思路 : 递归
// 	// 未全过 : 99 / 120  思路没问题 , 只不过题中val编号可能不是2i 2i+1，需要自己定义保证
// 	// 完全二叉树定义: 一棵深度为k的有n个结点的二叉树，对树中的结点按从上至下、从左到右的顺序进行编号，
// 	// 如果编号为i（1≤i≤n）的结点与满二叉树中编号为i的结点在二叉树中的位置相同，则这棵二叉树称为完全二叉树。

// 	if root == nil {
// 		return true
// 	}
// 	if root.Left == nil && root.Right == nil {
// 		return true
// 	}

// 	// 最后一层中的所有节点都尽可能靠左
// 	if root.Left == nil && root.Right != nil {
// 		return false
// 	}

// 	// 完全二叉树: 左子树节点val = 根节点val*2 右子树节点val = 根节点val*2+1
// 	t := root.Val * 2
// 	if root.Left != nil && root.Left.Val != t {
// 		return false
// 	}
// 	if root.Right != nil && root.Right.Val != t+1 {
// 		return false
// 	}

// 	return isCompleteTree(root.Left) && isCompleteTree(root.Right)

// }

func isCompleteTree(root *TreeNode) bool {
	return dfs958(root, 1, countNodes(root))
}

// 深度优先遍历，currNodes为当前节点的标号，totalNodes为总节点数
func dfs958(root *TreeNode, currNodes, totalNodes int) bool {
	//terminator（递归终止条件）
	if root == nil {
		return true
	}

	//当前节点的标号必须要小于等于总节点数
	if currNodes > totalNodes {
		return false
	}
	//向下递归（左子树的标号传递为currNodes*2，右子树的标号传递为currNodes*2+1）
	return dfs958(root.Left, currNodes*2, totalNodes) && dfs958(root.Right, currNodes*2+1, totalNodes)
}

// 用于计算树的总节点数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
