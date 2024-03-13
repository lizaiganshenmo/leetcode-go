package leetcode

// https://leetcode.cn/problems/path-sum-iii/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 没有全部通过84/119
// func pathSum(root *TreeNode, targetSum int) int {
// 	// 思路: dfs
// 	// 记录当前路径总和,以及节点各个path
// 	var ans int
// 	var path []int
// 	var curSum int
// 	var dfs func(root *TreeNode, curSum int)
// 	dfs = func(root *TreeNode, curSum int) {
// 		if root == nil {
// 			return
// 		}

// 		curSum += root.Val
// 		path = append(path, root.Val)

// 		if targetSum > 0 {
// 			for curSum > targetSum && len(path) > 0 {
// 				// path 取出首个元素直至小于等于
// 				curSum -= path[0]
// 				path = path[1:]
// 			}

// 		} else {
// 			for curSum < targetSum && len(path) > 0 {
// 				// path 取出首个元素直至大于等于
// 				curSum -= path[0]
// 				path = path[1:]
// 			}
// 		}

// 		if len(path) > 0 && curSum == targetSum {
// 			ans++
// 			// path = make([]int, 0)
// 		}

// 		old := make([]int, len(path))
// 		copy(old, path)
// 		dfs(root.Left, curSum)

// 		// 复位
// 		path = old
// 		dfs(root.Right, curSum)

// 	}

// 	dfs(root, curSum)

// 	return ans

// }

func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}
