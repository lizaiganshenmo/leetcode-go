package leetcode

import "math"

// https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
	// 思路:dfs 寻找节点的最大的子节点数和最小子节点数
	var ans int
	var dfs func(*TreeNode) (int, int) // 返回最大子节点值,最小子节点值
	dfs = func(tn *TreeNode) (int, int) {
		if tn == nil {
			return math.MinInt, math.MaxInt
		}

		lMax, lMin := dfs(tn.Left)
		rMax, rMin := dfs(tn.Right)
		max, min := max1026(tn.Val, max1026(lMax, rMax)), min1026(tn.Val, min1026(lMin, rMin))
		ans = max1026(ans, max1026(max-tn.Val, tn.Val-min))
		return max, min
	}
	dfs(root)

	return ans

}

func max1026(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min1026(a, b int) int {
	if a < b {
		return a
	}
	return b
}
