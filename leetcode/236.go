package leetcode

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 思路 :dfs
	// 如果遍历中节点等于p 或者q 则返回
	// 如果左子树和右子树都包含p或者q , 则root就是最近公共节点
	// 如果左子树不为nil 右为 nil 则pq都在左子树中, 反之亦然
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right

}
