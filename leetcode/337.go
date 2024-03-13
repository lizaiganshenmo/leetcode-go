package leetcode

// https://leetcode.cn/problems/house-robber-iii/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func ro1b(root *TreeNode) int {
	// 思路 : dfs
	// 条件： 相邻两个房子不能同时抢，否则报警
	// 则当前节点最大收益就是
	// max(当前节点抢 + 左边子节点不抢 + 右边子节点不抢, 当前节点不抢 + max(左子节点抢,左子节点不抢) + max(右子节点抢,右子节点不抢)
	var dfs func(root *TreeNode) (notRob, rob int)

	dfs = func(root *TreeNode) (notRob, rob int) {
		if root == nil {
			return 0, 0
		}

		lNotRob, lRob := dfs(root.Left)
		rNotRob, rRob := dfs(root.Right)

		// 选择抢当前root节点
		rob = root.Val + lNotRob + rNotRob
		// 不抢
		notRob = max34(lNotRob, lRob) + max34(rNotRob, rRob)

		return
	}

	notRob, rob := dfs(root)
	return max34(notRob, rob)

}

func max34(a, b int) int {
	if a > b {
		return a
	}

	return b
}
