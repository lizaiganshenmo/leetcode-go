package leetcode

// https://leetcode.cn/problems/maximum-width-of-binary-tree/description/?envType=company&envId=bytedance&favoriteSlug=bytedance-thirty-days
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
	// 思路 : bfs
	// 将这个二叉树视作与满二叉树结构相同 满二叉树 根val i 左子节点 2i 右子节点2i+1
	// 第一层 [1,1]
	// 第二层 [1,2]
	// 第三层 [1,4] 即 [1,2^(n-1)] n为层数
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	ans := 1
	root.Val = 1 // 满二叉树编号
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		ans = max662(ans, queue[len(queue)-1].Val-queue[0].Val+1) // 每层最大宽度 等于最右边节点编号 - 最左边节点编号 + 1
		tmp := queue
		queue = []*TreeNode{}
		for _, node := range tmp {
			t := node.Val * 2
			if node.Left != nil {
				node.Left.Val = t
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = t + 1
				queue = append(queue, node.Right)
			}
		}

	}

	return ans

}

func max662(a, b int) int {
	if a > b {
		return a
	}
	return b
}
