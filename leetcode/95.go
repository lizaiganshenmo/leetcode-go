package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// wrong
func generateTreesBad(n int) []*TreeNode {
	// 思路 dfs 遍历所有 [1,n]未根节点的所有情况
	ans := []*TreeNode{}
	if n == 0 {
		return ans
	}

	mapping := initMapping(n)
	var dfs func(i, left, right int) *TreeNode // i :以i为根节点的二叉搜索树遍历
	dfs = func(i, left, right int) *TreeNode {
		// 递归终止条件
		if len(mapping) == 0 {
			mapping = initMapping(n)
			return nil
		}

		tn := &TreeNode{Val: i}

		for t := left; t < i; t++ {
			if _, ok := mapping[t]; !ok {
				continue
			}
			// 左子树
			delete(mapping, t)
			tn.Left = dfs(t, 1, t-1)
			// mapping[left] = struct{}{} // 复位

		}

		for t := i + 1; t <= right; t++ {
			if _, ok := mapping[t]; !ok {
				continue
			}
			delete(mapping, t)
			// 右子树
			tn.Right = dfs(t, t+1, right)
		}

		return tn
	}

	for i := 1; i <= n; i++ {
		ans = append(ans, dfs(i, 1, n))
	}
	return ans

}

func initMapping(n int) map[int]struct{} {
	mapping := make(map[int]struct{}, n) // 是否遍历过了
	for i := 1; i <= n; i++ {
		mapping[i] = struct{}{}
	}

	return mapping
}

func generateTrees(n int) []*TreeNode {
	// 思路 dfs 遍历所有 [1,n]未根节点的所有情况
	var dfs func(left, right int) []*TreeNode // 生成 [left,right] 区间内的二叉搜索树集合
	dfs = func(left, right int) []*TreeNode {
		// 递归终止条件
		if left > right {
			return []*TreeNode{nil}
		}
		var ans []*TreeNode

		for i := left; i <= right; i++ {
			lTrees := dfs(left, i-1)
			rTrees := dfs(i+1, right)
			for _, lTree := range lTrees {
				for _, rTree := range rTrees {
					tn := &TreeNode{Val: i}
					tn.Left = lTree
					tn.Right = rTree
					ans = append(ans, tn)
				}
			}
		}
		return ans
	}

	return dfs(1, n)

}
