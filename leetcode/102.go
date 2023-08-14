package leetcode

// https://leetcode.cn/problems/binary-tree-level-order-traversal/?envType=study-plan-v2&envId=huawei-2023-spring-sprint

// 102. 二叉树的层序遍历
// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

// 示例 1：

// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
// 示例 2：

// 输入：root = [1]
// 输出：[[1]]
// 示例 3：

// 输入：root = []
// 输出：[]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// type queue struct{
// 	list []ite

// }
func levelOrder(root *TreeNode) [][]int {
	// 层序遍历 bfs
	// 思路 维护一个queue 同一层的入队,该层遍历完了遍历下一层
	var ans [][]int
	if root == nil {
		return ans
	}

	t := []int{root.Val}
	ans = append(ans, t)
	if root.Left == nil && root.Right == nil {
		return ans
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		nextQueue, tmpList := handleQueueBianli(queue)
		queue = nextQueue
		if len(tmpList) != 0 {
			ans = append(ans, tmpList)
		}
	}

	return ans

}

func handleQueueBianli(queue []*TreeNode) ([]*TreeNode, []int) {
	// var nextQueue []*TreeNode
	// var ans []int
	nextQueue := make([]*TreeNode, 0, 2*len(queue))
	ans := make([]int, 0, 2*len(queue))
	for _, v := range queue {
		if v.Left != nil {
			nextQueue = append(nextQueue, v.Left)
			ans = append(ans, v.Left.Val)
		}
		if v.Right != nil {
			nextQueue = append(nextQueue, v.Right)
			ans = append(ans, v.Right.Val)
		}

	}

	return nextQueue, ans
}
