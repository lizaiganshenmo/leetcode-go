package leetcode

// https://leetcode.cn/problems/binary-tree-right-side-view/description/

func rightSideView(root *TreeNode) (ans []int) {
	// 思路 : dfs
	// 优先遍历右侧节点 , 同时记录当前已经看过的节点层深
	high := -1                                  // 哪些层已经遍历, root根节点是0层
	var dfs func(root *TreeNode, curHeight int) // 节点 , 当前节点所在层高
	dfs = func(root *TreeNode, curHeight int) {
		if root == nil {
			return
		}

		if curHeight > high {
			ans = append(ans, root.Val)
			high++
		}

		dfs(root.Right, curHeight+1)
		dfs(root.Left, curHeight+1)

	}

	dfs(root, 0)

	return

}
